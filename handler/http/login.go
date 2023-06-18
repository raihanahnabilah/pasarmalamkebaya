package http

import (
	"net/http"
	"pasarmalamkebaya/dto"
	"pasarmalamkebaya/handler/response"
	"pasarmalamkebaya/usecase"

	"github.com/gin-gonic/gin"
)

type loginHandler struct {
	loginUsecase usecase.LoginUsecase
}

func NewLoginHandler(loginUsecase usecase.LoginUsecase) *loginHandler {
	return &loginHandler{loginUsecase}
}

func (h *loginHandler) Route(r *gin.RouterGroup) {
	r.POST("/api/v1/login", h.Login)
}

func (h *loginHandler) Login(c *gin.Context) {
	var input dto.LoginRequestBody

	// Validate the inputs!
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := response.Response("Invalid input for Login", 400, http.StatusText(http.StatusBadRequest), errorMessage)
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	// Plug in the data!
	user, err := h.loginUsecase.Login(input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := response.Response("Failed to Login", 400, http.StatusText(http.StatusBadRequest), errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.Response("Logged In Succesfully!", http.StatusOK, http.StatusText(http.StatusOK), user)
	c.JSON(http.StatusOK, response)

}
