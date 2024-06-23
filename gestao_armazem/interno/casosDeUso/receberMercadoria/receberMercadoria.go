package recebermercadoria

import (
	notarecebimento "github.com/acmllda/interno/dominio/agregados/notaRecebimento"
	notarecebimentorepositorio "github.com/acmllda/interno/dominio/repositorio/notaRecebimentoRepositorio"
)

type ReceberMercadoria struct {
	notaRepositorio notarecebimentorepositorio.NotaRecebimentoRepositorio
}

func (r ReceberMercadoria) Receber(n *notarecebimento.NotaRecebimento) bool {

	nota := r.notaRepositorio.New()
	nota.CriarNotaRecebimento(n)

	if existe, _ := nota.RecuperarNotaRecebimento(n.Id()); existe.Id() != "" {
		return true
	}

	return false
}
