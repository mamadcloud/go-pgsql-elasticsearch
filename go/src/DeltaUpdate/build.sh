# /bin/bash

GOOS=linux GOARCH=amd64 go build -o delta_update
mv ./delta_update ../../../dockers/go_services/delta_update