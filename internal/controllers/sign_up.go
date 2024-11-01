package controllers

import (
	"ReferralAPI/models"
	cerr "ReferralAPI/pkg/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User sign-up
// @Description Registers a new user and returns a success message
// @Tags authentication
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer token"
// @Param user body models.User true "User details"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/sign-up [post]
func (rc *RefController) SignUp(c *gin.Context) {
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

	err := rc.refServ.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, models.Response{Message: "account created"})
}
