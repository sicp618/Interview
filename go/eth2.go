package main

import (
	"context"
	"fmt"
	_ "math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	url := "http://localhost:7545"
	client, _ := ethclient.Dial(url)
	address := common.HexToAddress("0xEf41Deb6F24d50887Ed8c15a24862A7d5029Cf66")
	fmt.Println(address.Hash())
	// balance, err := client.BalanceAt(context.Background(), address, big.NewInt(1))
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(balance)
}