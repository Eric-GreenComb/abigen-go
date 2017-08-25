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
const pwd = `a1111111`

func main() {
	client, err := ethclient.Dial("/home/eric/.ethereum/geth.ipc")
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		os.Exit(0)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "token demo")
	})

	// step 1. Deploy a new awesome contract for the binding demo
	r.POST("deploy", func(c *gin.Context) {

		txOpt, err := bind.NewTransactor(strings.NewReader(key), pwd)
		if err != nil {
			c.String(200, err.Error())
			return
		}

		_initialAmount := big.NewInt(100000000000)
		_tokenName := "IPL"
		_address, _, _, err := DeployHumanStandardToken(txOpt, client, _initialAmount, _tokenName, 10, "IPLS")
		if err != nil {
			c.String(200, err.Error())
			return
		}

		fmt.Println(_address)

		address := fmt.Sprintf("0x%x", _address)

		c.JSON(200, address)
	})

	// step 2. SendCoin
	r.POST("/send/:conaddr/:to/:amount", func(c *gin.Context) {

		_conaddr := c.Params.ByName("conaddr")
		_to := c.Params.ByName("to")
		_amount := c.Params.ByName("amount")
		_int64, err := strconv.ParseInt(_amount, 10, 64)
		if err != nil {
			c.String(200, err.Error())
			return
		}

		txOpt, err := bind.NewTransactor(strings.NewReader(key), pwd)
		if err != nil {
			c.String(200, err.Error())
			return
		}
		ts, _ := NewHumanStandardTokenTransactor(common.HexToAddress(_conaddr), client)

		_bigint := big.NewInt(_int64)
		_, err = ts.Transfer(txOpt, common.HexToAddress(_to), _bigint)
		if err != nil {
			c.String(200, err.Error())
			return
		}

		c.String(200, "OK")
	})

	// step 3. get value
	r.GET("get/:conaddr/:addr", func(c *gin.Context) {
		_conaddr := c.Params.ByName("conaddr")
		_addr := c.Params.ByName("addr")
		_caller, err := NewHumanStandardTokenCaller(common.HexToAddress(_conaddr), client)
		if err != nil {
			c.String(200, err.Error())
			return
		}
		_bigint, err := _caller.BalanceOf(&bind.CallOpts{Pending: true}, common.HexToAddress(_addr))
		if err != nil {
			c.String(200, err.Error())
			return
		}
		fmt.Println("================================")
		fmt.Println(_bigint)
		c.JSON(200, _bigint.String())
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":3000")
}
