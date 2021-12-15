package silo

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"reflect"

	"github.com/sigurn/crc16"
)

func randomBytes(amount uint) ([]byte, error) {
	b := make([]byte, amount)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func bytesFromStructOfBytes(stct interface{}) (b []byte) {
	v := reflect.ValueOf(stct)

	for i:=0; i<v.NumField(); i++ {
		b = append(b, v.Field(i).Interface().([]byte)...)
	}

	return
}

func createHash(components... []byte) []byte {
	s := sha256.New()
	for _, m := range components {
		s.Write(m)
	}
	return s.Sum(nil)
}

func createChecksum(components... []byte) []byte {
	table := crc16.MakeTable(crc16.CRC16_CCITT_FALSE)

	c := crc16.New(table)
	for _, m := range components {
		c.Write(m)
	}
	csum := c.Sum16()

	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, csum)

	return b
}