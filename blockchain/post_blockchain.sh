#!/bin/bash


curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"BPM":50 }' \
  http://localhost:9000
