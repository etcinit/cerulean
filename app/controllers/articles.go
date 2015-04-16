package controllers

import (
	"strconv"

	"github.com/etcinit/cerulean/app/responses"
	"github.com/etcinit/cerulean/database/models"
	"github.com/etcinit/ohmygorm"
	"github.com/etcinit/pagination"
	"github.com/gin-gonic/gin"
)

// ArticlesController handles all routes related to articles
type ArticlesController struct {
	Connections *ohmygorm.ConnectionsService `inject:""`
	Repository  *ohmygorm.RepositoryService  `inject:""`
}

// Register registers all the routes for this controller
func (control *ArticlesController) Register(r *gin.RouterGroup) {
	articles := r.Group("/articles")
	{
		articles.GET("/", control.getIndex)
		articles.GET("/:id", control.getSingle)
	}
}

// getIndex returns a list of all keywords
func (control *ArticlesController) getIndex(c *gin.Context) {
	db, _ := control.Connections.Make()

	var totalItems int
	var articles []models.Article

	db.Model(&models.Article{}).Count(&totalItems)

	paginator := pagination.NewFromRequest(totalItems, 3, c.Request)
	pagination := paginator.ToPagination()

	db.Limit(pagination.ItemsPerPage).Offset(pagination.Offset).Find(&articles)

	responses.SendSingleResource(c, "articles", paginator.ToPaginationWithData(articles))
}

func (control *ArticlesController) getSingle(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil || !control.Repository.Exists(&models.Article{}, id) {
		responses.SendResourceNotFound(c)
		return
	}

	var article models.Article
	err = control.Repository.Find(&article, id)

	if err != nil {
		responses.SendResourceNotFound(c)
		return
	}

	responses.SendSingleResource(c, "article", article)
}
