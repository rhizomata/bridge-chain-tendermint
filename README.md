# bridge-chain-tendermint

``` 
# Init
TMHOME=chainroot1 go run ./cmd/. init
TMHOME=chainroot2 go run ./cmd/. init

# Check node1's node ID
TMHOME=chainroot1 go run ./cmd/. show_node_id

# Copy genesis file from chainroot1/config/config.toml to chainroot2/config/config.toml
# Change ports in node2's config (chainroot2/config/config.toml)

# Run Nodes

TMHOME=chainroot1 go run ./cmd/. node 

TMHOME=chainroot2 go run ./cmd/. node --p2p.persistent_peers={node1's node ID}@127.0.0.1:26656
TMHOME=chainroot2 go run ./cmd/. node --p2p.persistent_peers=6b24296bbbe0d90cf33ee9a6f9769cbef83082f6@127.0.0.1:26656



```