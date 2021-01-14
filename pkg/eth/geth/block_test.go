package geth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"nft_standard/config"
	"testing"
)

func TestGetBlock(t *testing.T) {
	block, err := GetBlock(big.NewInt(11))
	if err != nil {
		t.Error(err)
		return
	}

	block.Difficulty()
	block.Extra()
	block.GasLimit()
	block.GasUsed()
	block.Hash()
	block.Bloom()
	block.Coinbase()  // miner
	block.MixDigest() // mixHash
	n := block.Nonce()
	fmt.Println(n)
	block.Number()
	block.ParentHash()
	block.ReceiptHash() // receiptsRoot
	block.UncleHash()   // sha3Uncles
	block.Size()
	block.Root() // stateRoot
	block.Time()
	block.Transactions()
	block.TxHash() // transactionRoot
	uncles := block.Uncles()
	fmt.Println(uncles)

	fmt.Println(block.Header().ParentHash.String())
	fmt.Println(block.Header().UncleHash.String())
	fmt.Println(block.Header().Coinbase.String())
	fmt.Println(block.Header().Root.String())
	fmt.Println(block.Header().TxHash.String())
	fmt.Println(block.Header().ReceiptHash.String())
	fmt.Println(block.Header().ParentHash.String())

	for _, tx := range block.Transactions() {

		tx.To()

		proIDFlag, _, _ := config.ProjectIDS.GetMin()
		chainID, err := proIDFlag.HttpClient.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
			fmt.Println(msg.From().Hex())
		}

		tx.Gas()
		tx.GasPrice()
		fmt.Println("tx =>", tx.Hash().String())
		tx.Nonce()
		tx.Value()
		tx.Data()
	}
}
