package guiaRemessa

import (
	"time"
)

type GuiaRemessa struct {
	id            string
	produtoID     string
	quantidade    int
	dataExpedicao time.Time
}

func NewGuiaRemessa(i, pi string, q int) GuiaRemessa {
	return GuiaRemessa{
		id:            i,
		produtoID:     pi,
		quantidade:    q,
		dataExpedicao: time.Now(),
	}
}

func (g *GuiaRemessa) Id() string {
	return g.id
}

func (g *GuiaRemessa) ProdutoID() string {
	return g.produtoID
}

func (g *GuiaRemessa) DataExpedicao() time.Time {
	return g.dataExpedicao
}

func (g *GuiaRemessa) Quantidade() int {
	return g.quantidade
}
