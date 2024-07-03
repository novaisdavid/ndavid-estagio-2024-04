package calculodoirt

type Irt struct {
}

func (i Irt) NewIrt() Irt {
	return Irt{}
}

func (i Irt) CalcularSalariobruto(parame ...float64) float64 {
	salarioBruto := 0.0

	for _, valor := range parame {
		salarioBruto = salarioBruto + valor
	}

	return salarioBruto
}

func (i Irt) CalcularSubsidioMensal(diasUteis int, subsidioDiario float64) float64 {

	diasUteisPadrao := 22

	if diasUteis > 0 {
		diasUteisPadrao = diasUteis
	}

	return float64(diasUteisPadrao) * subsidioDiario
}

