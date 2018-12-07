#!/bin/bash

set -e
# install image
docker pull freewil/bitcoin-testnet-box
#run regtest
docker run -t -i -p 19001:19001 -p 19011:19011 freewil/bitcoin-testnet-box