package modelos

import (
	"errors"
	"strings"
	"time"
)

type Cliente struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Cpf      string    `json:"cpf,omitempty"`
	Fone     int       `json:"fone,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (cliente *Cliente) Preparar() error {
	if erro := cliente.validar(); erro != nil {
		return erro
	}

	cliente.formatar()
	return nil

}

func (cliente *Cliente) validar() error {
	if cliente.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco.")
	}

	if cliente.Cpf == "" {
		return errors.New("O CPF é obrigatório e não pode estar em branco.")
	}

	return nil
}

func (cliente *Cliente) formatar() {
	cliente.Nome = strings.TrimSpace(cliente.Nome)
}
