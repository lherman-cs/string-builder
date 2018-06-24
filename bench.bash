#!/usr/bin/env bash

START=10
END=100000000
INC=10
REPEAT=30
WORD="Lukas"

if [ $# != 1 ]; then
    echo "./bench.bash <program>"
    exit 1
fi

program=$1

for (( i = ${START}; i <= ${END}; i *= ${INC} ))
do
    ./${program} ${WORD} ${i} ${REPEAT}
done