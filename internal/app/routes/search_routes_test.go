package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/app/middlewares"
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTest() {
	config.GetDB("../../../words-in-the-bible-db/")
}

func TestSearchWordRouteSuccess(t *testing.T) {
	setupTest()

	router := gin.Default()
	router = SetupRoutes(middlewares.AddCorsMiddleWare(router))

	request := httptest.NewRequest("GET", "/search/jesus+Amor+jeSUs+a+", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, `{"amor":266,"jesus":1075}`, response.Body.String(), "Response body is expected")
}

func TestSearchWordRouteNotFound(t *testing.T) {
	router := gin.Default()
	router = SetupRoutes(middlewares.AddCorsMiddleWare(router))

	request := httptest.NewRequest("GET", "/search/a+", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code, "OK response is expected")
	assert.Equal(t, "\"Not Found\"", response.Body.String(), "Response body is expected")
}

func TestSearchWordRouteKeyNotFoundReturn0(t *testing.T) {
	defer config.CloseDB()
	router := gin.Default()
	router = SetupRoutes(middlewares.AddCorsMiddleWare(router))

	request := httptest.NewRequest("GET", "/search/aaaa", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, `{"aaaa":0}`, response.Body.String(), "Response body is expected")
}
