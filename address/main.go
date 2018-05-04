package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "token demo")
	})

	r.POST("/account/create/:name", func(c *gin.Context) {

		_name := c.Params.ByName("name")

		key := NewKey("userKeys/" + _name)
		if key.LoadKey() == nil {
			c.JSON(200, "User already exists")
			return
		}
		if err := key.RestoreOrCreate(); err != nil {
			c.JSON(200, err)
			return
		}
		if err := key.SaveKey(); err != nil {
			c.JSON(200, "SaveKey error")
			return
		}
		// acc := etherdb.Account{User: user, Address: key.PublicKeyAsHexString()}
		// if err := acc.Add(); err != nil {
		// 	log.Println(user, key.PublicKeyAsHexString(), err)
		// }

		c.JSON(200, key.PublicKeyAsHexString())
	})

	r.GET("/account/address/:name", func(c *gin.Context) {
		_name := c.Params.ByName("name")

		addr, err := userAddress(_name)
		if err != nil {
			c.JSON(200, "User does not exist")
			return
		}
		c.JSON(200, addr)
	})

	r.GET("/account/private/:name", func(c *gin.Context) {
		_name := c.Params.ByName("name")
		key := NewKey("userKeys/" + _name)

		if key.LoadKey() != nil {
			c.JSON(200, "User does not exist")
			return
		}
		c.JSON(200, fmt.Sprintf("0x%x", key.Key.D))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":4000")
}
