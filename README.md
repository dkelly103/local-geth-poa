# Geth Ethereum - Single node PoA

## What is this?

- Quickly spin up a single PoA ethereum node for local testing purposes

## Starting the network

    docker compose up --build

## Generating new keys

- To create a keystore (Print flag not required, source available at `./key-gen`):

        ./keygen print

- Add the ethereum address to the genesis with a balance
