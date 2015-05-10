package app

import (
	"net/http"

	"github.com/etcinit/cerulean/app/controllers"
	"github.com/gin-gonic/gin"
)

// EngineService provides the API engine
type EngineService struct {
	Front    controllers.FrontController    `inject:""`
	Articles controllers.ArticlesController `inject:""`
	Tumblr   controllers.TumblrController   `inject:""`
}

// New creates a new instance of an API engine
func (e *EngineService) New() *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	})

	e.Front.Register(router)

	v1 := router.Group("/v1")
	{
		e.Articles.Register(v1)
		e.Tumblr.Register(v1)
	}

	return router
}
