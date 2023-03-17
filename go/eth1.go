package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	url := "https://mainnet.infura.io/v3/your-api-key"
	url = "http://localhost:7545"
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")
	fmt.Println(client.BlockNumber(context.Background()))
}