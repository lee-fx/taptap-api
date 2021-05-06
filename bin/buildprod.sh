#! /bin/bash

# build services

cd ~/opt/lfxpupa/taptap/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd ~/opt/lfxpupa/taptap/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd ~/opt/lfxpupa/taptap/stream
env GOOS=linux GOARCH=amd64 go build -o ../bin/stream
