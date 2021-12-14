package controllers

import (
	"api/repositorios"
	"api/src/banco"
	"api/src/modelos"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CriarCliente(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var cliente modelos.Cliente

	if erro = json.Unmarshal(corpoRequest, &cliente); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = cliente.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeClientes(db)
	cliente.ID, erro = repositorio.Criar(cliente)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, cliente)
}

func BuscarCliente(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	clienteID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeClientes(db)
	cliente, erro := repositorio.BuscarPorID(clienteID)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	respostas.JSON(w, http.StatusOK, cliente)
}

func AtualizarCliente(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	clienteID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var cliente modelos.Cliente
	if erro = json.Unmarshal(corpoRequest, &cliente); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = cliente.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeClientes(db)
	erro = repositorio.Atualizar(clienteID, cliente)

	respostas.JSON(w, http.StatusNoContent, nil)

}

func DeletarCliente(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	clienteID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeClientes(db)
	if erro = repositorio.Deletar(clienteID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

func BuscarClientes(w http.ResponseWriter, r *http.Request) {
	nomeOuCpf := strings.ToLower(r.URL.Query().Get("cliente"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeClientes(db)
	clientes, erro := repositorio.Buscar(nomeOuCpf)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, clientes)
}
