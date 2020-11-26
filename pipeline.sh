#!/bin/bash
cd /Users/prima/playground/crudsample-deployment/crudsample
# build bin
env GOOS=linux GOARCH=amd64 go build -o ./bin/crudsample ./cmd/server/main.go
# build docker image using compose
docker-compose build --no-cache
#tag the image
docker tag crudsample_backend prima101112/crudsample:$1
#push the image
docker push prima101112/crudsample:$1
#deploy latest main
argocd app create crudsample --repo https://github.com/prima101112/crudsample.git --path helm --dest-server https://192.168.64.3:8443 --dest-namespace crudsample


#trial argo