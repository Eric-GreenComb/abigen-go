# Token

## abigen

abigen --sol HumanStandardToken.sol -pkg main --lang go --out cc.go


solc --libraries BigInt:0x6fee4929bbe6156903ffa602f2eb997ea6c8b20f --bin --abi --overwrite Test.sol -o ./
abigen --abi Test.abi --bin Test.bin -solc /usr/local/bin/solc -type Test -pkg main --lang go --out Test.go

## run

- deploy

curl -s -X POST http://localhost:3000/deploy/IPL/IPL 

0xedfc9c2f4cfa7495c1a95cfe1cb856f5980d5e18

- send

curl -s -X POST http://localhost:3000/send/:conaddr/:to/:amount

curl -s -X POST http://localhost:3000/send/0xedfc9c2f4cfa7495c1a95cfe1cb856f5980d5e18/0x00Aa39d30F0D20FF03a22cCfc30B7EfbFca597C2/1025

- send 2 eric

curl -s -X POST http://localhost:3000/send/0xEa4C690E7560070071643960f20bA984c8A922cd/0x00eaf5c65f32917c922364e4892b5ec5a35baa6e/1000000

- eric send 2 alan

curl -s -X POST http://localhost:3000/eric/send/0xEa4C690E7560070071643960f20bA984c8A922cd/0x00a5524ceac2e06c7f48085fb07d0913cfcda2b6/10000

- eric send 2 node2

curl -s -X POST http://localhost:3000/eric/send/0xEa4C690E7560070071643960f20bA984c8A922cd/0x002E28950558Fbede1A9675Cb113F0BD20912019/1000

- eric send 2 node3

curl -s -X POST http://localhost:3000/eric/send/0xEa4C690E7560070071643960f20bA984c8A922cd/0x00a94Ac799442FB13De8302026fd03068bA6A428/1022

- get

http://localhost:3000/get/0xEa4C690E7560070071643960f20bA984c8A922cd/0x00eaf5c65f32917c922364e4892b5ec5a35baa6e

- add user

curl --data '{"jsonrpc":"2.0","method":"parity_newAccountFromPhrase","params":["eric", "eric"],"id":0}' -H "Content-Type: application/json" -X POST localhost:8541

0x00eaf5c65f32917c922364e4892b5ec5a35baa6e

curl --data '{"jsonrpc":"2.0","method":"parity_newAccountFromPhrase","params":["alan", "alan"],"id":0}' -H "Content-Type: application/json" -X POST localhost:8541

0x00a5524ceac2e06c7f48085fb07d0913cfcda2b6