
# syntax=docker/dockerfile:1
ARG GO_VERSION=1.20 
ARG ALPINE_VERSION=3.17
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder
# use our own image above

ARG VERSION=2
ARG REVISION=2.8.0
ARG COMMIT_HASH=a3f056b1de49d6b1911cb14d1767a3a6f1020110
ARG ENTERPRISE_TOKEN
ARG RACE_ENABLED=false
ARG CGO_ENABLED=0
ARG PKG_NAME=github.com/rudderlabs/release-demo

RUN apk add --update make tzdata ca-certificates

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN BUILD_DATE=$(date "+%F,%T") \
    LDFLAGS="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT_HASH} -X main.buildDate=$BUILD_DATE -X main.builtBy=${REVISION} -X main.builtBy=${REVISION} -X main.enterpriseToken=${ENTERPRISE_TOKEN} " \
    make build

# RUN go build -o devtool ./cmd/devtool/


FROM alpine:${ALPINE_VERSION}

RUN apk update && apk add tzdata
RUN apk -U --no-cache upgrade && \
    apk add --no-cache ca-certificates postgresql-client curl bash

ENV WORKSPACECONFIG_S3_URI=

WORKDIR /app

COPY --from=builder app/rudder-server .
# COPY --from=builder app/build/wait-for-go/wait-for-go .
COPY --from=builder app/build/regulation-worker .
# COPY --from=builder app/devtool .

COPY startup.sh .
RUN chmod ug+x startup.sh

# Create required directories for nginx and beaver
RUN mkdir -p shared/pids public/system shared/sockets nginx/cache nginx/logs
RUN touch /etc/logrotate.d/app

# COPY build/wait-for /
# COPY ./rudder-cli/rudder-cli.linux.x86_64 /usr/bin/rudder-cli
# COPY scripts/generate-event /scripts/generate-event
# COPY scripts/batch.json /scripts/batch.json

ENTRYPOINT ["/app/startup.sh"]
# CMD ["/rudder-server"]
