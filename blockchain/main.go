package main

// uninstalled packages will be installed by the go.mod
import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	// hash is the hash value of this block
	Hash []byte
	// data is the data/information on this block
	Data []byte
	// is the hash value of the previous block that's linked to
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	// joining together the slices of bytes into info
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// creating the hash of our info with contains our data and prevHash (the secret part)
	hash := sha256.Sum256(info)
	// giving our block the hash value we just created
	b.Hash = hash[:]
}

func createBlock(data string, prevHash []byte) *Block {
	// creating a block with empty hash and data, precHash
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) addBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := createBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return createBlock("Genesis", []byte{})
}

func initBlockChain() *BlockChain {
	// returning our block chain which is an array that contain
	// our genesis block
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := initBlockChain()

	chain.addBlock("First after genesis")
	chain.addBlock("Second after genesis")
	chain.addBlock("Third after genesis")

	for _, block := range chain.blocks {
		//fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		// %s is just a string interpolation
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
