package recebermercadoria

import (
	notarecebimento "ACMELDA/interno/dominio/agregados/notaRecebimento"
	notarecebimentorepositorio "ACMELDA/interno/dominio/repositorio/notaRecebimentoRepositorio"
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
