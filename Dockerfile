FROM 540798973460.dkr.ecr.us-east-1.amazonaws.com/golang_base:latest

RUN apk add --update go git build-base
RUN go get github.com/onsi/ginkgo/ginkgo

ENV LC_ALL=C

ARG deploy_env=staging

ENV APP_ENV=staging
ENV appName=eventservice-backend
ENV DEPLOY_ENV=staging
ENV etcdHost=stg-docker-secrets.urbanladder.com 
ENV JOB_NAME=pm2
ENV robotsurl=

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Create required directories for nginx and beaver
RUN mkdir -p shared/pids public/system shared/sockets nginx/cache nginx/logs beaver
RUN touch /etc/logrotate.d/app beaver/beaver.conf

COPY . ./
RUN go build -o ./bin/event-gateway .
ENV WORKSPACECONFIG_S3_URI=
RUN chmod ug+x startup.sh
ENTRYPOINT ["sh","/app/startup.sh"]