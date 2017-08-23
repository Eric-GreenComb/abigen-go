package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// config address info
const key = `{"address":"73e091f128342f54595f6bde83970cb5ec84a6bf","crypto":{"cipher":"aes-128-ctr","ciphertext":"005ba9999e7be29352eb40b95734596c4ecf4ad26ea65b0d1ca28521fa7a2d80","cipherparams":{"iv":"c2203f266a4f442508733a30f001274b"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"600353ee2885f0238d912734a0638869b4ff73f50d3a839a1dfd800bed7051de"},"mac":"c4dcc6de182c9a30aa233883785b06a150000349340f504ee827068ba791fa60"},"id":"753f08b0-04f9-4b9b-8391-070583f6d13d","version":3}`
const pwd = `a1111111`

func main() {
	client, err := ethclient.Dial("/home/eric/.ethereum/geth.ipc")
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		os.Exit(0)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "abigen demo")
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		str := "pong"
		c.String(200, str)
	})

	// step 1. Deploy a new awesome contract for the binding demo
	r.POST("deploy", func(c *gin.Context) {

		txOpt, err := bind.NewTransactor(strings.NewReader(key), pwd)
		if err != nil {
			c.String(200, err.Error())
			return
		}

		address, _, _, err := DeployTest(txOpt, client)
		if err != nil {
			c.String(200, err.Error())
			return
		}

		_address := fmt.Sprintf("0x%x", address)

		c.JSON(200, _address)
	})

	// step 2. ModifyValue
	r.POST("/modify/:address/:value", func(c *gin.Context) {

		_address := c.Params.ByName("address")
		_value := c.Params.ByName("value")

		// address := "0xbcfd694d336e910143a339fad6bcc81a67ff3b7c"

		txOpt, err := bind.NewTransactor(strings.NewReader(key), pwd)
		if err != nil {
			c.String(200, err.Error())
			return
		}
		ts, _ := NewTestTransactor(common.HexToAddress(_address), client)

		_, err = ts.ModifyValue(txOpt, _value)
		if err != nil {
			c.String(200, err.Error())
			return
		}

		c.String(200, "OK")
	})

	// step 3. get value
	r.GET("get/:address", func(c *gin.Context) {
		_address := c.Params.ByName("address")
		// address := "0xbcfd694d336e910143a339fad6bcc81a67ff3b7c"
		tc, _ := NewTestCaller(common.HexToAddress(_address), client)
		str, _ := tc.Value(&bind.CallOpts{Pending: true})
		c.JSON(200, str)
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":3000")
}
