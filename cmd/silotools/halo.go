package main

import (
	"fmt"
	"encoding/hex"

	"github.com/ebfe/scard"

	"github.com/rozifus/silotools/pkg/halo"
)

type HaloCmd struct {
	Keys HaloKeysCmd `cmd`
	Sign HaloSignCmd `cmd`
}

type HaloKeysCmd struct {}
type HaloSignCmd struct {
	Key uint8 `field`
	Data string `field`
}

func (cmd *HaloKeysCmd) Run(cliCtx *CliContext) error {
	cardCtx, err := scard.EstablishContext()
	if err != nil {
		return err
	}
	defer cardCtx.Release()

	card, err := InitializeCard(cardCtx)
	if err != nil {
		return err
	}
	defer card.Disconnect(scard.ResetCard)

	keys, err := halo.Keys(card)
	if err != nil {
		return err
	}

	for index, key := range keys {
		fmt.Printf("key%v: %s\n", index+1, key)
	}

	return nil
}

func (cmd *HaloSignCmd) Run(cliCtx *CliContext) error {
	cardCtx, err := scard.EstablishContext()
	if err != nil {
		return err
	}
	defer cardCtx.Release()

	card, err := InitializeCard(cardCtx)
	if err != nil {
		return err
	}
	defer card.Disconnect(scard.ResetCard)

	dataToSign, err := hex.DecodeString(cmd.Data)
	if err != nil {
		fmt.Println("Invalid hex data for signing:")
		fmt.Printf("%v\n", err)
		return err
	}

	res, err := halo.Sign(card, cmd.Key, dataToSign)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	fmt.Println(hex.EncodeToString(res))

	return nil
}
