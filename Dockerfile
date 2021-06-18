FROM golang:1.13-alpine

RUN apk add --no-cache \
        python3 \
        py3-pip \
    && pip3 install --upgrade pip \
    && pip3 install \
        awscli \
    && rm -rf /var/cache/apk/*
RUN apk add --update go git build-base
RUN go get github.com/onsi/ginkgo/ginkgo

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir /app
ADD ./ /app
WORKDIR /app
ENV WORKSPACECONFIG_S3_URI=
RUN chmod +x startup.sh
ENTRYPOINT ["sh","/app/startup.sh"]