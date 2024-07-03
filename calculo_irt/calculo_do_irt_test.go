package calculodoirt_test

import (
	"testing"

	calculo "github.com/calculo_irt"
)

func TestCalcularIrt(t *testing.T) {
	t.Run("Deve calcular o subsidio de alimentação mensal caso for informado dias uteis de trabalho e o subsidio diario de alimentação", func(t *testing.T) {
		// arrange
		s := calculo.Irt{}.NewIrt()
		diasUteis := 30
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
		s := calculo.Irt{}.NewIrt()
		subsidioAlimentacaoDiario := 500.00
		//act
		subsidioAlimentacaoMensal := s.CalcularSubsidioMensal(0, subsidioAlimentacaoDiario)

		//assert
		if subsidioAlimentacaoMensal <= 0 {
			t.Errorf("O subsidio de alimentaçao mensal não pode ser difrente de %.2f ", subsidioAlimentacaoMensal)
		}
	})

	t.Run("Deve calcular o subsidio de transporte mensal caso for informado os dias uteis de trabalho subsidio diario de transporte", func(t *testing.T) {
		// arrange
		s := calculo.Irt{}.NewIrt()
		subsidioTransporteDiario := 300.00
		diasUteis := 30
		//act
		subsidioTransporteMensal := s.CalcularSubsidioMensal(diasUteis, subsidioTransporteDiario)

		//assert
		if subsidioTransporteMensal <= 0 {
			t.Errorf("O subsidio de transporte mensal não pode ser difrente de %.2f ", subsidioTransporteMensal)
		}
	})

	t.Run("Não deve calcular o subsidio de transporte mensal caso não for informado os dias uteis de trabalho subsidio diario de transporte", func(t *testing.T) {
		// arrange
		s := calculo.Irt{}.NewIrt()
		//act
		subsidioTransporteMensal := s.CalcularSubsidioMensal(0, 0)

		//assert
		if subsidioTransporteMensal > 0 {
			t.Errorf("O subsidio de transporte mensal não pode ser difrente de %.2f ", subsidioTransporteMensal)
		}
	})

	t.Run("Deve calcular o subsidio de transporte mensal caso for informado apenas subsidio diario de transporte", func(t *testing.T) {
		// arrange
		s := calculo.Irt{}.NewIrt()
		subsidioTransporteDiario := 300.00
		//act
		subsidioTransporteMensal := s.CalcularSubsidioMensal(0, subsidioTransporteDiario)

		//assert
		if subsidioTransporteMensal <= 0 {
			t.Errorf("O subsidio de transporte mensal não pode ser difrente de %.2f ", subsidioTransporteMensal)
		}
	})

	t.Run("Não deve calcular o subsidio de transporte mensal caso não for informado os dias uteis de trabalho e o subsidio diario de transporte", func(t *testing.T) {
		// arrange
		s := calculo.Irt{}.NewIrt()
		//act
		subsidioTransporteMensal := s.CalcularSubsidioMensal(0, 0)

		//assert
		if subsidioTransporteMensal > 0 {
			t.Errorf("O subsidio de transporte mensal não pode ser difrente de %.2f ", subsidioTransporteMensal)
		}
	})
	t.Run("Deve calcular o salario bruto", func(t *testing.T) {
		// arrange
		s := calculo.Irt{}.NewIrt()
		subTransport := 120.23
		subAliment := 23.23
		premio := 1000.0
		salarioBase := 30000

		//act
		sl := s.CalcularSalariobruto(subTransport, subAliment, float64(premio), float64(salarioBase))

		//assert

		if sl <= 0 {
			t.Errorf("O salário não pode ser %.2f ", sl)
		}
	})

}
