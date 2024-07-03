package calculodoirt_test

import (
	"testing"

	calculo "github.com/calculo_irt"
)

func TestCalcularIrt(t *testing.T) {
	s := calculo.Irt{}.NewIrt()
	diasUteis := 30
	salarioBase := 30000.0
	subsidioAlimentacaoDiario := 400.00
	subsidioTransporteDiario := 300.00

	t.Run("Deve calcular o subsidio de alimentação mensal caso for informado dias uteis de trabalho e o subsidio diario de alimentação", func(t *testing.T) {
		// arrange
		subsidioAlimentacaoDiario := 500.00
		//act
		subsidioAlimentacaoMensal := s.CalcularSubsidioMensal(diasUteis, subsidioAlimentacaoDiario)

		//assert
		if subsidioAlimentacaoMensal <= 0 {
			t.Errorf("O subsidio de alimentaçao mensal não pode ser difrente de %.2f ", subsidioAlimentacaoMensal)
		}
	})

	t.Run("Deve calcular o subsidio de alimentação mensal caso for informado apenas o subsidio diario de alimentação", func(t *testing.T) {
		// arrange
		//act
		subsidioAlimentacaoMensal := s.CalcularSubsidioMensal(0, subsidioAlimentacaoDiario)

		//assert
		if subsidioAlimentacaoMensal <= 0 {
			t.Errorf("O subsidio de alimentaçao mensal não pode ser difrente de %.2f ", subsidioAlimentacaoMensal)
		}
	})

	t.Run("Deve calcular o subsidio de transporte mensal caso for informado os dias uteis de trabalho subsidio diario de transporte", func(t *testing.T) {
		// arrange
		//act
		subsidioTransporteMensal := s.CalcularSubsidioMensal(diasUteis, subsidioTransporteDiario)

		//assert
		if subsidioTransporteMensal <= 0 {
			t.Errorf("O subsidio de transporte mensal não pode ser difrente de %.2f ", subsidioTransporteMensal)
		}
	})

	t.Run("Não deve calcular o subsidio de transporte mensal caso não for informado os dias uteis de trabalho subsidio diario de transporte", func(t *testing.T) {
		// arrange
		//act
		subsidioTransporteMensal := s.CalcularSubsidioMensal(0, 0)

		//assert
		if subsidioTransporteMensal > 0 {
			t.Errorf("O subsidio de transporte mensal não pode ser difrente de %.2f ", subsidioTransporteMensal)
		}
	})

	t.Run("Deve calcular o subsidio de transporte mensal caso for informado apenas subsidio diario de transporte", func(t *testing.T) {
		// arrange
		//act
		subsidioTransporteMensal := s.CalcularSubsidioMensal(0, subsidioTransporteDiario)

		//assert
		if subsidioTransporteMensal <= 0 {
			t.Errorf("O subsidio de transporte mensal não pode ser difrente de %.2f ", subsidioTransporteMensal)
		}
	})

	t.Run("Não deve calcular o subsidio de transporte mensal caso não for informado os dias uteis de trabalho e o subsidio diario de transporte", func(t *testing.T) {
		// arrange
		//act
		subsidioTransporteMensal := s.CalcularSubsidioMensal(0, 0)

		//assert
		if subsidioTransporteMensal > 0 {
			t.Errorf("O subsidio de transporte mensal não pode ser difrente de %.2f ", subsidioTransporteMensal)
		}
	})
	t.Run("Deve calcular o salario bruto com dias uteis de trabalho informado", func(t *testing.T) {
		// arrange
		subTransporteMensal := s.CalcularSubsidioMensal(diasUteis, 200.23)
		subAlimentacaoMensal := s.CalcularSubsidioMensal(diasUteis, 300.23)
		premio := 1000.0

		//act
		salarioBruto := s.CalcularSalariobruto(subTransporteMensal, subAlimentacaoMensal, float64(premio), float64(salarioBase))

		//assert

		if salarioBruto <= 0 {
			t.Errorf("O salário não pode ser %.2f ", salarioBruto)
		}
	})

	t.Run("Deve calcular o salario bruto sem dias uteis de trabalho informado", func(t *testing.T) {
		// arrange
		subTransporteMensal := s.CalcularSubsidioMensal(0, 200.23)
		subAlimentacaoMensal := s.CalcularSubsidioMensal(0, 300.23)
		premio := 1000.0

		//act
		salarioBruto := s.CalcularSalariobruto(subTransporteMensal, subAlimentacaoMensal, float64(premio), float64(salarioBase))

		//assert

		if salarioBruto <= 0 {
			t.Errorf("O salário não pode ser %.2f ", salarioBruto)
		}
	})

	t.Run("Deve calcular o salario base apôs falta com dias uteis de trabalho informado", func(t *testing.T) {
		// arrange
		falta := 2
		//act
		salarioBaseAposFalta := s.CalcualrSalarioBaseApoisFalta(diasUteis, falta, salarioBase)

		//assert

		if salarioBaseAposFalta <= 0 {
			t.Errorf("O salario base não deve ser %.2f ", salarioBase)
		}
	})

	t.Run("Deve calcular o salario base apôs falta sem dias uteis de trabalho informado", func(t *testing.T) {
		// arrange
		falta := 2
		//act
		salarioBaseAposFalta := s.CalcualrSalarioBaseApoisFalta(0, falta, salarioBase)

		//assert

		if salarioBaseAposFalta <= 0 {
			t.Errorf("O salario base não deve ser %.2f ", salarioBase)
		}
	})

	t.Run("Deve calcular o inss do trabalhador", func(t *testing.T) {
		// arrange
		//act
		inss := s.CalcualrInssDoTrabalhador(salarioBase)

		//assert

		if inss <= 0 {
			t.Errorf("O inss não deve ser %.2f ", inss)
		}
	})

	t.Run("Deve calcular a sujeição do irt", func(t *testing.T) {
		// arrange
		subsidioTransporteMensal := s.CalcularSubsidioMensal(diasUteis, subsidioTransporteDiario)
		subsidioAlimentacaoMensal := s.CalcularSubsidioMensal(diasUteis, subsidioAlimentacaoDiario)
		//act
		totalSujeicaoIrt := s.CalcularSujeicaoIrt(subsidioTransporteMensal, subsidioAlimentacaoMensal)
		//assert

		if totalSujeicaoIrt < 0 {
			t.Errorf("A sujeição de irt não pode ser menor %.2f ", totalSujeicaoIrt)
		}
	})

	t.Run("Deve calcular o material coletavel", func(t *testing.T) {
		// arrange
		subsidioTransporteMensal := s.CalcularSubsidioMensal(diasUteis, subsidioTransporteDiario)
		subsidioAlimentacaoMensal := s.CalcularSubsidioMensal(diasUteis, subsidioAlimentacaoDiario)
		totalSujeicaoIrt := s.CalcularSujeicaoIrt(subsidioTransporteMensal, subsidioAlimentacaoMensal)
		inss := s.CalcualrInssDoTrabalhador(salarioBase)
		//act
		materialColetavel := s.CalcularMaterialColetavel(salarioBase, totalSujeicaoIrt, inss)
		//assert
		//saber se material coletavel pode ser zero
		if materialColetavel < 0 {
			t.Errorf("O material coletavel não deve ser igual %.2f", materialColetavel)
		}
	})
}
