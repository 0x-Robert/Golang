package main

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	cl, err := ethclient.Dial("/tmp/geth.ipc")
	if err != nil {
		panic(err)
	}
	_ = cl

	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		return err
	}

}
