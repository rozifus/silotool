package byteable

import (
	"encoding/hex"
)


type Node struct {
	Name string
	Bytes []byte
}

func (bn *Node)GetName() string {
	return bn.Name
}

func (bn *Node)ToBytes() []byte {
	return bn.Bytes
}

func NewNode(s string, bs ...byte) *Node {
	return &Node{s, bs}
}

func NewNodeFromHex(s string, h string) *Node {
	bs, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}

	return NewNode(s, bs...)
}