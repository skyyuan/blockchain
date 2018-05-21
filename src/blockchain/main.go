package main

import (
	_ "blockchain/routers"
	"github.com/astaxie/beego"
	"blockchain/blocks"
	"fmt"
	"strconv"
)

func main() {
	bc := blocks.NewBlockchain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blocks.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
	beego.Run()
}

