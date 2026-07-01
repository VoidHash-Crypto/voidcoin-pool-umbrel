#!/bin/bash
set -e
mkdir -p /etc/voidpool
envsubst < /etc/voidpool/config.template.yaml > /etc/voidpool/config.yaml
echo "Stratum config written"
exec /usr/local/bin/stratum -config /etc/voidpool/config.yaml
