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
const eric = `{"address":"008f0194bf7b1da7528132ed098d4351168eb77b","crypto":{"cipher":"aes-128-ctr","cipherparams":{"iv":"533a6eb7934a1c4649ee56e613d66bff"},"ciphertext":"062dbca93704bb1ec3018a78752f7b4870e0cdf7a93b04392547152fd162c74d","kdf":"pbkdf2","kdfparams":{"c":10240,"dklen":32,"prf":"hmac-sha256","salt":"56f335f8a29d87feb97ff2931258b58598f475f53346e26c9d6c9df0f0c6afef"},"mac":"eeabffc95026e0b8a01f686ff7136a24d704719231baaf2efdcd6e7d5c3fe67e"},"id":"4aab680f-ad79-7fca-ef4e-7ab3f46d1919","meta":"{\"description\":\"\",\"passwordHint\":\"\",\"timestamp\":1523529995560}","name":"eric","version":3}`
const ericpwd = `a1111111`

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
	r.POST("/ericsend/:conaddr/:to/:amount", func(c *gin.Context) {

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
