# bridge-chain-tendermint

``` 
# Init
TMHOME=chainroot1 go run ./cmd/. init
TMHOME=chainroot2 go run ./cmd/. init
TMHOME=chainroot3 go run ./cmd/. init

# Check node1's node ID
TMHOME=chainroot1 go run ./cmd/. show_node_id

# Copy genesis file from chainroot1/config/config.toml to chainroot2/config/config.toml
# Change ports in node1's config (chainroot1/config/config.toml)

# allow_duplicate_ip = true

# Run Nodes

TMHOME=chainroot1 go run ./cmd/. node 

TMHOME=chainroot2 go run ./cmd/. node --p2p.persistent_peers={node1's node ID}@127.0.0.1:26656
TMHOME=chainroot2 go run ./cmd/. node --p2p.persistent_peers=66df038cb19856c2c2142334bf23158864285252@127.0.0.1:26656,75020c9dabbc75b9fdb0a3f40b3fcb0fa29458a9@127.0.0.1:17756


TMHOME=chainroot2 go run ./cmd/. node --p2p.seeds=66df038cb19856c2c2142334bf23158864285252@127.0.0.1:26656 --p2p.persistent_peers=75020c9dabbc75b9fdb0a3f40b3fcb0fa29458a9@127.0.0.1:17756

TMHOME=chainroot3 go run ./cmd/. node --p2p.persistent_peers=66df038cb19856c2c2142334bf23158864285252@127.0.0.1:26656,4d01c82696d8dbf26c2c2b765a6fbdfcfa38be06@127.0.0.1:16656

TMHOME=chainroot3 go run ./cmd/. node --p2p.seeds=66df038cb19856c2c2142334bf23158864285252@127.0.0.1:26656,4d01c82696d8dbf26c2c2b765a6fbdfcfa38be06@127.0.0.1:16656

```