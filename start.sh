#!/bin/sh

DATA_DIRECTORY=/root/.ethereum

rm -rf $DATA_DIRECTORY/geth | true

geth --datadir $DATA_DIRECTORY \
        --miner.etherbase $ETHER_BASE \
        init /genesis.json

geth --datadir $DATA_DIRECTORY --dev \
        --miner.etherbase $ETHER_BASE \
        --http --http.addr '0.0.0.0' --http.port 8545 --http.corsdomain '*' --http.api 'admin,eth,miner,net,web3,personal,txpool' --allow-insecure-unlock \
        --ws --ws.addr '0.0.0.0' --ws.port 8546 --ws.origins '*' \
        --password "/dev/null" \
        --unlock $UNLOCK \
        --verbosity 3 \
        --syncmode full \
        --mine