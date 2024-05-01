package produto_stock

type Produto struct{
	Descricao             string
	CondicaoArmazenamento string
	Categoria             string
}


func (p *Produto ) CadastroProduto(d, ca, cat string){
	p.Descricao = d
	p.CondicaoArmazenamento = ca
	p.Categoria = cat

}

func (p Produto) VerProdutoCadastrado() string {
	return p.Descricao +" "+p.CondicaoArmazenamento+" "+p.Categoria
}
