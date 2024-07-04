package calculodoirt_test

import (
	"testing"

	calculo "github.com/calculo_irt"
)

func tabelaIrt() []calculo.TabelaIrt {
	return []calculo.TabelaIrt{
		{0, 100000, 0, 0, 0},
		{1001000, 150000, 0, 100001, 0.13},
		{150001, 200000, 12500, 150001, 0.16},
		{200001, 300000, 31250, 200001, 0.18},
		{300001, 500000, 49250, 300001, 0.19},
		{500001, 1000000, 87250, 500001, 0.2},
		{1000001, 1500000, 187249, 1000001, 0.21},
		{1500001, 2000000, 292249, 1500001, 0.22},
		{2000001, 2500000, 402249, 2000001, 0.23},
		{2500001, 5000000, 517249, 2500001, 0.24},
		{5000001, 10000000, 1117249, 5000001, 0.245},
		{10000001, 0, 2342248, 10000001, 0.25},
	}
}

var tabela = tabelaIrt()
var s = calculo.CalculoIrt{}.NewIrt(tabela)
var diasUteis = 30
var salarioBase = 30000.0
var premios = 5000.0
var falta = 2
var subsidioAlimentacaoDiario = 400.00
var subsidioTransporteDiario = 500.00

var subsidioTransporteMensal = s.CalcularSubsidioMensal(diasUteis, subsidioTransporteDiario)
var subsidioAlimentacaoMensal = s.CalcularSubsidioMensal(diasUteis, subsidioAlimentacaoDiario)
var salarioBruto = s.CalcularSalariobruto(subsidioTransporteMensal, subsidioAlimentacaoMensal)
var totalSujeicaoIrt = s.CalcularSujeicaoIrt(subsidioTransporteMensal, subsidioAlimentacaoMensal)
var inss = s.CalcualrInssDoTrabalhador(salarioBruto) // rever se na materia coletavel o salrio é base ou bruto
var materialColetavel = s.CalcularMateriaColetavel(salarioBase, totalSujeicaoIrt, inss)
var irt = s.CalcularIrt(materialColetavel)
var totalDesconto = s.CalcularTotalDesconto(inss, irt)

func TestCalcularSubsidioMensalDeAlimentacao(t *testing.T) {
	t.Run("Deve calcular o subsidio de alimentação mensal caso for informado dias uteis de trabalho e o subsidio diario de alimentação", func(t *testing.T) {
		//act
		subsidioAlimentacaoMensal := s.CalcularSubsidioMensal(diasUteis, subsidioAlimentacaoDiario)

		//assert
		if subsidioAlimentacaoMensal <= 0 {
			t.Errorf("O subsidio de alimentaçao mensal não pode ser difrente de %.2f ", subsidioAlimentacaoMensal)
		}
	})

	t.Run("Deve calcular o subsidio de alimentação mensal caso for informado apenas o subsidio diario de alimentação", func(t *testing.T) {
		//act
		subsidioAlimentacaoMensal := s.CalcularSubsidioMensal(0, subsidioAlimentacaoDiario)

		//assert
		if subsidioAlimentacaoMensal <= 0 {
			t.Errorf("O subsidio de alimentaçao mensal não pode ser difrente de %.2f ", subsidioAlimentacaoMensal)
		}
	})

}

func TestCalcularSubsidioMensalDeTransporte(t *testing.T) {
	t.Run("Deve calcular o subsidio de transporte mensal caso for informado os dias uteis de trabalho subsidio diario de transporte", func(t *testing.T) {
		//act
		subsidioTransporteMensal := s.CalcularSubsidioMensal(diasUteis, subsidioTransporteDiario)

		//assert
		if subsidioTransporteMensal <= 0 {
			t.Errorf("O subsidio de transporte mensal não pode ser difrente de %.2f ", subsidioTransporteMensal)
		}
	})

	t.Run("Não deve calcular o subsidio de transporte mensal caso não for informado os dias uteis de trabalho subsidio diario de transporte", func(t *testing.T) {
		//act
		subsidioTransporteMensal := s.CalcularSubsidioMensal(0, 0)

		//assert
		if subsidioTransporteMensal > 0 {
			t.Errorf("O subsidio de transporte mensal não pode ser difrente de %.2f ", subsidioTransporteMensal)
		}
	})

	t.Run("Deve calcular o subsidio de transporte mensal caso for informado apenas subsidio diario de transporte", func(t *testing.T) {
		//act
		subsidioTransporteMensal := s.CalcularSubsidioMensal(0, subsidioTransporteDiario)

		//assert
		if subsidioTransporteMensal <= 0 {
			t.Errorf("O subsidio de transporte mensal não pode ser difrente de %.2f ", subsidioTransporteMensal)
		}
	})

	t.Run("Não deve calcular o subsidio de transporte mensal caso não for informado os dias uteis de trabalho e o subsidio diario de transporte", func(t *testing.T) {
		//act
		subsidioTransporteMensal := s.CalcularSubsidioMensal(0, 0)

		//assert
		if subsidioTransporteMensal > 0 {
			t.Errorf("O subsidio de transporte mensal não pode ser difrente de %.2f ", subsidioTransporteMensal)
		}
	})

}

