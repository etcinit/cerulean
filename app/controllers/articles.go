package controllers

import (
	"strconv"

	"github.com/etcinit/cerulean/app/cache"
	"github.com/etcinit/cerulean/app/responses"
	"github.com/etcinit/cerulean/database/models"
	"github.com/etcinit/ohmygorm"
	"github.com/etcinit/pagination"
	"github.com/etcinit/speedbump"
	"github.com/etcinit/speedbump/ginbump"
	"github.com/gin-gonic/gin"
)

// ArticlesController handles all routes related to articles
type ArticlesController struct {
	Connections *ohmygorm.ConnectionsService `inject:""`
	Repository  *ohmygorm.RepositoryService  `inject:""`
	Cache       *cache.RedisService          `inject:""`
}

// Register registers all the routes for this controller
func (control *ArticlesController) Register(r *gin.RouterGroup) {
	articles := r.Group("/articles")
	{
		articles.Use(ginbump.RateLimit(
			control.Cache.Make(),
			speedbump.PerMinuteHasher{},
			100,
		))

		articles.GET("/", control.getIndex)
		articles.GET("/title/:encoded", control.getSingleByTitle)
		articles.GET("/id/:id", control.getSingle)
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

	db.Preload("Author").Limit(pagination.ItemsPerPage).Offset(pagination.Offset).Order("created_at desc").Find(&articles)

	responses.SendSingleResource(c, "articles", paginator.ToPaginationWithData(articles))
}

func (control *ArticlesController) getSingle(c *gin.Context) {
	db, _ := control.Connections.Make()

	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil || !control.Repository.Exists(&models.Article{}, id) {
		responses.SendResourceNotFound(c)
		return
	}

	var article models.Article
	err = control.Repository.FirstOrFail(&article, db.Preload("Author").Where(&models.Article{ID: id}))

	if err != nil {
		responses.SendResourceNotFound(c)
		return
	}

	responses.SendSingleResource(c, "article", article)
}

func (control *ArticlesController) getSingleByTitle(c *gin.Context) {
	db, _ := control.Connections.Make()

	titleEncoded := c.Params.ByName("encoded")

	var article models.Article
	err := control.Repository.FirstOrFail(&article, db.Preload("Author").Where(&models.Article{TitleEncoded: titleEncoded}))

	if err != nil || titleEncoded == "" {
		responses.SendResourceNotFound(c)
		return
	}

	responses.SendSingleResource(c, "article", article)
}
