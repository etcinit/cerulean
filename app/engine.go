package app

import (
	"github.com/etcinit/cerulean/app/controllers"
	"github.com/gin-gonic/gin"
)

// EngineService provides the API engine
type EngineService struct {
	Front    controllers.FrontController    `inject:""`
	Articles controllers.ArticlesController `inject:""`
}

// New creates a new instance of an API engine
func (e *EngineService) New() *gin.Engine {
	router := gin.Default()

	e.Front.Register(router)

	v1 := router.Group("/v1")
	{
		e.Articles.Register(v1)
	}

	return router
}
