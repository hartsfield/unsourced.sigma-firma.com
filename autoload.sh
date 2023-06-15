#!/bin/bash
trap -- '' SIGTERM
pkill -f $1
export servicePort=9394
export logFilePath=./log_file.txt
go build -o $1
nohup ./$1
# nohup ./$1 > /dev/null & disown
# ./$1
