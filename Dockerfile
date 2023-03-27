
# syntax=docker/dockerfile:1
ARG GO_VERSION=1.20
ARG ALPINE_VERSION=3.17
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder
ARG VERSION
ARG REVISION
ARG COMMIT_HASH
ARG ENTERPRISE_TOKEN=0
ARG RACE_ENABLED=false
ARG CGO_ENABLED=0
ARG PKG_NAME=github.com/rudderlabs/release-demo

ENV LC_ALL=C
ARG deploy_env=staging
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --update make tzdata ca-certificates

WORKDIR /app
# Create required directories for nginx and beaver
RUN mkdir -p shared/pids public/system shared/sockets nginx/cache nginx/logs beaver
RUN touch /etc/logrotate.d/app beaver/beaver.conf
COPY . ./
# RUN go build -o ./bin/event-gateway .

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN BUILD_DATE=$(date "+%F,%T") \
    LDFLAGS="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT_HASH} -X main.buildDate=$BUILD_DATE -X main.builtBy=${REVISION} -X main.builtBy=${REVISION} -X main.enterpriseToken=0" \
    make build

RUN go build -o devtool ./cmd/devtool/


FROM alpine:${ALPINE_VERSION}

RUN apk update && apk add tzdata
RUN apk -U --no-cache upgrade && \
    apk add --no-cache ca-certificates postgresql-client curl bash

COPY --from=builder app/rudder-server .
COPY --from=builder app/build/wait-for-go/wait-for-go .
COPY --from=builder app/build/regulation-worker .
COPY --from=builder app/devtool .

COPY build/startup.sh /
COPY build/wait-for /
COPY ./rudder-cli/rudder-cli.linux.x86_64 /usr/bin/rudder-cli
COPY scripts/generate-event /scripts/generate-event
COPY scripts/batch.json /scripts/batch.json

ENTRYPOINT ["/startup.sh"]

# ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["/rudder-server"]
