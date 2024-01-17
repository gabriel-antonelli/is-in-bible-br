package controller

import (
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/app/initial"
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/modules/search/service"
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/modules/shared"
)

type searchController struct {
	service service.SearchService
}

func NewSearchController(service service.SearchService) shared.Controller {
	return &searchController{service}
}

func (c *searchController) Handle(req shared.Request) shared.Response {
	result := c.service.Search(req[0].(string), initial.StringFile)
	if len(result) == 0 {
		return shared.Response{
			StatusCode: 404,
			Result:     "Not Found",
		}
	}
	return shared.Response{
		StatusCode: 200,
		Result:     result,
	}
}
