package tumblr

import (
	"github.com/etcinit/pagination"
	"github.com/etcinit/tumble/client"
	"github.com/etcinit/tumble/client/blog"
	"github.com/jacobstr/confer"
)

// Service provides access to the Tumblr API
type Service struct {
	Tumblr *client.TumblrClient `inject:""`
	Config *confer.Config       `inject:""`
}

// Paginate get blog posts by page
func (s *Service) Paginate(page int) pagination.Pagination {
	limit := s.Config.GetInt("tumblr.postsPerPage")

	if page < 1 {
		page = 1
	}

	response, _ := s.Tumblr.Blog.GetPosts(
		s.Config.GetString("tumblr.hostname"),
		blog.GetPostsParameters{
			Limit:  limit,
			Offset: page * limit,
		},
	)

	paginator := pagination.New(
		response.Response.TotalPosts,
		limit,
		page,
	)

	return paginator.ToPaginationWithData(response.Response.Posts)
}
