package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// Gerar string base 64 para usar como key do env
/*func init() {
	chave := make([]byte, 64)
	fmt.Println(chave)

	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}*/

func main() {

	config.Carregar()

	r := router.Gerar()

	fmt.Printf("Escutando na porta %d", config.Porta)
	//fmt.Println("URL", config.StringConexaoBanco)
	//	fmt.Println("Rodando a API!")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
