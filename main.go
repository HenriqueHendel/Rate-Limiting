package main

import (
	"net/http"
	"os"

	"github.com/henriquehendel/rateLimiting/client"
	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")
	server := echo.New()

	server.GET("/", handleTokenedRequest)
	server.POST("/client", handleCreateClient)

	server.Logger.Fatal(server.Start(":" + port))
}

func handleTokenedRequest(c echo.Context) error {
	name := c.QueryParam("clientName")
	if name == "" {
		return c.JSON(http.StatusBadRequest, "Query param 'clientName' is empty!")
	}
	tokenBucket, error := client.GetBucket(name)
	if error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}
	if !tokenBucket.IsRequestAllowed(2) {
		return c.JSON(http.StatusTooManyRequests, "Try later!")
	}
	return c.JSON(http.StatusOK, "Hello World!")
}

func handleCreateClient(c echo.Context) error {
	newClient := client.Client{}
	c.Bind(&newClient)

	createClient := client.SetNewClient(newClient)

	if createClient != nil {
		return c.JSON(http.StatusBadRequest, createClient.Error())
	}

	return c.JSON(http.StatusCreated, newClient)
}
