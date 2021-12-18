package main

import (
	"app/src/router"
	"app/src/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Rodando webapp!")
	log.Fatal(http.ListenAndServe(":3000", r))
}
