package encomenda

import (
	lote   "Stock_Acme/lote"
	"fmt"
	"time"
)

type Encomenda struct {

	Cliente              string
	IdentificadorProduto string
	Quantidade           int
}


func (e Encomenda) RetiraEncomenda(en Encomenda, lotes []lote.Lote) {
	 
	r :=en.produtoComDataValidadeProxima(lotes, en.IdentificadorProduto)
	fmt.Println("DATAS: ", r)
}


// ao escolher o produto traz com a data de expiração mais proxima

func (e Encomenda) produtoComDataValidadeProxima(lotes []lote.Lote, idProdduto string) []lote.Lote{
	// Obtendo a data atual
	dataAtual := time.Now()
	// Inicializando com a primeira data de validade e índices
	var indicesMaisProximos []int
	dataMaisProxima := time.Time{}
	// Iterando sobre os lotes para encontrar a data de validade mais próxima
	
	for i, lote := range lotes {
		if lote.IdProduto == idProdduto {
			dataValidade, err := time.Parse("2006-01-02", lote.DataDeValidade)
			if err != nil {
				fmt.Println("Erro ao analisar a data de validade:", err)
				continue
			}

			if dataValidade.After(dataAtual) && (dataMaisProxima.IsZero() || dataValidade.Before(dataMaisProxima)) {
				dataMaisProxima = dataValidade
				indicesMaisProximos = []int{i}
			} else if dataValidade.Equal(dataMaisProxima) {
				indicesMaisProximos = append(indicesMaisProximos, i)
			}
		}
	}

	// Verificando se foram encontrados lotes
	if len(indicesMaisProximos) > 0 {
		fmt.Println("Lotes com data de validade mais próxima:")
		for _, indice := range indicesMaisProximos {
			fmt.Println(lotes[indice])
		}
	} else {
		fmt.Println("Não há lotes com data de validade futura.")
	}

	return lotes
}

