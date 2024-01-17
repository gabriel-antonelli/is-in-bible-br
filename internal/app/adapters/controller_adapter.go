package adapters

import (
	shared "github.com/gabriel-antonelli/is-in-the-bible-br/internal/modules/shared"
	"github.com/gin-gonic/gin"
)

func ControllerAdapter(controller shared.Controller, keys []string, c *gin.Context) {
	request := []any{}
	for _, key := range keys {
		value := c.Query(key)
		if value == "" {
			value = c.Param(key)
		}
		request = append(request, value)
	}
	request = append(request, c.Request.Body)
	response := controller.Handle(request)
	c.JSON(response.StatusCode, response.Result)
}
