package byteable


type Byteable interface {
	ToBytes() []byte
	GetName() string
}