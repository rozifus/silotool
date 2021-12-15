package access

import (
	"errors"
	"encoding/binary"
	"time"

	"github.com/ebfe/scard"
)


const(
	BinaryRead uint8 = 0xb0
	BinaryWrite = 0xd6
)

const TransmitSuccess uint16 = 0x9000


func BinaryReadBlocks(card *scard.Card, startBlock, endBlock uint16) ([]byte, error) {
	var blockCount = endBlock - startBlock

	result := make([]byte, 4 * blockCount)

	d := 0
	for i:=startBlock; i<endBlock; i++ {
		b, err := BinaryReadBlock(card, i)
		if err != nil {
			return nil, err
		}

		copy(result[d:d+4], b)
		d += 4
	}

	return result, nil
}

func BinaryWriteBlocks(card *scard.Card, block uint8, data []byte) ([]byte, error) {
	if !(len(data) % 4 == 0) {
		return nil, errors.New("data must be multiple of 4 bytes")
	}

	for i:=0; i<len(data); i+=4 {
		err := BinaryWriteBlock(card, block, data[i:i+4])
		if err != nil {
			return nil, err
		}
		block++
	}

	time.Sleep(200 * time.Millisecond)

	return BinaryReadBlock(card, 0xcb)
}

func BinaryReadBlock(card *scard.Card, block uint16) ([]byte, error) {
	h := Header{
		0xff,
		BinaryRead,
		uint8(block >> 8),
		uint8(block),
		4,
	}

	result, err := DoTransmit(card, h.Bytes())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func BinaryWriteBlock(card *scard.Card, block uint8, data []byte) error {
	if (len(data) != 4) {
		return errors.New("data must be 4 bytes")
	}

	payload := make([]byte, 9)

	h := Header{
		0xff,
		BinaryWrite,
		0x00,
		block,
		4,
	}

	copy(payload[:5], h.Bytes())
	copy(payload[5:], data)

	_, err := DoTransmit(card, payload)
	return err
}

func DoTransmit(card *scard.Card, payload []byte) (rdata []byte, err error) {
	rc, err := card.Transmit(payload)
	if err != nil {
		return nil, err
	}

	result := rc[:len(rc)-2]
	code := binary.BigEndian.Uint16(rc[len(rc)-2:])

	if code != TransmitSuccess {
		return nil, errors.New("Transmit Failed")
	}

	return result, nil
}