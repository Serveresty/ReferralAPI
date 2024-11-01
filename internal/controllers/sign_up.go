package controllers

import (
	"ReferralAPI/models"
	cerr "ReferralAPI/pkg/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rc *RefController) SignUp(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if _, err := rc.refServ.TokenValidation(token); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": cerr.ErrAlreadyAuthorized.Error()})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := rc.refServ.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "account created"})
}
