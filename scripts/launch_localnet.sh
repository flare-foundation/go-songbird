#!/usr/bin/env bash

mkdir -p ./db/node1 ./db/node2 ./db/node3 ./db/node4 ./db/node5 ./db/node6
mkdir -p ./logs/node1 ./logs/node2 ./logs/node3 ./logs/node4 ./logs/node5 ./logs/node6

export CUSTOM_VALIDATORS="NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg,NodeID-GWPcbFJZFfZreETSoWjPimr846mXEKCtu,NodeID-K9vx5sYL3aAq4Jt4SmXY1FasP2hpwpPNu"

printf "Launching node 1 at 127.0.0.1:9650\n"
./build/flare --network-id=local \
    --public-ip=127.0.0.1 \
    --http-port=9650 \
    --staking-port=9651 \
    --log-dir=./logs/node1 \
    --db-dir=./db/node1 \
    --bootstrap-ips= \
    --bootstrap-ids= \
    --staking-tls-cert-file=./scripts/keys/NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg.crt \
    --staking-tls-key-file=./scripts/keys/NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg.key \
    --db-type=leveldb \
    --log-level=debug 2>&1 > ./logs/node1/console.log &
NODE_1_PID=`echo $!`
sleep 1

printf "Launching node 2 at 127.0.0.1:9660\n"
./build/flare --network-id=local \
    --public-ip=127.0.0.1 \
    --http-port=9660 \
    --staking-port=9661 \
    --log-dir=./logs/node2 \
    --db-dir=./db/node2 \
    --bootstrap-ips=127.0.0.1:9651 \
    --bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
    --staking-tls-cert-file=./scripts/keys/NodeID-GWPcbFJZFfZreETSoWjPimr846mXEKCtu.crt \
    --staking-tls-key-file=./scripts/keys/NodeID-GWPcbFJZFfZreETSoWjPimr846mXEKCtu.key \
    --db-type=leveldb \
    --log-level=debug 2>&1 > ./logs/node2/console.log &
NODE_2_PID=`echo $!`
sleep 1

printf "Launching node 3 at 127.0.0.1:9670\n"
./build/flare --network-id=local \
    --public-ip=127.0.0.1 \
    --http-port=9670 \
    --staking-port=9671 \
    --log-dir=./logs/node3 \
    --db-dir=./db/node3 \
    --bootstrap-ips=127.0.0.1:9651 \
    --bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
    --staking-tls-cert-file=./scripts/keys/NodeID-K9vx5sYL3aAq4Jt4SmXY1FasP2hpwpPNu.crt \
    --staking-tls-key-file=./scripts/keys/NodeID-K9vx5sYL3aAq4Jt4SmXY1FasP2hpwpPNu.key \
    --db-type=leveldb \
    --log-level=debug 2>&1 > ./logs/node3/console.log &
NODE_3_PID=`echo $!`
sleep 1

printf "Launching node 4 at 127.0.0.1:9680\n"
./build/flare --network-id=local \
    --public-ip=127.0.0.1 \
    --http-port=9680 \
    --staking-port=9681 \
    --log-dir=./logs/node4 \
    --db-dir=./db/node4 \
    --bootstrap-ips=127.0.0.1:9651 \
    --bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
    --staking-tls-cert-file=./scripts/keys/NodeID-MFrZFVCXPv5iCn6M9K6XduxGTYp891xXZ.crt \
    --staking-tls-key-file=./scripts/keys/NodeID-MFrZFVCXPv5iCn6M9K6XduxGTYp891xXZ.key \
    --db-type=leveldb \
    --log-level=debug 2>&1 > ./logs/node4/console.log &
NODE_4_PID=`echo $!`
sleep 1

printf "Launching node 5 at 127.0.0.1:9690\n"
./build/flare --network-id=local \
    --public-ip=127.0.0.1 \
    --http-port=9690 \
    --staking-port=9691 \
    --log-dir=./logs/node5 \
    --db-dir=./db/node5 \
    --bootstrap-ips=127.0.0.1:9651 \
    --bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
    --staking-tls-cert-file=./scripts/keys/NodeID-NFBbbJ4qCmNaCzeW7sxErhvWqvEQMnYcN.crt \
    --staking-tls-key-file=./scripts/keys/NodeID-NFBbbJ4qCmNaCzeW7sxErhvWqvEQMnYcN.key \
    --db-type=leveldb \
    --log-level=debug 2>&1 > ./logs/node5/console.log &
NODE_5_PID=`echo $!`
sleep 1

printf "Launching node 6 at 127.0.0.1:9690\n"
./build/flare --network-id=local \
    --public-ip=127.0.0.1 \
    --http-port=9700 \
    --staking-port=9701 \
    --log-dir=./logs/node6 \
    --db-dir=./db/node6 \
    --bootstrap-ips=127.0.0.1:9651 \
    --bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
    --staking-tls-cert-file=./scripts/keys/NodeID-P7oB2McjBGgW2NXXWVYjV8JEDFoW9xDE5.crt \
    --staking-tls-key-file=./scripts/keys/NodeID-P7oB2McjBGgW2NXXWVYjV8JEDFoW9xDE5.key \
    --db-type=leveldb \
    --log-level=debug 2>&1 > ./logs/node6/console.log &
NODE_6_PID=`echo $!`
sleep 1

printf "\n"
read -p "Press enter to kill all nodes"
kill $NODE_1_PID
kill $NODE_2_PID
kill $NODE_3_PID
kill $NODE_4_PID
kill $NODE_5_PID