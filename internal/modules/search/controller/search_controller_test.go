package controller

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockSearchService struct {
	mock.Mock
}

func (m *MockSearchService) Search(words string) map[string]int {
	args := m.Called(words)
	return args.Get(0).(map[string]int)
}

func TestSearchControllerHandle200AndResult(t *testing.T) {
	searchService := new(MockSearchService)

	request := "amor+Jesus+a+amOR"
	response := map[string]int{
		"jesus": 42,
		"amor":  23,
	}

	searchService.On("Search", request).Return(response)

	searchController := NewSearchController(searchService)

	result := searchController.Handle([]any{request})

	assert.Equal(t, 200, result.StatusCode, "Expected 200 status code")
	assert.Equal(t, response, result.Result, "Expected response")
}

func TestSearchControllerHandle404(t *testing.T) {
	searchService := new(MockSearchService)

	searchService.On("Search", "test").Return(make(map[string]int))

	searchController := NewSearchController(searchService)

	result := searchController.Handle([]any{"test"})

	assert.Equal(t, 404, result.StatusCode, "Expected 404 status code")
	assert.Equal(t, "Not Found", result.Result, "Expected not found response")
}

func TestSearchControllerHandle400(t *testing.T) {
	searchService := new(MockSearchService)

	searchController := NewSearchController(searchService)

	result := searchController.Handle([]any{})

	assert.Equal(t, 400, result.StatusCode, "Expected 400 status code")
	assert.Equal(t, "No words to search for", result.Result, "Expected no words to search for response")
}
