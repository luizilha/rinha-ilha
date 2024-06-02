package people

import (
	"luizilha/rinha-ilha/internal/storage"
	"luizilha/rinha-ilha/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (con *Controller) Count(c *gin.Context) {
	db := storage.PostgresConnection()
	repo := NewRepository(db)
	value, err := repo.Count()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}
	c.IndentedJSON(http.StatusOK, value)

}

func (con *Controller) Insert(c *gin.Context) {
	db := storage.PostgresConnection()
	repo := NewRepository(db)

	var newPeople types.People
	if err0 := c.BindJSON(&newPeople); err0 != nil {
		c.IndentedJSON(http.StatusBadRequest, err0)
	}

	row, err := repo.Insert(&newPeople)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	} else {
		c.IndentedJSON(http.StatusCreated, row)
	}
}

func (con *Controller) FindById(c *gin.Context) {
	db := storage.PostgresConnection()
	repo := NewRepository(db)

	var uri types.UriID
	if err := c.ShouldBindUri(&uri); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	p, err0 := repo.FindById(uri.ID)
	if err0 != nil {
		c.IndentedJSON(http.StatusBadRequest, err0.Error())
	} else {
		c.IndentedJSON(http.StatusOK, p)
	}
}
