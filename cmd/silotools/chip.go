package main

import (
	"fmt"

	"github.com/ebfe/scard"

	"github.com/rozifus/silotools/pkg/chip"
)

type ChipCmd struct {
	Uid ChipUidCmd `cmd`
}

type ChipUidCmd struct {}

func (cmd *ChipUidCmd) Run(cliCtx *CliContext) error {
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

	res, err := chip.GetUid(card)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
