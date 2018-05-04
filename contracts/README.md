# Token

## abigen

abigen --sol Test.sol --pkg main --lang go --out cc.go

## run

- deploy

curl -s -X POST http://localhost:3000/deploy 
  
- /modify/:address/:value

curl -s -X POST http://localhost:3000/modify/0xcf4e066ae93e86b079d4897b8595745515e0cd61/test

- get

http://localhost:3000/get/0xcf4e066ae93e86b079d4897b8595745515e0cd61

