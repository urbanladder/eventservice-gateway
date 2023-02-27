FROM ${BASE_IMAGE_REPO_URL}/golang_base:latest

RUN apk add --update go git build-base
#RUN go get github.com/onsi/ginkgo/ginkgo

ENV LC_ALL=C
ARG deploy_env=staging

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