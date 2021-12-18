package router

import (
	"app/src/router/rotas"

	"github.com/gorilla/mux"
)

//Gerar() - Retorna um routar com todas as roas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
