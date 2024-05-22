package people

import (
	"luizilha/rinha-ilha/internal/storage"
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

