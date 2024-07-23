#!/bin/bash

# build for amd64
env GOOS=linux GOARCH=amd64 go build -o ./target/ip-updater_amd64

# build for arm64 (raspberry pi 4)
env GOOS=linux GOARCH=arm64 go build -o ./target/ip-updater_arm64

## build for arm5
#env GOOS=linux GOARCH=arm GOARM=arm5 go build -d /target/ -o ip-updater_arm5
#
## build for arm6
#env GOOS=linux GOARCH=arm GOARM=arm6 go build -d /target/ -o ip-updater_arm6
#
## build for arm7
#env GOOS=linux GOARCH=arm GOARM=arm7 go build -d /target/ -o /target/ip-updater_arm7