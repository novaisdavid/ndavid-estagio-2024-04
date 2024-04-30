package produto_stock

type Produto struct{
	Descicao string
	CondicaoArmazenamento string
	Categoria string
}

func (p Produto) Now() string{
	return ""
}
func (p *Produto) Cadastro_produto(d, ca, c string) {
	/*p.Descicao = "Sal_hymalaiano"
	p.CondicaoArmazenamento = "Ambiente_Frio"
	p.Categoria = "Alimento"*/
	
	p.Descicao = d
	p.CondicaoArmazenamento = ca
	p.Categoria = c
}