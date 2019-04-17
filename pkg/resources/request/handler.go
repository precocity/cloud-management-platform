package request

import (
	"net/http"

	"github.com/labstack/echo"
)

//Handler will handle all api requests for resource requests
type Handler struct {
	service *Service
}

//CreateHandler creates a new instance of the resource request handler
func CreateHandler(service *Service) *Handler {

	h := &Handler{
		service: service,
	}

	return h
}

//HandleCreateRequest will handle api requests to create resource requests
func (h Handler) HandleCreateRequest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
