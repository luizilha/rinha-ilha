package people

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func NewPeopleRouter(e *gin.Engine) *Router {
	return &Router{
		engine: e,
	}
}

func (p *Router) Registry() {
	controller := &Controller{}
	p.engine.GET("contagem-pessoas", controller.Count)
	p.engine.GET("pessoas/:id", controller.FindById)
	p.engine.POST("pessoa", controller.Insert)

}
