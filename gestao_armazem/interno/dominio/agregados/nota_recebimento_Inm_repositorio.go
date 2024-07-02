package agregados

import (
	"errors"
)

type NotaRecebimentoRepositorio struct {
	notasRecebimento map[string]*NotaRecebimento
}

func (n NotaRecebimentoRepositorio) New() *NotaRecebimentoRepositorio {
	return &NotaRecebimentoRepositorio{notasRecebimento: make(map[string]*NotaRecebimento)}

}

func (n *NotaRecebimentoRepositorio) CriarNotaRecebimento(nt *NotaRecebimento) {
	n.notasRecebimento[nt.Id()] = nt
}

func (n *NotaRecebimentoRepositorio) RecuperarNotaRecebimento(id string) (*NotaRecebimento, error) {
	nt, existe := n.notasRecebimento[id]
	if !existe {
		return nil, errors.New("nota de recebimento não encontrada")
	}
	return nt, nil
}

func (n *NotaRecebimentoRepositorio) RecuperarTodasNotasRecebimento() []*NotaRecebimento {
	var todasNotasDeRecebimento []*NotaRecebimento
	for _, notas := range n.notasRecebimento {
		todasNotasDeRecebimento = append(todasNotasDeRecebimento, notas)
	}
	return todasNotasDeRecebimento
}
