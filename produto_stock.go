package produto_stock

type Produto struct {
	IdProduto      string
	NomeDoProduto  string
	Categoria      string
	DataDeProducao string
	DataDeValidade string
}

func (p Produto) VerificaSeProdutoExiste(produtoEsperado []Produto, IdProduto string) Produto {
	for _, produto := range produtoEsperado {
		if produto.IdProduto == IdProduto {
			return produto
		}
	}
	return Produto{}
}
