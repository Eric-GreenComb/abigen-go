# Token

## abigen

abigen --sol HumanStandardToken.sol -pkg main --lang go --out cc.go


solc --libraries BigInt:0x6fee4929bbe6156903ffa602f2eb997ea6c8b20f --bin --abi --overwrite Test.sol -o ./
abigen --abi Test.abi --bin Test.bin -solc /usr/local/bin/solc -type Test -pkg main --lang go --out Test.go

## run

- deploy

curl -s -X POST http://localhost:3000/deploy 
  
- send

curl -s -X POST http://localhost:3000/send/:conaddr/:to/:amount

curl -s -X POST http://localhost:3000/send/7ba4324585cb5597adc283024819254345cd7c62/00Aa39d30F0D20FF03a22cCfc30B7EfbFca597C2/10000

curl -s -X POST http://localhost:3000/send/7ba4324585cb5597adc283024819254345cd7c62/0x008f0194Bf7B1dA7528132Ed098D4351168eB77b/100000

curl -s -X POST http://localhost:3000/ericsend/7ba4324585cb5597adc283024819254345cd7c62/0x00Aa39d30F0D20FF03a22cCfc30B7EfbFca597C2/10000

- get

http://localhost:3000/get/7ba4324585cb5597adc283024819254345cd7c62/00Aa39d30F0D20FF03a22cCfc30B7EfbFca597C2

granola robin cone reverb stick pristine entryway flint canopy landing walnut finite