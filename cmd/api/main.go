package main

import (
	"github.com/gin-gonic/gin"
	"luizilha/rinha-ilha/internal/people"
)

func main() {
	router := gin.Default()
	peopleRouter := people.NewPeopleRouter(router)
	peopleRouter.Registry()
	router.Run(":8080")
}
