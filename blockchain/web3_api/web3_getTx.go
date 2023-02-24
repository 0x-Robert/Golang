package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccountTransactionsInRange(client *ethclient.Client, account common.Address, startBlock uint64, endBlock uint64) ([]common.Hash, error) {
	var txHashes []common.Hash

	for blockNumber := startBlock; blockNumber <= endBlock; blockNumber++ {
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		fmt.Println("block", block)
		if err != nil {
			return nil, err
		}

		for _, tx := range block.Transactions() {
			//if tx.To() != nil && *tx.To() == account || tx.From() == account {
			//fmt.Println("tx", tx)
			if tx.To() != nil && *tx.To() == account {
				txHashes = append(txHashes, tx.Hash())
			}
		}
	}

	return txHashes, nil
}

func main() {
	client, err := ethclient.Dial("https://goerli.infura.io/v3/8f0f6f1d3eba4554818d7e69ef83053f")
	if err != nil {
		fmt.Println("Failed to connect to Ethereum node:", err)
		return
	}

	account := common.HexToAddress("0x7684992428a8E5600C0510c48ba871311067d74c")
	startBlock := uint64(8536151)
	endBlock := uint64(8536152)

	txHashes, err := GetAccountTransactionsInRange(client, account, startBlock, endBlock)
	if err != nil {
		fmt.Println("Failed to get account transactions:", err)
		return
	}

	fmt.Println("Found", len(txHashes), "transactions involving account", account.Hex(), "in block range", startBlock, "-", endBlock)
}
