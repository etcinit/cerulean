package controllers

import (
	"strconv"

	"github.com/etcinit/cerulean/app/responses"
	"github.com/etcinit/cerulean/app/tumblr"
	"github.com/gin-gonic/gin"
)

// TumblrController handles main routes
type TumblrController struct {
	Tumblr *tumblr.Service `inject:""`
}

// Register registers the route handlers for this controller
func (control *TumblrController) Register(r *gin.RouterGroup) {
	tumblr := r.Group("/tumblr")
	{
		tumblr.GET("/", control.getIndex)
		//tumblr.GET("/:id", control.getSingle)
	}
}

// getIndex returns a list of all keywords
func (control *TumblrController) getIndex(c *gin.Context) {
	page, err := strconv.Atoi(c.Request.URL.Query().Get("page"))

	if err != nil {
		responses.SendResourceNotFound(c)
		return
	}

	posts := control.Tumblr.Paginate(page)

	responses.SendSingleResource(c, "posts", posts)
}
