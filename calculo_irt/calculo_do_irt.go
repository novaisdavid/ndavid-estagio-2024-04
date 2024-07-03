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

func (i Irt) CalcualrSalarioBaseApoisFalta(diasUteis, falta int, salarioBase float64) float64 {
	diasUteisPadrao := 22

	if diasUteis > 0 {
		diasUteisPadrao = diasUteis
	}

	return salarioBase - float64(falta)*(float64(salarioBase)/float64(diasUteisPadrao))
}

func (i Irt) CalcualrInssDoTrabalhador(salarioBase float64) float64 {
	return salarioBase * 0.03
}

func (i Irt) CalcularSujeicaoIrt(subsidioAlimetacoMensal, subATransporteMensal float64) float64 {
	excessoAlimentcao := i.calcularExcesso(subsidioAlimetacoMensal)
	excessoTransporte := i.calcularExcesso(subATransporteMensal)

	if (excessoAlimentcao + excessoTransporte) < 0 {
		return 0
	}
	return excessoAlimentcao + excessoTransporte
}

func (i Irt) calcularExcesso(subsidio float64) float64 {
	limite := 30000.0

	return subsidio - limite
}

func (i Irt) CalcularMaterialColetavel(paramentro ...float64) float64 {

	if paramentro == nil {
		return 0.0
	}

	return paramentro[0] + paramentro[1] - paramentro[2]

}
