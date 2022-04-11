#!/bin/bash

GO_MOD_NAME=`cat go.mod | grep "^module " | awk -F ' ' '{print $2}'`

function compile()
{
    go mod tidy

    ldflagsv="-X '${GO_MOD_NAME}/gval.GitCommitID=`git rev-parse HEAD`'"
    ldflagsv+=" -X '${GO_MOD_NAME}/gval.BuildTime=`date +"%Y-%m-%d %H:%M:%S"`'"

    rm -rf output/
    mkdir -p output/bin output/conf
    for app in $(ls cmd/); do
	    go build -o output/bin/${app} -ldflags "${ldflagsv}" cmd/${app}/main.go
        cp -f cmd/${app}/*.ini output/conf
    done
}

compile
