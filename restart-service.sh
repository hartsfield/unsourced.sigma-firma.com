#!/bin/bash
export servicePort=9395
export logFilePath=./log_file.txt
trap -- '' SIGTERM
git pull
go build -o unsourced
pkill -f unsourced
nohup ./unsourced > /dev/null & disown
sleep 2
