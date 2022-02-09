#!/usr/bin/env bash

mkdir -p ./db/node1 ./db/node2 ./db/node3 ./db/node4 ./db/node5
mkdir -p ./logs/node1 ./logs/node2 ./logs/node3 ./logs/node4 ./logs/node5

export FBA_VALs=./scripts/configs/local/fba_validators.json

printf "Launching node 1 at 127.0.0.1:9650\n"
export WEB3_API=enabled
./build/flare --network-id=network-99999 \
    --genesis=scdev.json \
    --public-ip=127.0.0.1 \
    --http-port=9650 \
    --staking-port=9651 \
    --log-dir=./logs/node1 \
    --db-dir=./db/node1 \
    --bootstrap-ips= \
    --bootstrap-ids= \
    --staking-tls-cert-file=./scripts/configs/local/NodeID-5dDZXn99LCkDoEi6t9gTitZuQmhokxQTc.crt \
    --staking-tls-key-file=./scripts/configs/local/NodeID-5dDZXn99LCkDoEi6t9gTitZuQmhokxQTc.key \
    --db-type=leveldb \
    --log-level=debug 2>&1 > ./logs/node1/console.log &
NODE_1_PID=`echo $!`
sleep 3

printf "Launching node 2 at 127.0.0.1:9660\n"
export WEB3_API=enabled
./build/flare --network-id=network-99999 \
    --genesis=scdev.json \
    --public-ip=127.0.0.1 \
    --http-port=9660 \
    --staking-port=9661 \
    --log-dir=./logs/node2 \
    --db-dir=./db/node2 \
    --bootstrap-ips=127.0.0.1:9651 \
    --bootstrap-ids=NodeID-5dDZXn99LCkDoEi6t9gTitZuQmhokxQTc \
    --staking-tls-cert-file=./scripts/configs/local/NodeID-AQghDJTU3zuQj73itPtfTZz6CxsTQVD3R.crt \
    --staking-tls-key-file=./scripts/configs/local/NodeID-AQghDJTU3zuQj73itPtfTZz6CxsTQVD3R.key \
    --db-type=leveldb \
    --log-level=debug 2>&1 > ./logs/node2/console.log &
NODE_2_PID=`echo $!`
sleep 3

printf "Launching node 3 at 127.0.0.1:9670\n"
export WEB3_API=enabled
./build/flare --network-id=network-99999 \
    --genesis=scdev.json \
    --public-ip=127.0.0.1 \
    --http-port=9670 \
    --staking-port=9671 \
    --log-dir=./logs/node3 \
    --db-dir=./db/node3 \
    --bootstrap-ips=127.0.0.1:9651 \
    --bootstrap-ids=NodeID-5dDZXn99LCkDoEi6t9gTitZuQmhokxQTc \
    --staking-tls-cert-file=./scripts/configs/local/NodeID-EkH8wyEshzEQBToAdR7Fexxcj9rrmEEHZ.crt \
    --staking-tls-key-file=./scripts/configs/local/NodeID-EkH8wyEshzEQBToAdR7Fexxcj9rrmEEHZ.key \
    --db-type=leveldb \
    --log-level=debug 2>&1 > ./logs/node3/console.log &
NODE_3_PID=`echo $!`
sleep 3

printf "Launching node 4 at 127.0.0.1:9680\n"
export WEB3_API=enabled
./build/flare --network-id=network-99999 \
    --genesis=scdev.json \
    --public-ip=127.0.0.1 \
    --http-port=9680 \
    --staking-port=9681 \
    --log-dir=./logs/node4 \
    --db-dir=./db/node4 \
    --bootstrap-ips=127.0.0.1:9651 \
    --bootstrap-ids=NodeID-5dDZXn99LCkDoEi6t9gTitZuQmhokxQTc \
    --staking-tls-cert-file=./scripts/configs/local/NodeID-FPAwqHjs8Mw8Cuki5bkm3vSVisZr8t2Lu.crt \
    --staking-tls-key-file=./scripts/configs/local/NodeID-FPAwqHjs8Mw8Cuki5bkm3vSVisZr8t2Lu.key \
    --db-type=leveldb \
    --log-level=debug 2>&1 > ./logs/node4/console.log &
NODE_4_PID=`echo $!`
sleep 3

printf "Launching node 5 at 127.0.0.1:9690\n"
export WEB3_API=enabled
./build/flare --network-id=network-99999 \
    --genesis=scdev.json \
    --public-ip=127.0.0.1 \
    --http-port=9690 \
    --staking-port=9691 \
    --log-dir=./logs/node5 \
    --db-dir=./db/node5 \
    --bootstrap-ips=127.0.0.1:9651 \
    --bootstrap-ids=NodeID-5dDZXn99LCkDoEi6t9gTitZuQmhokxQTc \
    --staking-tls-cert-file=./scripts/configs/local/NodeID-HaZ4HpanjndqSuN252chFsTysmdND5meA.crt \
    --staking-tls-key-file=./scripts/configs/local/NodeID-HaZ4HpanjndqSuN252chFsTysmdND5meA.key \
    --db-type=leveldb \
    --log-level=debug 2>&1 > ./logs/node5/console.log &
NODE_5_PID=`echo $!`
sleep 3

printf "\n"
read -p "Press enter to kill all nodes"
kill $NODE_1_PID
kill $NODE_2_PID
kill $NODE_3_PID
kill $NODE_4_PID
kill $NODE_5_PID