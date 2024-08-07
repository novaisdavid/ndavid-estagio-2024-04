package entidade

type TabelaIrt struct {
	ValorInicial float64
	Limite       float64
	ParcelaFixa  float64
	Excesso      float64
	Taxa         float64
}

type CalculoIrt struct {
	tabelaIrt       []TabelaIrt
	diasUteisPadrao int
}

func (i CalculoIrt) NewIrt(tabela []TabelaIrt) CalculoIrt {
	return CalculoIrt{tabelaIrt: tabela, diasUteisPadrao: 22}
}

func (i CalculoIrt) CalcularSalariobruto(parame ...float64) float64 {
	salarioBruto := 0.0

	for _, valor := range parame {
		salarioBruto = salarioBruto + valor
	}
	return salarioBruto
}

func (i CalculoIrt) CalcularSubsidioMensal(diasUteis int, subsidioDiario float64) float64 {
	if diasUteis == 0 {
		diasUteis = i.diasUteisPadrao
	}
	return float64(diasUteis) * subsidioDiario
}

func (i CalculoIrt) CalcualrSalarioBaseApoisFalta(diasUteis, falta int, salarioBase float64) float64 {
	if diasUteis == 0 {
		diasUteis = i.diasUteisPadrao
	}
	return salarioBase - float64(falta)*(float64(salarioBase)/float64(diasUteis))
}

func (i CalculoIrt) CalcualrInssDoTrabalhador(salarioBruto float64) float64 {
	return salarioBruto * 0.03
}

func (i CalculoIrt) CalcularSujeicaoIrt(subsidioAlimetacoMensal, subATransporteMensal float64) float64 {
	excessoAlimentcao := i.calcularExcesso(subsidioAlimetacoMensal)
	excessoTransporte := i.calcularExcesso(subATransporteMensal)

	if (excessoAlimentcao + excessoTransporte) < 0 {
		return 0
	}

	return excessoAlimentcao + excessoTransporte
}

func (i CalculoIrt) calcularExcesso(subsidio float64) float64 {
	limite := 30000.0

	return subsidio - limite
}

func (i CalculoIrt) CalcularMateriaColetavel(paramentro ...float64) float64 {

	if paramentro == nil {
		return 0.0
	}

	return paramentro[0] + paramentro[1] - paramentro[2]

}

func (i CalculoIrt) CalcularIrt(materialColetavel float64) float64 {
	irt := 0.0
	for _, valor := range i.tabelaIrt {
		if materialColetavel >= valor.ValorInicial && materialColetavel <= valor.Limite {
			irt = valor.ParcelaFixa + (materialColetavel-valor.Excesso)*valor.Taxa
			return irt
		}

		if valor.ValorInicial >= 100000001 {
			irt = valor.ParcelaFixa + (materialColetavel-valor.Excesso)*valor.Taxa
			return irt
		}
	}

	return irt
}

func (i CalculoIrt) CalcularTotalDesconto(inss, irt float64) float64 {
	return inss + irt
}

func (i CalculoIrt) CalcularSalarioLiquido(salarioBruto, totalDesconto float64) float64 {
	return salarioBruto - totalDesconto
}
