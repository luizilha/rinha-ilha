package people

import (
	"luizilha/rinha-ilha/internal/storage"
	"luizilha/rinha-ilha/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct { }

/*
func (con *controller) AllPeople(c *gin.Context) {
  db := storage.PostgresConnection()
  repo := NewRepository(db)
	all, err := repo.AllPeople()
	if err != nil {
    c.IndentedJSON(http.StatusBadRequest, err.Error())
	}
	c.IndentedJSON(http.StatusOK, all)
}
*/

func (con *controller) Count(c *gin.Context) {
  db := storage.PostgresConnection()
  repo := NewRepository(db)
  value, err := repo.Count()
  if err != nil {
    c.IndentedJSON(http.StatusBadRequest, err.Error())
  }
  c.IndentedJSON(http.StatusOK, value)

}

func (con *controller) Insert(c *gin.Context) {
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
