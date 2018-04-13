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
const key = `{"address":"004ec07d2329997267ec62b4166639513386f32e","crypto":{"cipher":"aes-128-ctr","cipherparams":{"iv":"87b493d4dd7f328e4c3d4400c306b552"},"ciphertext":"4853877ee45b1ec41dbedc2d237260aef2291d5dca6f640e053e40cb0686343d","kdf":"pbkdf2","kdfparams":{"c":10240,"dklen":32,"prf":"hmac-sha256","salt":"e992921213cb010c901f368b761b91cfab956ec6a1371617af985b1fe5af791c"},"mac":"43c63c503b872f0a9312ac5c5c24a0039edcb708c54e289e6db096341b80c851"},"id":"ca5df0e3-c4c5-9b56-0457-9e23e47b3221","meta":"{}","name":"","version":3}`
const pwd = `user`
const eric = `{"address":"00eaf5c65f32917c922364e4892b5ec5a35baa6e","crypto":{"cipher":"aes-128-ctr","cipherparams":{"iv":"7d2ea255e51c5a20530744590bb51bba"},"ciphertext":"09b5f7ca2fe47a9f522da8a87207e2ada9c93858c915c891c1e5663695bad669","kdf":"pbkdf2","kdfparams":{"c":10240,"dklen":32,"prf":"hmac-sha256","salt":"e2ad7727dae631c8dc00f1efc656bede4945532e2a6736412363d2f8e9838336"},"mac":"4096a009a108fd0a6462d373eb68e0e4bfd885d2061b33029c1d8296754b5015"},"id":"8cca6ede-dc66-b2a8-dfe9-435803d2a474","meta":"{}","name":"","version":3}`
const ericpwd = `eric`

func main() {
	client, err := ethclient.Dial("http://localhost:8540")
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

	// step 2. SendCoin
	r.POST("/eric/send/:conaddr/:to/:amount", func(c *gin.Context) {

		_conaddr := c.Params.ByName("conaddr")
		_to := c.Params.ByName("to")
		_amount := c.Params.ByName("amount")
		_int64, err := strconv.ParseInt(_amount, 10, 64)
		if err != nil {
			c.String(200, err.Error())
			return
		}

		txOpt, err := bind.NewTransactor(strings.NewReader(eric), ericpwd)
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
