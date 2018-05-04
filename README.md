# abigen-go

abigen go demo

## run geth

### init genesis

geth --datadir "/home/eric/go/geth" init /home/eric/go/geth/peihua.json

### run geth console

geth --rpc --rpccorsdomain "*" --port "30303" --rpcapi "db,eth,net,web3" --unlock '0' --password ~/Library/Ethereum/password --nodiscover --maxpeers '50' --networkid '1234574' --datadir '~/Library/Ethereum' console

geth --rpc --rpccorsdomain "*" --port "30303" --rpcapi "db,eth,net,web3" --unlock '0' --password /home/eric/go/geth/password --nodiscover --maxpeers '50' --networkid '1234574' --datadir '/home/eric/go/geth' --ipcpath /home/eric/.ethereum/geth.ipc console

## run abigen

abigen --sol Test.sol -pkg main --lang go --out cc.go

go run *.go
