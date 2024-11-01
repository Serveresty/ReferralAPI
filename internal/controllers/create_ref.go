package controllers

import (
	cerr "ReferralAPI/pkg/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rc *RefController) CreateRef(c *gin.Context) {
	token := c.GetHeader("Authorization")

	claims, err := rc.refServ.TokenValidation(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": cerr.ErrNotAuthorized.Error()})
		return
	}

	refCode, err := rc.refServ.CreateRefToken(claims.Subject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "referral code created", "referral_code": refCode})
}
