#!/bin/sh
curl -s \
  -X POST "http://localhost:8081/subjects/topic1/versions" \
  -H "Content-Type: application/vnd.schemaregistry.v1+json" \
  -d "$(cat blobschema.avsc | jq '{schema: . | tostring}')" \
  | jq
