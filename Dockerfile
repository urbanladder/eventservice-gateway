FROM golang:1.13-alpine
ENV GO111MODULE=on
RUN apk add --no-cache \
        python3 \
        py3-pip \
    && pip3 install --upgrade pip \
    && pip3 install \
        awscli \
    && rm -rf /var/cache/apk/*
RUN apk add --update go git build-base
RUN go get github.com/onsi/ginkgo/ginkgo
RUN mkdir /app
ADD ./ /app
WORKDIR /app
ENV WORKSPACECONFIG_S3_URI=
RUN chmod +x backend-entrypoint.sh
ENTRYPOINT ["sh","/app/backend-entrypoint.sh"]