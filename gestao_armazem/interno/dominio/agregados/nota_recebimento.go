package notarecebimento

type NotaRecebimento struct {
	id         string
	produtoId  string
	quantidade int
	validade   string
}

func New(i, pi string, q int, v string) NotaRecebimento {
	return NotaRecebimento{id: i, produtoId: pi, quantidade: q, validade: v}
}

func (n NotaRecebimento) Id() string {
	return n.id
}

func (n NotaRecebimento) Quantidade() int {
	return n.quantidade
}

func (n NotaRecebimento) Validade() string {
	return n.validade
}

func (n NotaRecebimento) ProdutoID() string {
	return n.produtoId
}
