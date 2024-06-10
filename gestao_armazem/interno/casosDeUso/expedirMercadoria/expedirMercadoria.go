package expedirMercadoria

import (
	guiaremessa "ACMELDA/interno/dominio/agregados/guiaRemessa"
	"ACMELDA/interno/dominio/entidades/lote"
	guiaRemessaRepositorio "ACMELDA/interno/dominio/repositorio/guiaRemessaRepositorio"
)

type ExpedirMercadoria struct {
	lote            lote.Lote
	guiaRepositorio guiaRemessaRepositorio.GuiaRemessaRepositorio
}

func (e ExpedirMercadoria) Executa(g *guiaremessa.GuiaRemessa) bool {
	lotes := e.lote.EncontraProdutoEmUmOuMaisLotes(g.ProdutoID())
	retirou := e.lote.RetirarProdutoNoLote(lotes, g.Quantidade())
	guia := e.guiaRepositorio.New()

	if retirou != -1 {
		guia.CriarGuia(g)
		return true

	}

	return false
}
