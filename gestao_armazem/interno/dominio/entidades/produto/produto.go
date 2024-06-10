package produto

type Produto struct {
	id           string
	descricao    string
	dataValidade string
}

func New(i, dc, d string) *Produto {
	return &Produto{
		id:           i,
		descricao:    dc,
		dataValidade: d,
	}
}

func (p Produto) Id() string {
	return p.id
}

func (p Produto) DataValidade() string {
	return p.dataValidade
}
