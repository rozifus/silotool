package silo

import (
	"crypto/elliptic"
	"crypto/ecdsa"
	"math/big"
)

func verifyCurve(key, hash, signature []byte) bool {
	x := new(big.Int).SetBytes(key[:32])
	y := new(big.Int).SetBytes(key[32:])

	pubKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X: x,
		Y: y,
	}

	r := new(big.Int).SetBytes(signature[:32])
	s := new(big.Int).SetBytes(signature[32:])

	return ecdsa.Verify(&pubKey, hash, r, s)
}
