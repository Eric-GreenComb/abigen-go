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
const key = `{"address":"004ec07d2329997267ec62b4166639513386f32e","crypto":{"cipher":"aes-128-ctr","cipherparams":{"iv":"87b493d4dd7f328e4c3d4400c306b552"},"ciphertext":"4853877ee45b1ec41dbedc2d237260aef2291d5dca6f640e053e40cb0686343d","kdf":"pbkdf2","kdfparams":{"c":10240,"dklen":32,"prf":"hmac-sha256","salt":"e992921213cb010c901f368b761b91cfab956ec6a1371617af985b1fe5af791c"},"mac":"43c63c503b872f0a9312ac5c5c24a0039edcb708c54e289e6db096341b80c851"},"id":"ca5df0e3-c4c5-9b56-0457-9e23e47b3221","meta":"{}","name":"","version":3}`
const pwd = `user`

func main() {
	client, err := ethclient.Dial("http://localhost:8540")
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
