package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//eth client 설정, 나의 경우에는 Ganache를 설정했다.
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}
	//client의 값 출력
	fmt.Println("client", client) //&{0xc000118100}
	account := common.HexToAddress("0x09B225298c2C9F8428c198068588d10B677d1A08")
	fmt.Println("account", account)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance", balance) // balance 90288025923715000000

	// 272번 블록 시점의 계정잔액을 확인할 수 있다. 블록번호는 big.Int여야한다.
	blockNumber := big.NewInt(272)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balanceAt 잔액:", balanceAt) // balanceAt 잔액: 90288025923715000000

	// 이더리움의 숫자는 고정 소수점 정밀도이기 때문에 가장 작은 단위로 처리되며
	// ETH값을 읽으려면 Wei이다
	// 큰 숫자를 다루기 때문에 go math, big package를 가져와야한다.
	fbalance := new(big.Float)
	// balance에 balanceAt 값을 문자열로 변환하여 할당한다.
	fbalance.SetString(balanceAt.String())
	//fbalance를 10^18로 나누어 이더 값으로 변환한다.
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	//이더 값을 출력한다.
	fmt.Println("ethValue", ethValue) //ethValue 90.288025923715
	//아직 처리되지 않은 잔액을 가져온다.
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println("pendingBalance", pendingBalance) //pendingBalance 90288025923715000000
}
