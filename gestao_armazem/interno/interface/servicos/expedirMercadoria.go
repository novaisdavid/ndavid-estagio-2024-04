package servicos

import (
	"github.com/acmllda/interno/casosDeUso/expedirMercadoria"
	 guiaRemessa "github.com/acmllda/interno/dominio/agregados"
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
	expediu := expedir.Executa(guia)

	if expediu {
		fmt.Println("MERCADORIA EXPEDIDADA")
		return
	}

	fmt.Println("MERCADORIA NÂO EXPEDIDADA")

}
