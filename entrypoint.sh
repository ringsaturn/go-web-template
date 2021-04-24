#!/bin/bash
export PRJ_PATH=`pwd`

SERVICE=$1

if [ x"$SERVICE" = x"dev" ]; then
    ./build/app config/dev.yaml
elif [ x"$SERVICE" = x"prod" ]; then
    export GIN_MODE=release
    ./build/app config/dev.yaml
fi
