# /bin/bash

GOOS=linux GOARCH=amd64 go build -o go_api
mv ./go_api ../../../dockers/go_api/go_api