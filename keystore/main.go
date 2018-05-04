package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "token demo")
	})

	r.POST("/account/create/:password", func(c *gin.Context) {

		_password := c.Params.ByName("password")

		dir, ks := tmpKeyStore(true)
		fmt.Println(dir)
		fmt.Println(ks)
		a, err := ks.NewAccount(_password)
		if err != nil {
			c.JSON(200, err)
			return
		}
		fmt.Println(a.URL)
		c.JSON(200, a.Address)
	})

	r.GET("/account/address/:password", func(c *gin.Context) {
		_password := c.Params.ByName("password")

		keyjson, err := ioutil.ReadFile("key/UTC--2018-04-26T17-46-51.655094678Z--fa6fa302b62e33610062ff8cd540ed17b0981b86")
		if err != nil {
			c.JSON(200, "file is not exit")
			return
		}

		key, err := keystore.DecryptKey(keyjson, _password)
		if err != nil {
			c.JSON(200, "_password is error")
			return
		}
		c.JSON(200, key.Address)
	})

	r.GET("/account/private/:password", func(c *gin.Context) {

		_password := c.Params.ByName("password")

		keyjson, err := ioutil.ReadFile("key/UTC--2018-04-26T17-46-51.655094678Z--fa6fa302b62e33610062ff8cd540ed17b0981b86")
		if err != nil {
			c.JSON(200, "file is not exit")
			return
		}

		key, err := keystore.DecryptKey(keyjson, _password)
		if err != nil {
			c.JSON(200, "_password is error")
			return
		}
		c.JSON(200, fmt.Sprintf("%x", key.PrivateKey.D))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":3000")
}

func tmpKeyStore(encrypted bool) (string, *keystore.KeyStore) {
	d := "./key"

	new := keystore.NewPlaintextKeyStore
	if encrypted {
		new = func(kd string) *keystore.KeyStore {
			return keystore.NewKeyStore(kd, veryLightScryptN, veryLightScryptP)
		}
	}
	return d, new(d)
}
