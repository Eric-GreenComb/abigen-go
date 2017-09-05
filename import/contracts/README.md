
## deploy import

- abigen ConvertLib.sol  
abigen --sol ConvertLib.sol -pkg main --lang go --out cc.go --out ConvertLib.go

- solc MetaCoin.sol  
solc --libraries ConvertLib:0x21ea709d33e96b60e2951ed887349e92037cd23e --bin --abi --overwrite MetaCoin.sol -o ./

- abigen MetaCoin.sol  
abigen --abi MetaCoin.abi --bin MetaCoin.bin -solc /usr/bin/solc -type MetaCoin -pkg main --lang go --out MetaCoin.go