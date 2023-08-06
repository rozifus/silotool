package halo

import (
	"encoding/hex"

	"github.com/rozifus/silotools/pkg/byteable"
)


func NewCommandRequestBytes(payload []byte) (b []byte) {
	starterThing, _ := hex.DecodeString("B0510000")
	enderThing, _ := hex.DecodeString("00")

	payloadSize := byte(len(payload))

	b = append(b, starterThing...)
	b = append(b, payloadSize)
	b = append(b, payload...)
	b = append(b, enderThing...)
	return
}

func NewCommandRequestByteable(payload []byte) (b byteable.Byteable) {
	return byteable.NewTree("command-request",
		byteable.NewNodeFromHex("starter-thing", "B0510000"),
		byteable.NewNode("payload size", byte(len(payload))),
		byteable.NewNode("payload", payload...),
		byteable.NewNodeFromHex("ender-thing", "00"),
	)
}

func NewSignRequestBytes(keyNumber uint8, data []byte) (b []byte) {
	innerStart, _ := hex.DecodeString("01")

	b = append(b, innerStart...)
	b = append(b, keyNumber)
	b = append(b, data...)
	return b
}

