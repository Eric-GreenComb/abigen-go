package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// config address info
const key = `{"address":"73e091f128342f54595f6bde83970cb5ec84a6bf","crypto":{"cipher":"aes-128-ctr","ciphertext":"005ba9999e7be29352eb40b95734596c4ecf4ad26ea65b0d1ca28521fa7a2d80","cipherparams":{"iv":"c2203f266a4f442508733a30f001274b"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"600353ee2885f0238d912734a0638869b4ff73f50d3a839a1dfd800bed7051de"},"mac":"c4dcc6de182c9a30aa233883785b06a150000349340f504ee827068ba791fa60"},"id":"753f08b0-04f9-4b9b-8391-070583f6d13d","version":3}`

func main() {
	client, err := ethclient.Dial("/home/eric/.ethereum/geth.ipc")
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		os.Exit(0)
	}

	// step 1. Deploy a new awesome contract for the binding demo

	// txOpt, err := bind.NewTransactor(strings.NewReader(key), "a1111111")
	// if err != nil {
	// 	fmt.Println("Failed to create authorized transactor:", err)
	// 	os.Exit(0)
	// }

	// address, tx, _, err := DeployTest(txOpt, client)
	// if err != nil {
	// 	fmt.Println("Failed to deploy new token contract:", err)
	// 	os.Exit(0)
	// }
	// fmt.Printf("Contract pending deploy: 0x%x\n", address)
	// fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	// step 2. ModifyValue

	// address := "0xbcfd694d336e910143a339fad6bcc81a67ff3b7c"

	// txOpt, err := bind.NewTransactor(strings.NewReader(key), "a1111111")
	// if err != nil {
	// 	fmt.Println("Failed to create authorized transactor:", err)
	// 	os.Exit(0)
	// }
	// ts, _ := NewTestTransactor(common.HexToAddress(address), client)

	// _, err = ts.ModifyValue(txOpt, "3333")
	// if err != nil {
	// 	fmt.Println("TS : ", err)
	// 	os.Exit(0)
	// }

	// step 3. get value

	address := "0xbcfd694d336e910143a339fad6bcc81a67ff3b7c"

	tc, _ := NewTestCaller(common.HexToAddress(address), client)
	str, _ := tc.Value(&bind.CallOpts{Pending: true})
	fmt.Println(str)
}
