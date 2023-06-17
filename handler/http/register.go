package handler

import (
	"net/http"
	"pasarmalamkebaya/dto"
	"pasarmalamkebaya/handler/response"
	"pasarmalamkebaya/usecase"

	"github.com/gin-gonic/gin"
)

type registerHandler struct {
	registerUsecase usecase.RegisterUsecase
}

func NewRegisterHandler(registerUsecase usecase.RegisterUsecase) *registerHandler {
	return &registerHandler{registerUsecase}
}

func (h *registerHandler) Route(r *gin.RouterGroup) {
	r.GET("/api/v1/register", h.Register)
}

func (h *registerHandler) Register(c *gin.Context) {
	var dto dto.RegisterRequestBody

	// Validate the inputs!
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := response.Response("Failed to add user", 400, "Bad request", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	// Plug in the data!
	user, err := h.registerUsecase.Register(dto)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := response.Response("Failed to add user", 400, "Bad request", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.Response("Added user", 200, "Success", user)
	c.JSON(http.StatusOK, response)
}