func TestCalcularSalarioBase(t *testing.T) {
	t.Run("Deve calcular o salario bruto com dias uteis de trabalho informado", func(t *testing.T) {
		//act
		salarioBruto := s.CalcularSalariobruto(subsidioTransporteMensal, subsidioAlimentacaoMensal, float64(premios), float64(salarioBase))

		//assert

		if salarioBruto <= 0 {
			t.Errorf("O salário não pode ser %.2f ", salarioBruto)
		}
	})

	t.Run("Deve calcular o salario bruto sem dias uteis de trabalho informado", func(t *testing.T) {
		//act
		salarioBruto := s.CalcularSalariobruto(subsidioTransporteMensal, subsidioAlimentacaoMensal, float64(premios), float64(salarioBase))

		//assert
		if salarioBruto <= 0 {
			t.Errorf("O salário não pode ser %.2f ", salarioBruto)
		}
	})

	t.Run("Deve calcular o salario base apôs falta com dias uteis de trabalho informado", func(t *testing.T) {

		//act
		salarioBaseAposFalta := s.CalcualrSalarioBaseApoisFalta(diasUteis, falta, salarioBase)

		//assert

		if salarioBaseAposFalta <= 0 {
			t.Errorf("O salario base não deve ser %.2f ", salarioBase)
		}
	})

	t.Run("Deve calcular o salario base apôs falta sem dias uteis de trabalho informado", func(t *testing.T) {
		//act
		salarioBaseAposFalta := s.CalcualrSalarioBaseApoisFalta(0, falta, salarioBase)

		//assert
		if salarioBaseAposFalta <= 0 {
			t.Errorf("O salario base não deve ser %.2f ", salarioBase)
		}
	})

}

func TestCalcularInssDoTrabalhador(t *testing.T) {
	t.Run("Deve calcular o inss do trabalhador no salario bruto ou iliquido", func(t *testing.T) {
		//act
		inss := s.CalcualrInssDoTrabalhador(salarioBruto)

		//assert

		if inss <= 0 {
			t.Errorf("O inss não deve ser %.2f ", inss)
		}
	})

	t.Run("Não deve calcular o inss do trabalhador sem o salário bruto ou iliquido", func(t *testing.T) {
		//act
		inss := s.CalcualrInssDoTrabalhador(0)
		//assert
		if inss > 0 {
			t.Errorf("O inss não deve ser maior %.2f ", inss)
		}
	})
}

func TestCalcularSujeicaoDoIrt(t *testing.T) {
	t.Run("Deve calcular a sujeição do irt", func(t *testing.T) {
		//act
		totalSujeicaoIrt := s.CalcularSujeicaoIrt(subsidioTransporteMensal, subsidioAlimentacaoMensal)
		//assert

		if totalSujeicaoIrt < 0 {
			t.Errorf("A sujeição de irt não pode ser menor %.2f ", totalSujeicaoIrt)
		}
	})

	t.Run("Deve calcular a sujeição do irt sem dias uteis de trabalho informado", func(t *testing.T) {
		//act
		totalSujeicaoIrt := s.CalcularSujeicaoIrt(subsidioTransporteMensal, subsidioAlimentacaoMensal)
		//assert

		if totalSujeicaoIrt < 0 {
			t.Errorf("A sujeição de irt não pode ser menor %.2f ", totalSujeicaoIrt)
		}
	})

	t.Run("Não deve calcular a sujeição do irt sem subsidio de transporte, alimentação diario e dias uteis de trabalho informado", func(t *testing.T) {
		//act
		totalSujeicaoIrt := s.CalcularSujeicaoIrt(0, 0)
		//assert

		if totalSujeicaoIrt > 0 {
			t.Errorf("A sujeição de irt não pode ser menor %.2f ", totalSujeicaoIrt)
		}
	})
}

