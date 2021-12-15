package access


type Header struct {
	Class, Ins, P1, P2, Length byte
}

func (he *Header) Bytes() []byte {
	return []byte{
		byte(he.Class),
		byte(he.Ins),
		byte(he.P1),
		byte(he.P2),
		byte(he.Length),
	}
}
