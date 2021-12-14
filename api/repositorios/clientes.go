package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type clientes struct {
	db *sql.DB
}

func NovoRepositorioDeClientes(db *sql.DB) *clientes {
	return &clientes{db}
}

func (repositorio clientes) Criar(cliente modelos.Cliente) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into clientes(nome, cpf, fone) values (?,?,?) ")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(cliente.Nome, cliente.Cpf, cliente.Fone)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

func (repositorio clientes) Buscar(nomeOuCpf string) ([]modelos.Cliente, error) {
	nomeOuCpf = fmt.Sprintf("%%%s%%", nomeOuCpf)

	linhas, erro := repositorio.db.Query(
		"select id, nome, cpf, fone, criadoEm from clientes where nome like ? or cpf like ?", nomeOuCpf, nomeOuCpf)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var clientes []modelos.Cliente

	for linhas.Next() {
		var cliente modelos.Cliente

		if erro = linhas.Scan(
			&cliente.ID,
			&cliente.Nome,
			&cliente.Cpf,
			&cliente.Fone,
			&cliente.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		clientes = append(clientes, cliente)
	}

	return clientes, nil

}

func (repositorio clientes) BuscarPorID(ID uint64) (modelos.Cliente, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, cpf, fone, criadoEm from clientes where id = ? ", ID)

	if erro != nil {
		fmt.Println("Não localizei o cliente")
		return modelos.Cliente{}, erro
	}

	defer linhas.Close()

	var cliente modelos.Cliente

	if linhas.Next() {
		if erro = linhas.Scan(
			&cliente.ID,
			&cliente.Nome,
			&cliente.Cpf,
			&cliente.Fone,
			&cliente.CriadoEm,
		); erro != nil {
			return modelos.Cliente{}, erro
		}

	}

	return cliente, nil

}

func (repositorio clientes) Atualizar(ID uint64, cliente modelos.Cliente) error {
	statement, erro := repositorio.db.Prepare(
		"update clientes set nome = ?, cpf = ?, fone = ? where id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(cliente.Nome, cliente.Cpf, cliente.Fone, ID); erro != nil {
		return erro
	}

	return nil

}

func (repositorio clientes) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from clientes where id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}

	return nil

}
