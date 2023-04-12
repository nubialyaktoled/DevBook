package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas rotas da api
type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {

	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
	}

	return r

}
