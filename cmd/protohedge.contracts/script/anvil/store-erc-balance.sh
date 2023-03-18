#!/bin/bash

export $(cat ../../.env | xargs)

deps=$(cat ../../config/addresses.arbitrum.json)

read_address () {
    echo $deps | jq $1 | tr -d '"'
}

export USDC=$(read_address ".usdc")

cast call --private-key $PERSONAL_PRIVATE_KEY $USDC "transfer()" $1 $2 
