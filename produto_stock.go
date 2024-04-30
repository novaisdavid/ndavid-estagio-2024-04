package produto_stock

type Produto struct{
	Descricao             string
	CondicaoArmazenamento string
	Categoria             string
}


func New() string {
	return "q"
}

func (p *Produto ) CadastroProduto(d, ca, cat string) int{
	p.Descricao = d
	p.CondicaoArmazenamento = ca
	p.Categoria = cat

	return 1	
}
