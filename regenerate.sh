#!/bin/bash

openapi-generator generate \
  -i "api/openapi.yaml" \
  --skip-validate-spec \
  -c "api/openapitools.json" \
  -g go \
  --git-user-id Turing-Chu \
  --git-repo-id binanceapi-go
