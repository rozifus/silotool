package silo

import (
	"time"

	"github.com/ebfe/scard"

	"github.com/rozifus/silotools/pkg/access"
)


func ReadCore(card *scard.Card) (*CoreData, error) {
	coreBytes, err := access.BinaryReadBlocks(card, 0x0, 0x63)
	if err != nil {
		return nil, err
	}

	return NewCoreData(coreBytes), nil
}

func TestSignature(card *scard.Card) (bool, error) {
	coreData, err := ReadCore(card)
	if err != nil {
		return false, err
	}

	time.Sleep(200 * time.Millisecond)

	address, err := randomBytes(20)
	if err != nil {
		return false, err
	}

	block, err := randomBytes(32)
	if err != nil {
		return false, err
	}

	sigRequestBytes := NewSignatureRequestBytes(0x00, address, block)

	_, err = access.BinaryWriteBlocks(card, 0xb0, sigRequestBytes)
	if err != nil {
		return false, err
	}

	time.Sleep(2500 * time.Millisecond)

	sigResultBytes, err := access.BinaryReadBlocks(card, 0x64, 0xa5)
	if err != nil {
		return false, err
	}

	sr := NewSignatureResult(sigResultBytes)

	combinedHash := createHash(address, block)

	return verifyCurve(coreData.ExternalPublicKey, combinedHash, sr.ExternalSignature), nil
}