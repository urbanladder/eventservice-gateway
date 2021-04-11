#!/bin/bash
set -e

FILEPATH=/etc/rudderstack/
aws s3 cp --profile ul_stg ${WORKSPACECONFIG_S3_URI} ${FILEPATH}/workspaceConfig.json
go run -mod=vendor main.go