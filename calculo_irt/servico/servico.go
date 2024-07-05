package servico

import (
	"fmt"
	"strings"

	dados "github.com/calculo_irt/entidade"
)

func CalcularIrt(irt dados.CalculoIrt, dadosIrt dados.DadosIrt) {

	subsidioAlimentacaoMensal := irt.CalcularSubsidioMensal(dadosIrt.DiasUteis(), dadosIrt.SubsidioAlimentacao())
	subsidioTransporteMensal := irt.CalcularSubsidioMensal(dadosIrt.DiasUteis(), dadosIrt.SubsidioTransporte())
	salarioBaseAposFalta := irt.CalcualrSalarioBaseApoisFalta(dadosIrt.DiasUteis(), dadosIrt.Falta(), dadosIrt.SalarioBase())
	salarioBruto := irt.CalcularSalariobruto(float64(dadosIrt.DiasUteis()),
		dadosIrt.SubsidioAlimentacao(),
		dadosIrt.SubsidioTransporte(),
		dadosIrt.Premios(),
		dadosIrt.SalarioBase(),
	)
	sujeicaoIrt := irt.CalcularSujeicaoIrt(subsidioAlimentacaoMensal, subsidioTransporteMensal)
	inss := irt.CalcualrInssDoTrabalhador(salarioBruto)
	materiaColetavel := irt.CalcularMateriaColetavel(salarioBruto, sujeicaoIrt, inss)
	resultadoIrt := irt.CalcularIrt(materiaColetavel)
	totalDesconto := irt.CalcularTotalDesconto(inss, resultadoIrt)
	salarioLíquido := irt.CalcularSalarioLiquido(salarioBruto, totalDesconto)

	fmt.Println()
	fmt.Println("Salário Base: ", formatarNumero(dadosIrt.SalarioBase()), "kz")
	fmt.Println("========= Incrementos ========")
	fmt.Println("Subsídio de Alimentação: ", formatarNumero(subsidioAlimentacaoMensal), "kz")
	fmt.Println("Subsídio de Transporte: ", formatarNumero(subsidioTransporteMensal), "kz")
	fmt.Println("Prémios:  ", formatarNumero(dadosIrt.Premios()), "kz")
	fmt.Println()
	fmt.Println("========= Descontos ========")
	fmt.Println("Salário Base Após Falta: ", formatarNumero(salarioBaseAposFalta), "kz")
	fmt.Println("Segurança Social: ", formatarNumero(inss), "kz")
	fmt.Println("IRT: ", formatarNumero(resultadoIrt), "kz")
	fmt.Println("Salário Líquido: ", formatarNumero(salarioLíquido), "kz")
	fmt.Println()

}

func formatarNumero(valor float64) string {
	numero := fmt.Sprintf("%.2f", valor)
	partes := strings.Split(numero, ".")
	parteintegral := partes[0]
	partedecimal := partes[1]

	var builder strings.Builder
	tamanho := len(parteintegral)
	for i, c := range parteintegral {
		if i > 0 && (tamanho-i)%3 == 0 {
			builder.WriteRune('.')
		}
		builder.WriteRune(c)
	}
	return builder.String() + "," + partedecimal
}
