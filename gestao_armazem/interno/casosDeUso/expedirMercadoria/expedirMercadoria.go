package expedirMercadoria

import (
	guiaremessa "github.com/acmllda/interno/dominio/agregados/guiaRemessa"
	"github.com/acmllda/interno/dominio/entidades/lote"
	guiaRemessaRepositorio "github.com/acmllda/interno/dominio/repositorio/guiaRemessaRepositorio"
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
