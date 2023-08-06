package halo

import (
	"encoding/hex"
	"errors"

	"github.com/ebfe/scard"

	"github.com/rozifus/silotools/pkg/access"
)


func ExecuteCommand(card *scard.Card, payload []byte) ([]byte, error) {
	opSelect, err := hex.DecodeString("00A4040007481199130e9f0100") // TODO: what is this magic hex?
	if err != nil {
		return nil, err
	}

	_, err = access.TransmitBytes(card, opSelect)
	if err != nil {
		return nil, err
	}

	opCommand := NewCommandRequestByteable(payload)

	resCommand, err := access.TransmitByteable(card, opCommand)
	if err != nil {
		return nil, err
	}

	return resCommand, nil
}

func Keys(card *scard.Card) (outKeys []string, err error) {
	opKeys, err := hex.DecodeString("02")

	resKeys, err := ExecuteCommand(card, opKeys)
	if err != nil {
		return nil, err
	}

	for len(resKeys) != 0 {
		keyLen := int(resKeys[0])
		rawKey := resKeys[1:keyLen+1]
		key := hex.EncodeToString(rawKey)

		outKeys = append(outKeys, key)

		resKeys = resKeys[keyLen+1:]
	}

	return
}

func Sign(card *scard.Card, keyNumber uint8, data []byte) ([]byte, error) {
	data = zeroPadBytes(data, 32)

	payload := NewSignRequestBytes(keyNumber, data)

	res, err := ExecuteCommand(card, payload)
	if err != nil {
		return nil, err
	}

	if res[0] == 0xe1 && len(res) == 2 {
		return nil, errors.New("Sign failed!")
	}

	return res, nil
}


func zeroPadBytes(bytes []byte, targetSize int) []byte {
	if len(bytes) >= targetSize {
		return bytes
	}

	padSize := targetSize - len(bytes)

	paddedBytes := make([]byte, targetSize)
	copy(paddedBytes[padSize:], bytes)
	return paddedBytes
}

