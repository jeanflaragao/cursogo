package controllers

import (
	"app/src/utils"
	"net/http"
)

//CarregarTelaDeLogin() - Carrega a tela de login da aplicação
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

//CarregarCadastroDeUsuario() - Carrega a tela de cadastro de usuário
func CarregarCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}
