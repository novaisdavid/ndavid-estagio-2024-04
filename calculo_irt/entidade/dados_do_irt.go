package entidade

type DadosIrt struct {
	diasUteis          int
	falta              int
	subsiAlimentacao   float64
	subsidioTransporte float64
	premios            float64
	salarioBase        float64
}

func (d DadosIrt) NewDadosIrt(diasUteis, falta int, SubsiAlimentacao, SubsiTransporte, premios, salarioBase float64) DadosIrt {
	return DadosIrt{
		diasUteis:          diasUteis,
		falta:              falta,
		subsiAlimentacao:   SubsiAlimentacao,
		subsidioTransporte: SubsiTransporte,
		premios:            premios,
		salarioBase:        salarioBase,
	}
}

func (d DadosIrt) DiasUteis() int {
	return d.diasUteis
}

func (d DadosIrt) Falta() int {
	return d.falta
}

func (d DadosIrt) SubsidioAlimentacao() float64 {
	return d.subsiAlimentacao
}

func (d DadosIrt) SubsidioTransporte() float64 {
	return d.subsidioTransporte
}

func (d DadosIrt) Premios() float64 {
	return d.premios
}

func (d DadosIrt) SalarioBase() float64 {
	return d.salarioBase
}
