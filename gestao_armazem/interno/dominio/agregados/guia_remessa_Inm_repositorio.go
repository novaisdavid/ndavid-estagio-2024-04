package agregados

import (
	"errors"
)

type GuiaRemessaRepositorio struct {
	guias map[string]*GuiaRemessa
}

func (r GuiaRemessaRepositorio) New() *GuiaRemessaRepositorio {
	return &GuiaRemessaRepositorio{guias: make(map[string]*GuiaRemessa)}
}

func (r *GuiaRemessaRepositorio) CriarGuia(g *GuiaRemessa) {
	r.guias[g.Id()] = g
}

func (r *GuiaRemessaRepositorio) RecuperarGuia(id string) (*GuiaRemessa, error) {
	g, existe := r.guias[id]
	if !existe {
		return nil, errors.New("guia de remessa não encontrada")
	}
	return g, nil
}

func (r *GuiaRemessaRepositorio) RecuperarTodasGuias() []*GuiaRemessa {
	var todasGuias []*GuiaRemessa
	for _, guia := range r.guias {
		todasGuias = append(todasGuias, guia)
	}
	return todasGuias
}
