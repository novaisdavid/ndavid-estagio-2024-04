package notarecebimentorepositorio

import (
	notarecebimento "github.com/acmllda/interno/dominio/agregados/notaRecebimento"
	"errors"
)

type NotaRecebimentoRepositorio struct {
	notasRecebimento map[string]*notarecebimento.NotaRecebimento
}

func (n NotaRecebimentoRepositorio) New() *NotaRecebimentoRepositorio {
	return &NotaRecebimentoRepositorio{notasRecebimento: make(map[string]*notarecebimento.NotaRecebimento)}

}

func (n *NotaRecebimentoRepositorio) CriarNotaRecebimento(nt *notarecebimento.NotaRecebimento) {
	n.notasRecebimento[nt.Id()] = nt
}

func (n *NotaRecebimentoRepositorio) RecuperarNotaRecebimento(id string) (*notarecebimento.NotaRecebimento, error) {
	nt, existe := n.notasRecebimento[id]
	if !existe {
		return nil, errors.New("nota de recebimento não encontrada")
	}
	return nt, nil
}

func (n *NotaRecebimentoRepositorio) RecuperarTodasNotasRecebimento() []*notarecebimento.NotaRecebimento {
	var todasNotasDeRecebimento []*notarecebimento.NotaRecebimento
	for _, notas := range n.notasRecebimento {
		todasNotasDeRecebimento = append(todasNotasDeRecebimento, notas)
	}
	return todasNotasDeRecebimento
}
