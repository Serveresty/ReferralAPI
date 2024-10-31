package controllers

import (
	"ReferralAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rc *RefController) SignUp(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if _, err := rc.refServ.TokenValidation(token); err == nil {
		c.JSON(http.StatusBadRequest, models.Response{Status: "error", Message: "already authorized"})
		return
	}
}
