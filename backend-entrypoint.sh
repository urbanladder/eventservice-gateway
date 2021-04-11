#!/bin/bash
set -e

FILEPATH=/etc/rudderstack/
aws s3 cp ${WORKSPACECONFIG_S3_URI} ${FILEPATH}/workspaceConfig.json
go run -mod=vendor main.go