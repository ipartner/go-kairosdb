#!/bin/bash

set -e

for (( ; ; ))
do
    if curl -s --head --request GET $1 | grep "204 No Content" > /dev/null
    then
        break
    fi
    sleep 1
done

