#!/bin/bash
PROJECT_NAME=applet-server
FILE=$(date +%F).log

PID=$(ps x | grep $PROJECT_NAME | grep -v grep | awk '{print $1}')
if [[  $PID ]] ; then
  kill -9 $PID
fi


export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:`pwd`/internal/vad/libs
export GODEBUG=cgocheck=0
go mod download
go mod verify
go mod tidy
mkdir -p bin/
go build  -ldflags "-s -w " -o ./bin/$PROJECT_NAME  `pwd`/cmd/$PROJECT_NAME/...&& ulimit -c unlimited && mkdir -p log/ && nohup bin/$PROJECT_NAME  >log/$FILE 2>&1 &