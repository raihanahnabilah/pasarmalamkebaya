package http

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
	r.POST("/api/v1/register", h.Register)
}

func (h *registerHandler) Register(c *gin.Context) {
	var input dto.RegisterRequestBody

	// Validate the inputs!
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := response.Response("Failed to add user", 400, http.StatusText(http.StatusBadRequest), errorMessage)
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	// Plug in the data!
	user, err := h.registerUsecase.Register(input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := response.Response("Failed to add user", 400, http.StatusText(http.StatusBadRequest), errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.Response("Added user", http.StatusOK, http.StatusText(http.StatusOK), user)
	c.JSON(http.StatusOK, response)
}
