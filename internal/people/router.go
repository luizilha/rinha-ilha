package people

import (
	"github.com/gin-gonic/gin"
)

type router struct {
  engine *gin.Engine
}

func NewPeopleRouter(e *gin.Engine) (*router){
  return &router{
    engine: e,
  }
}

func (p *router) Registry() {
  controller := &controller{}
  p.engine.GET("contagem-pessoas", controller.Count)
}
