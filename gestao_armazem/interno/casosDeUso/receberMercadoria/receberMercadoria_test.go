package recebermercadoria_test

import (
	"testing"

	recebermercadoria "github.com/acmllda/interno/casosDeUso/receberMercadoria"
	notarecebimento "github.com/acmllda/interno/dominio/agregados"
	"github.com/acmllda/interno/dominio/entidades/lote"
	"github.com/acmllda/interno/dominio/entidades/produto"
)

func TestReceberMercadoria(t *testing.T) {
	t.Run("se nota de recebimento foi criada", func(t *testing.T) {
		//arrange

		// act
		n := notarecebimento.NotaRecebimento{}.New("", "", 0, "")

		//assert
		if n.Quantidade() != 0 {
			t.Fail()
		}
	})

	t.Run("se há dados na nota de recebimento", func(t *testing.T) {
		//arrange
		id := "n001"
		Quantidade := 12
		validade := "2024-07-21"
		// act
		n := notarecebimento.NotaRecebimento{}.New(id, "", Quantidade, validade)

		//assert
		if n.Validade() == "" {
			t.Fail()
		}
	})

	t.Run("se há produto na nota de recebimento", func(t *testing.T) {
		//arrange
		produtoID := "013"
		id := "n001"
		Quantidade := 12
		validade := "2024-07-21"
		// act
		n := notarecebimento.NotaRecebimento{}.New(id, produtoID, Quantidade, validade)

		//assert
		if n.ProdutoID() == "" {
			t.Fail()
		}
	})

	t.Run("recuperar agregado no repositorio da nota de recebimento", func(t *testing.T) {
		//arrange
		produtoID := "013"
		id := "n001"
		Quantidade := 12
		validade := "2024-07-21"

		n := notarecebimento.NotaRecebimento{}.New(id, produtoID, Quantidade, validade)
		repo := notarecebimento.NotaRecebimentoRepositorio{}.New()
		repo.CriarNotaRecebimento(n)

		// act
		rp, err := repo.RecuperarNotaRecebimento(id)

		//assert
		if err != nil || rp == nil {

			t.Fail()
		}

	})

	t.Run("se criou mais de um agregado no repositorio da nota de recebimento", func(t *testing.T) {
		//arrange
		var notasRecebimento []notarecebimento.NotaRecebimento
		repo := notarecebimento.NotaRecebimentoRepositorio{}.New()

		notasRecebimento = append(notasRecebimento, *notarecebimento.NotaRecebimento{}.New("n001", "005", 1432, "2025-03-21"))
		notasRecebimento = append(notasRecebimento, *notarecebimento.NotaRecebimento{}.New("n001", "002", 1672, "2025-02-11"))
		notasRecebimento = append(notasRecebimento, *notarecebimento.NotaRecebimento{}.New("n001", "022", 112, "2025-04-01"))
		notasRecebimento = append(notasRecebimento, *notarecebimento.NotaRecebimento{}.New("n001", "202", 120, "2025-05-22"))

		for _, nota := range notasRecebimento {
			repo.CriarNotaRecebimento(&nota)
		}

		// act
		rp := repo.RecuperarTodasNotasRecebimento()

		//assert
		if len(rp) == 0 {
			t.Fail()
		}

	})

	t.Run("nota de recebimento sem dado", func(t *testing.T) {
		//arrange

		//act
		n := notarecebimento.NotaRecebimento{}.New("", "", 0, "")

		// assert
		if n.Quantidade() > 0 {
			//fmt.Println("NOTA: ", n)
			t.Fail()
		}

	})

	t.Run("verificar se há produto no lote", func(t *testing.T) {
		//arrange
		var lot *lote.Lote
		repo := notarecebimento.NotaRecebimentoRepositorio{}.New()
		produto := produto.New("23", "arroz", "2024-08-12")
		quantidade := 100
		nota := notarecebimento.NotaRecebimento{}.New("nt003", produto.Id(), quantidade, produto.DataValidade())
		repo.CriarNotaRecebimento(nota)

		//act
		if tem, _ := repo.RecuperarNotaRecebimento(nota.Id()); tem.Id() != "" {
			lot = lote.New("lt102", *produto, quantidade, produto.DataValidade())

		}

		//assert
		if lot.ProdutoId() == "" {
			t.Fail()
		}
	})

	t.Run("verificar se os dados do produto que está na nota de recebimento é igual com os dados do produto que está no lote", func(t *testing.T) {
		//arrange
		dadosIguais := false
		repo := notarecebimento.NotaRecebimentoRepositorio{}.New()
		produto := produto.New("23", "arroz", "2024-08-12")
		quantidade := 100
		nota := notarecebimento.NotaRecebimento{}.New("nt003", produto.Id(), quantidade, produto.DataValidade())
		repo.CriarNotaRecebimento(nota)
		lot := lote.New("lt102", *produto, quantidade, produto.DataValidade())
		//act
		if tem, _ := repo.RecuperarNotaRecebimento(nota.Id()); tem.Id() != "" && tem.ProdutoID() == lot.ProdutoId() && tem.Quantidade() == lot.Quantidade() {
			dadosIguais = true

		}

		//assert
		if !dadosIguais {
			t.Fail()
		}
	})

	t.Run("se recebeu a mercadoria e a nota de remessa foi criada", func(t *testing.T) {
		//arrange
		recebermercadori := recebermercadoria.ReceberMercadoria{}
		produto := produto.New("23", "arroz", "2024-08-12")
		quantidade := 100
		nota := notarecebimento.NotaRecebimento{}.New("nt003", produto.Id(), quantidade, produto.DataValidade())

		//act
		d := recebermercadori.Receber(nota)

		// assert

		if !d {
			t.Fail()
		}
	})

}
