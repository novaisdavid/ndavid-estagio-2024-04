package expedirMercadoria_test

import (
	"testing"

	"github.com/acmllda/interno/casosDeUso/expedirMercadoria"
	agregado "github.com/acmllda/interno/dominio/agregados"
	"github.com/acmllda/interno/dominio/entidades/lote"
	"github.com/acmllda/interno/dominio/entidades/produto"
)

func TestExpedirMercadoria(t *testing.T) {

	t.Run("se guia de remessa foi criada", func(t *testing.T) {
		// arrange
		pedidoID := "001"
		quantidade := 100

		// act
		guia := agregado.NewGuiaRemessa("g001", pedidoID, quantidade)
		// assert
		if guia.Id() == "" {
			t.Fail()
		}
	})

	t.Run("recuperar um agregado de guia remessa no repositorio", func(t *testing.T) {
		// arrange
		pedidoID := "001"
		quantidade := 100
		guia := agregado.NewGuiaRemessa("g001", pedidoID, quantidade)
		r := agregado.GuiaRemessaRepositorio{}.New()
		r.CriarGuia(guia)

		// act
		gr, err := r.RecuperarGuia("g001")

		// assert
		if err != nil || gr == nil {
			t.Fatalf("Erro ao recuperar guia de remessa: %v", err)
		}
	})

	t.Run("se criou mais de um agregado de guia remessa", func(t *testing.T) {
		// arrange
		var guias []agregado.GuiaRemessa

		guias = append(guias, *agregado.NewGuiaRemessa("g001", "pedido123", 10))
		guias = append(guias, *agregado.NewGuiaRemessa("g002", "pedido456", 20))
		guias = append(guias, *agregado.NewGuiaRemessa("g003", "pedido789", 30))

		r := agregado.GuiaRemessaRepositorio{}.New()
		for _, g := range guias {
			r.CriarGuia(&g)
		}

		// act
		tr := r.RecuperarTodasGuias()

		//assert
		if len(tr) == 0 {
			t.Fail()
		}

	})

	t.Run("se o produto for encontrado no lote", func(t *testing.T) {
		// arrange
		produtoID := "003"
		quantidade := 100
		guia := agregado.NewGuiaRemessa("g001", produtoID, quantidade)
		l := lote.New("", produto.Produto{}, 0, "")

		//act
		p := l.EncontraProdutoEmUmOuMaisLotes(guia.ProdutoID())

		// assert
		if len(p) == 0 {
			t.Fail()
		}

	})

	t.Run("retorna produtos em lotes com data validade mais próxima", func(t *testing.T) {
		// arrange
		pedidoID := "003"
		quantidade := 30
		guia := agregado.NewGuiaRemessa("g001", pedidoID, quantidade)
		l := lote.New("", produto.Produto{}, 0, "")
		//act
		lotes := l.EncontraProdutoEmUmOuMaisLotes(guia.ProdutoID())

		// assert
		if len(lotes) == 0 {
			t.Fail()
		}

	})

	t.Run("Não retornar produtos em lotes com data validade já passada", func(t *testing.T) {
		// arrange
		pedidoID := "001"
		quantidade := 30
		guia := agregado.NewGuiaRemessa("g001", pedidoID, quantidade)
		l := lote.New("", produto.Produto{}, 0, "")

		//act
		lotes := l.EncontraProdutoEmUmOuMaisLotes(guia.ProdutoID())

		// assert
		if len(lotes) > 0 {
			t.Fail()
		}

	})

	t.Run("buscar um produto que não existe em lotes", func(t *testing.T) {
		// arrange
		pedidoID := "103"
		quantidade := 0
		guia := agregado.NewGuiaRemessa("g001", pedidoID, quantidade)
		l := lote.New("", produto.Produto{}, 0, "")

		//act
		lotes := l.EncontraProdutoEmUmOuMaisLotes(guia.ProdutoID())

		// assert
		if len(lotes) > 0 {
			t.Fail()
		}

	})

	t.Run("encontrar produto em lotes sem dados da guia de remessa", func(t *testing.T) {
		// arrange
		guia := agregado.NewGuiaRemessa("", "", 0)
		l := lote.New("", produto.Produto{}, 0, "")

		//act
		lotes := l.EncontraProdutoEmUmOuMaisLotes(guia.ProdutoID())

		// assert
		if len(lotes) > 0 {
			t.Fail()
		}

	})

	t.Run("retirar uma determinada quantidade de um produto em lote", func(t *testing.T) {
		// arrange
		produtoID := "003"
		quantidade := 1200

		guia := agregado.NewGuiaRemessa("g045", produtoID, quantidade)
		l := lote.New("", produto.Produto{}, 0, "")
		lotes := l.EncontraProdutoEmUmOuMaisLotes(guia.ProdutoID())
		//act
		lote := l.RetirarProdutoNoLote(lotes, guia.Quantidade())

		// assert
		if lote == -1 {
			t.Fail()
		}

	})

	t.Run("retirar quantidade acima de um produto mais do que existe no lote", func(t *testing.T) {
		// arrange
		produtoID := "003"
		quantidade := 2900

		guia := agregado.NewGuiaRemessa("g045", produtoID, quantidade)
		l := lote.New("", produto.Produto{}, 0, "")
		lotes := l.EncontraProdutoEmUmOuMaisLotes(guia.ProdutoID())
		//act
		lote := l.RetirarProdutoNoLote(lotes, guia.Quantidade())

		// assert
		if lote != -1 {
			t.Fail()
		}

	})

	t.Run("retirar um produto que não existe no lote", func(t *testing.T) {
		// arrange
		produtoID := "103"
		quantidade := 900

		guia := agregado.NewGuiaRemessa("g045", produtoID, quantidade)
		l := lote.New("", produto.Produto{}, 0, "")
		lotes := l.EncontraProdutoEmUmOuMaisLotes(guia.ProdutoID())
		//act
		lote := l.RetirarProdutoNoLote(lotes, guia.Quantidade())

		// assert
		if lote != -1 {
			t.Fail()
		}

	})

	t.Run("retirar um produto sem dados da guia de remessa", func(t *testing.T) {
		// arrange
		quantidade := 199
		guia := agregado.NewGuiaRemessa("", "", 0)
		l := lote.New("", produto.Produto{}, 0, "")
		lotes := l.EncontraProdutoEmUmOuMaisLotes(guia.ProdutoID())
		//act
		lote := l.RetirarProdutoNoLote(lotes, quantidade)

		// assert
		if lote != -1 {
			t.Fail()
		}

	})

	t.Run("expedir mercadoria", func(t *testing.T) {
		// arrange
		expedir := expedirMercadoria.ExpedirMercadoria{}
		idProduto := "003"
		quantidade := 200
		g := agregado.NewGuiaRemessa("g002", idProduto, quantidade)

		// act
		r := expedir.Executa(g)

		// assert
		if !r {
			t.Fail()
		}

	})

	t.Run("expedir uma quantidade de mercadoria mais do que a existente no/s lote/s", func(t *testing.T) {
		// arrange
		expedir := expedirMercadoria.ExpedirMercadoria{}
		idProduto := "003"
		quantidade := 20000
		g := agregado.NewGuiaRemessa("g002", idProduto, quantidade)

		// act
		resutado := expedir.Executa(g)

		// assert
		if resutado {
			t.Fail()
		}

	})

}
