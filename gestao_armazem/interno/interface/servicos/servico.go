package servicos

type Servicos interface {
	ExpedirMercadoria(idproduto string, quantidade int)
}