func TestCalcularMateriaColetave(t *testing.T) {
	t.Run("Deve calcular a materia coletavel", func(t *testing.T) {
		// arrange
		subsidioTransporteMensal := s.CalcularSubsidioMensal(diasUteis, subsidioTransporteDiario)
		subsidioAlimentacaoMensal := s.CalcularSubsidioMensal(diasUteis, subsidioAlimentacaoDiario)
		totalSujeicaoIrt := s.CalcularSujeicaoIrt(subsidioTransporteMensal, subsidioAlimentacaoMensal)
		inss := s.CalcualrInssDoTrabalhador(salarioBase)

		//act
		materialColetavel := s.CalcularMateriaColetavel(salarioBase, totalSujeicaoIrt, inss)
		//assert
		if materialColetavel < 0 {
			t.Errorf("O material coletavel não deve ser igual %.2f", materialColetavel)
		}
	})

	t.Run("Deve calcular a materia coletavel sem dias uteis de trabalho informado", func(t *testing.T) {
		// arrange
		subsidioTransporteMensal := s.CalcularSubsidioMensal(0, subsidioTransporteDiario)
		subsidioAlimentacaoMensal := s.CalcularSubsidioMensal(0, subsidioAlimentacaoDiario)
		totalSujeicaoIrt := s.CalcularSujeicaoIrt(subsidioTransporteMensal, subsidioAlimentacaoMensal)
		inss := s.CalcualrInssDoTrabalhador(salarioBase)

		//act
		materialColetavel := s.CalcularMateriaColetavel(salarioBase, totalSujeicaoIrt, inss)
		//assert
		if materialColetavel < 0 {
			t.Errorf("O material coletavel não deve ser igual %.2f", materialColetavel)
		}
	})

	t.Run("Não deve calcular a materia coletavel sem subsidio de transporte, alimentação mensal e  dias uteis de trabalho informado", func(t *testing.T) {
		// arrange
		totalSujeicaoIrt := s.CalcularSujeicaoIrt(0, 0)
		inss := s.CalcualrInssDoTrabalhador(salarioBase)

		//act
		materialColetavel := s.CalcularMateriaColetavel(salarioBase, totalSujeicaoIrt, inss)
		//assert
		if materialColetavel < 0 {
			t.Errorf("O material coletavel não deve ser igual %.2f", materialColetavel)
		}
	})
}

func TestCalcularDescontoTotal(t *testing.T) {
	t.Run("Deve calcular o total de desconto", func(t *testing.T) {
		// arrange

		//act
		totalDesconto := s.CalcularTotalDesconto(inss, irt)
		//assert

		if totalDesconto < 0 {
			t.Errorf("O total de desconto não pode ser negativo %.2f", totalDesconto)
		}
	})

	t.Run("Deve calcular o total de desconto sem dias uteis de trabalho informado", func(t *testing.T) {
		//act
		totalDesconto := s.CalcularTotalDesconto(inss, irt)
		//assert

		if totalDesconto < 0 {
			t.Errorf("O total de desconto não pode ser negativo %.2f", totalDesconto)
		}
	})

	t.Run("Deve calcular o total de desconto sem dias uteis de trabalho informado", func(t *testing.T) {
		//act
		totalDesconto := s.CalcularTotalDesconto(inss, irt)
		//assert

		if totalDesconto < 0 {
			t.Errorf("O total de desconto não pode ser negativo %.2f", totalDesconto)
		}
	})

	t.Run("Deve calcular o total de desconto sem subsidio de transporte, alimentação mensal e dias uteis de trabalho", func(t *testing.T) {
		//act
		totalDesconto := s.CalcularTotalDesconto(inss, irt)
		//assert

		if totalDesconto < 0 {
			t.Errorf("O total de desconto não pode ser calculado %.2f ", totalDesconto)
		}
	})

	t.Run("Deve calcular o total de desconto sem dias uteis de trabalho", func(t *testing.T) {
		//act
		totalDesconto := s.CalcularTotalDesconto(inss, irt)
		//assert

		if totalDesconto < 0 {
			t.Errorf("O total de desconto não pode ser negativo %.2f", totalDesconto)
		}
	})

}
func TestCalcularIrt(t *testing.T) {

	t.Run("Deve calcular o irt", func(t *testing.T) {
		//act
		irt := s.CalcularIrt(materialColetavel)
		//assert

		if irt < 0 {
			t.Errorf("O irt não pode ser negativo %.2f", irt)
		}
	})
}

func TestCalcularSalarioLiquido(t *testing.T) {
	t.Run("Deve calcular o salário líquido", func(t *testing.T) {
		//act
		salaioLiquido := s.CalcularSalarioLiquido(salarioBruto, totalDesconto)
		//assert
		if salaioLiquido < 0 {
			t.Errorf("O salário líquido não pode ser negativo %.2f", totalDesconto)
		}
	})
}
