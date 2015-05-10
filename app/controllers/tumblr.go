package controllers

import (
	"strconv"

	"github.com/etcinit/cerulean/app/cache"
	"github.com/etcinit/cerulean/app/responses"
	"github.com/etcinit/cerulean/app/tumblr"
	"github.com/etcinit/speedbump"
	"github.com/etcinit/speedbump/ginbump"
	"github.com/gin-gonic/gin"
)

// TumblrController handles main routes
type TumblrController struct {
	Tumblr *tumblr.Service     `inject:""`
	Cache  *cache.RedisService `inject:""`
}

// Register registers the route handlers for this controller
func (control *TumblrController) Register(r *gin.RouterGroup) {
	tumblr := r.Group("/tumblr")
	{
		tumblr.Use(ginbump.RateLimit(
			control.Cache.Make(),
			speedbump.PerMinuteHasher{},
			100,
		))

		tumblr.GET("/posts", control.getIndex)
		//tumblr.GET("/:id", control.getSingle)
	}
}

// getIndex returns a list of all tumblr posts
func (control *TumblrController) getIndex(c *gin.Context) {
	page := 1

	if c.Request.URL.Query().Get("page") != "" {
		var err error
		page, err = strconv.Atoi(c.Request.URL.Query().Get("page"))

		if err != nil {
			responses.SendResourceNotFound(c)
			return
		}
	}

	posts := control.Tumblr.Paginate(page)

	responses.SendSingleResource(c, "posts", posts)
}
