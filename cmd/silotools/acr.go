package main

import (
	"github.com/ebfe/scard"

	"github.com/rozifus/silotools/pkg/access"
)

type AcrCmd struct {
	Buzzer AcrBuzzerCmd `cmd`
}

type AcrBuzzerCmd struct {
	Enable AcrBuzzerEnableCmd `cmd`
	Disable AcrBuzzerDisableCmd `cmd`
	Status AcrBuzzerStatusCmd `cmd`
}

type AcrBuzzerEnableCmd struct {}
type AcrBuzzerDisableCmd struct {}
type AcrBuzzerStatusCmd struct {}

func (cmd *AcrBuzzerEnableCmd) Run(cliCtx *CliContext) error {
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

	return setBuzzerEnabled(card, true)
}

func (cmd *AcrBuzzerDisableCmd) Run(cliCtx *CliContext) error {
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

	return setBuzzerEnabled(card, false)
}

func setBuzzerEnabled(card *scard.Card, enabled bool) error {
	// from https://stackoverflow.com/a/41550221

	var buzzerTargetState byte = 0x00
	if enabled {
		buzzerTargetState = 0xff
	}

	h := access.Header{
		0xff,
		0x00,
		0x52,
		buzzerTargetState,
		0x00,
	}

	_, err := access.DoTransmit(card, h.Bytes())
	return err
}

