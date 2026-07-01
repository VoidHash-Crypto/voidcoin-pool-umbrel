#!/bin/bash
set -e

DATA_DIR="${APP_DATA_DIR:-/data}/node"
mkdir -p "$DATA_DIR"

# Write voidcoin.conf
cat > "$DATA_DIR/voidcoin.conf" << CONF
server=1
listen=1
port=7777
rpcport=7778
rpcuser=${RPC_USER:-voidrpc}
rpcpassword=${RPC_PASSWORD}
rpcallowip=0.0.0.0/0
rpcbind=0.0.0.0
txindex=1
v2transport=0
prune=0
CONF

exec voidcoind -datadir="$DATA_DIR" -conf="$DATA_DIR/voidcoin.conf" "$@"
