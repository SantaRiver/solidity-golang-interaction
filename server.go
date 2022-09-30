package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/santariver/solidity-golang/api"
	"net/http"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	conn, err := api.NewApi(common.HexToAddress("0x0Fa19Dc4fA976E382D06C6f99bFac8F3fea689E6"), client)
	if err != nil {
		panic(err)
	}
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/greet/:message", func(c echo.Context) error {
		message := c.Param("message")
		reply, err := conn.Greet(&bind.CallOpts{}, message)
		if err != nil {
			return err
		}
		err = c.JSON(http.StatusOK, reply)
		if err != nil {
			return err
		}
		return err
	})
	e.GET("/hello", func(c echo.Context) error {
		reply, err := conn.Hello(&bind.CallOpts{})
		if err != nil {
			return err
		}
		err = c.JSON(http.StatusOK, reply)
		if err != nil {
			return err
		} // Hello World
		return err
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
