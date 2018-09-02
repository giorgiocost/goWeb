#!/bin/bash

if [ $* == 'start' ]; then
    sudo ./development.sh env && go run main.go
    exit 1
fi

if [ $* == 'env' ]; then
    dep ensure
    exit 1
fi
 
if [ $# -lt 1 ]; then
   echo "Faltou utilizar pelo menos um argumento!"
   exit 1
fi


#  ./development.sh