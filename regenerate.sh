#!/bin/bash

export GO_POST_PROCESS_FILE="$(which gofmt) -w"
openapi-generator generate \
  -i "api/openapi.yaml" \
  --skip-validate-spec \
  -c "api/openapitools.yaml" \
  -g go \
  --git-user-id Turing-Chu \
  --git-repo-id binanceapi-go
