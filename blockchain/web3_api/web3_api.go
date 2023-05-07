package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	// cl, err := ethclient.Dial("/tmp/geth.ipc")
	// if err != nil {
	// 	panic(err)
	// }
	// _ = cl

	//chainid, err := cl.ChainID(context.Background())
	// if err != nil {
	// 	return err
	// }

	addr := common.HexToAddress("0x7684992428a8E5600C0510c48ba871311067d74c")
	//nonce, err := cl.NonceAt(context.Background(), addr, big.NewInt(14000000))

	fmt.Println("addr", addr)

	blockNum, err := cl.BlockNumber(context.Background())
	if err != nil {
		return err
	}
	q := ethereum.FilterQuery{
		FromBlock: new(big.Int).Sub(blockNum, big.NewInt(10)),
		ToBlock:   blockNum,
		Topics:    [][]common.Hash{common.HexToHash("0x7684992428a8E5600C0510c48ba871311067d74c")},
	}
	logs, err := cl.FilterLogs(context.Background(), q)
	if err != nil {
		return err
	}

}
