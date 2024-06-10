package servicos

import (
	"ACMELDA/interno/casosDeUso/expedirMercadoria"
	"ACMELDA/interno/dominio/agregados/guiaRemessa"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func ExpedirMercadoria(idProduto string, quantidade int) {
	rand.Seed(time.Now().UnixNano())
	idGuia := rand.Intn(101)

	guia := guiaRemessa.NewGuiaRemessa(strconv.Itoa(idGuia), idProduto, quantidade)

	expedir := expedirMercadoria.ExpedirMercadoria{}
	expediu := expedir.Executa(&guia)

	if expediu {
		fmt.Println("MERCADORIA EXPEDIDADA")
		return
	}

	fmt.Println("MERCADORIA NÂO EXPEDIDADA")

}
