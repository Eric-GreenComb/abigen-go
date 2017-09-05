package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// config address info
const key = `{"address":"73e091f128342f54595f6bde83970cb5ec84a6bf","crypto":{"cipher":"aes-128-ctr","ciphertext":"005ba9999e7be29352eb40b95734596c4ecf4ad26ea65b0d1ca28521fa7a2d80","cipherparams":{"iv":"c2203f266a4f442508733a30f001274b"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"600353ee2885f0238d912734a0638869b4ff73f50d3a839a1dfd800bed7051de"},"mac":"c4dcc6de182c9a30aa233883785b06a150000349340f504ee827068ba791fa60"},"id":"753f08b0-04f9-4b9b-8391-070583f6d13d","version":3}`
const contAddr = `0xb7003552dc796f0d045b27c991d2e268649bda08`
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

	// step 1. deploy-lib
	// abigen --sol ConvertLib.sol -pkg main --lang go --out cc.go --out ConvertLib.go
	// return address: 0x21ea709d33e96b60e2951ed887349e92037cd23e
	r.POST("deploy-lib", func(c *gin.Context) {

		txOpt, err := bind.NewTransactor(strings.NewReader(key), pwd)
		if err != nil {
			c.String(200, err.Error())
			return
		}

		_address, _, _, err := DeployConvertLib(txOpt, client)
		if err != nil {
			c.String(200, err.Error())
			return
		}

		fmt.Println(_address)

		address := fmt.Sprintf("0x%x", _address)

		c.JSON(200, address)
	})

	// step 2. Deploy a new awesome contract for the binding demo
	// - solc MetaCoin.sol
	// - abigen MetaCoin.sol
	// return 0xb7003552dc796f0d045b27c991d2e268649bda08
	r.POST("deploy-metacoin", func(c *gin.Context) {

		txOpt, err := bind.NewTransactor(strings.NewReader(key), pwd)
		if err != nil {
			c.String(200, err.Error())
			return
		}

		_address, _, _, err := DeployMetaCoin(txOpt, client)
		if err != nil {
			c.String(200, err.Error())
			return
		}

		fmt.Println(_address)

		address := fmt.Sprintf("0x%x", _address)

		c.JSON(200, address)
	})

	// step 3. ModifyValue
	r.POST("/send/:address/:value", func(c *gin.Context) {

		_address := c.Params.ByName("address")
		_value := c.Params.ByName("value")
		_int64, err := strconv.ParseInt(_value, 10, 64)

		// address := "0xbcfd694d336e910143a339fad6bcc81a67ff3b7c"

		txOpt, err := bind.NewTransactor(strings.NewReader(key), pwd)
		if err != nil {
			c.String(200, err.Error())
			return
		}
		ts, _ := NewMetaCoinTransactor(common.HexToAddress(contAddr), client)

		_bigint := big.NewInt(_int64)
		_, err = ts.SendCoin(txOpt, common.HexToAddress(_address), _bigint)
		if err != nil {
			c.String(200, err.Error())
			return
		}

		c.String(200, "OK")
	})

	// step 4. get value
	r.GET("get/:address", func(c *gin.Context) {
		_address := c.Params.ByName("address")
		// address := "0xbcfd694d336e910143a339fad6bcc81a67ff3b7c"
		tc, _ := NewMetaCoinCaller(common.HexToAddress(contAddr), client)
		_bigint, _ := tc.GetBalance(&bind.CallOpts{Pending: true}, common.HexToAddress(_address))
		c.JSON(200, _bigint.String())
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":3000")

}
