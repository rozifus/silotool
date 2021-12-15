package main

import (
	"github.com/ebfe/scard"
)

func InitializeCard(cardCtx *scard.Context) (*scard.Card, error) {
	readers, err := cardCtx.ListReaders()
	if err != nil {
		return nil, err
	}

	index, err := waitUntilCardPresent(cardCtx, readers)
	if err != nil {
		return nil, err
	}

	card, err := cardCtx.Connect(readers[index], scard.ShareExclusive, scard.ProtocolAny)
	if err != nil {
		return nil, err
	}

	_, err = card.Status()
	if err != nil {
		return nil, err
	}

	return card, nil
}

func waitUntilCardPresent(ctx *scard.Context, readers []string) (int, error) {
	rs := make([]scard.ReaderState, len(readers))
	for i := range rs {
		rs[i].Reader = readers[i]
		rs[i].CurrentState = scard.StateUnaware
	}

	for {
		for i := range rs {
			if rs[i].EventState&scard.StatePresent != 0 {
				return i, nil
			}
			rs[i].CurrentState = rs[i].EventState
		}
		err := ctx.GetStatusChange(rs, -1)
		if err != nil {
			return -1, err
		}
	}
}