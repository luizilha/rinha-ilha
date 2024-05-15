package main

import (
//	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type pessoa struct {
	ID         string `json:"id"`
	apelido    string
	nome       string
	nascimento string
	stack      []string
}

var pessoas = []pessoa{
  { ID: "nsestrtsne", apelido: "ilha", nome: "luiz", nascimento: "07/03/1993", stack: []string{} },
}

func contagemPessoas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, len(pessoas))
}

func main() {
	router := gin.Default()
	router.GET("/contagem-pessoas", contagemPessoas)
	router.Run(":8080")
}
