package chip

import (
	"encoding/hex"
	"github.com/ebfe/scard"

	"github.com/rozifus/silotools/pkg/access"
)

func GetUid(card *scard.Card) ([]byte, error) {
	payload, err := hex.DecodeString("ffca000000")
	if err != nil {
		return nil, err
	}

	res, err := access.DoTransmit(card, payload)
	if err != nil {
		return nil, err
	}

	return res, nil
}




