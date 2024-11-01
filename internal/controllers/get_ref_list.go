package controllers

import (
	cerr "ReferralAPI/pkg/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rc *RefController) GetRefList(c *gin.Context) {
	token := c.GetHeader("Authorization")

	_, err := rc.refServ.TokenValidation(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": cerr.ErrNotAuthorized.Error()})
		return
	}

	id := c.Param("id")
	users, err := rc.refServ.GetRefList(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "got referrals list", "referrals": users})
}
