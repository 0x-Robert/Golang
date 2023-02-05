#!/bin/bash

#curl -d "BPM=50" localhost:8888 
#curl -d "userId=kimcoding&password=1234" https://localhost:4000/login 

#curl -X POST https://localhost:8888
#   -H 'Content-Type: application/json'
#   -d '{"BPM":50}'

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"BPM":50 }' \
  http://localhost:8888
