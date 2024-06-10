package guiaRemessaRepository

import (
	"ACMELDA/interno/dominio/agregados/guiaRemessa"
	"errors"
)

type GuiaRemessaRepositorio struct {
	guias map[string]*guiaRemessa.GuiaRemessa
}

func (r GuiaRemessaRepositorio) New() *GuiaRemessaRepositorio {
	return &GuiaRemessaRepositorio{guias: make(map[string]*guiaRemessa.GuiaRemessa)}
}

func (r *GuiaRemessaRepositorio) CriarGuia(g *guiaRemessa.GuiaRemessa) {
	r.guias[g.Id()] = g
}

func (r *GuiaRemessaRepositorio) RecuperarGuia(id string) (*guiaRemessa.GuiaRemessa, error) {
	g, existe := r.guias[id]
	if !existe {
		return nil, errors.New("guia de remessa não encontrada")
	}
	return g, nil
}

func (r *GuiaRemessaRepositorio) RecuperarTodasGuias() []*guiaRemessa.GuiaRemessa {
	var todasGuias []*guiaRemessa.GuiaRemessa
	for _, guia := range r.guias {
		todasGuias = append(todasGuias, guia)
	}
	return todasGuias
}
