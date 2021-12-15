package silo

import (
	"encoding/hex"
)


type CoreData struct {
	HardwareRevisionNumber,
	FirmwareNumber,
	SerialNumber,
	ExternalPublicKey,
	InternalPublicKey,
	SmartContractAddress,
	NxpI2cSerial,
	Nxp804Serial,
	Atecc608aSerial,
	ConfigZoneBytes []byte
}

func NewCoreData(b []byte) *CoreData {
	return &CoreData{
		HardwareRevisionNumber: b[70:74],
		FirmwareNumber: b[74:78],
		SerialNumber: b[78:86],
		ExternalPublicKey: b[86:150],
		InternalPublicKey: b[150:214],
		SmartContractAddress: b[214:234],
		NxpI2cSerial: b[234:241],
		Nxp804Serial: b[241:257],
		Atecc608aSerial: b[257:266],
		ConfigZoneBytes: b[266:394],
	}
}

func NewSignatureRequestBytes(commandCode byte, address, block []byte) (b []byte) {
	paddedAddress := make([]byte, 32)
	copy(paddedAddress[:20], address)

	combinedHash := createHash(address, block)
	checksum := createChecksum([]byte{commandCode}, paddedAddress, block, combinedHash)

	notSureWhatThisIs, _ := hex.DecodeString("550063")
	endThing, _ := hex.DecodeString("fe00")

	b = append(b, notSureWhatThisIs...)
	b = append(b, commandCode)
	b = append(b, paddedAddress...)
	b = append(b, block...)
	b = append(b, combinedHash...)
	b = append(b, checksum...)
	b = append(b, endThing...)
	return
}


type SignatureResult struct {
	LastHash,
	ExternalSignature,
	InternalSignature,
	Counter []byte
}

func NewSignatureResult(b []byte) *SignatureResult {
	return &SignatureResult{
		LastHash: b[65:98],
		ExternalSignature: b[129:193],
		InternalSignature: b[193:257],
		Counter: b[257:258],
	}
}
