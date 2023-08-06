package byteable


type Tree struct {
	Name string
	Children []Byteable
}

func (t *Tree)GetName() string {
	return t.Name
}

func (t *Tree)ToBytes() (b []byte) {
	var out = make([]byte, 0, len(t.Children))
	for _, c := range t.Children {
		out = append(out, c.ToBytes()...)
	}
	return out
}

func NewTree(name string, bs ...Byteable) *Tree {
	return &Tree{name, bs}
}