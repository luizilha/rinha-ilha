package main

import (
	"luizilha/rinha-ilha/internal/people"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	peopleRouter := people.NewPeopleRouter(router)
	peopleRouter.Registry()
	router.Run(":8080")
}
