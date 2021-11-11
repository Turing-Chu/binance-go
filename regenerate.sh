#!/bin/bash

openapi-generator generate -i "api/openapi.yaml" --skip-validate-spec -c "api/openapitools.json" -g go
# -c "api/openapitools.json"