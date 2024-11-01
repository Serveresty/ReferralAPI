package controllers

import (
	"ReferralAPI/models"
	cerr "ReferralAPI/pkg/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User sign-in
// @Description Authenticates a user and returns a JWT token
// @Tags authentication
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer token"
// @Param user body models.User true "User credentials"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/sign-in [post]
func (rc *RefController) SignIn(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if _, err := rc.refServ.TokenValidation(token); err == nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: cerr.ErrAlreadyAuthorized.Error()})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	token, err := rc.refServ.AuthorizeUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "authorized", Data: token})
}
