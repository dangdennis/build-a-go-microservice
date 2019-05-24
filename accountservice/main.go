package main

import (
	"fmt"

	"github.com/dangdennis/build-a-go-microservice/accountservice/dbclient"
	"github.com/dangdennis/build-a-go-microservice/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient() {
	var DBClient dbclient.IBoltClient
	// service.DBClient = dbclient.IBoltClient{}
	DBClient.OpenBoltDb()
	DBClient.Seed()
}
