package controllers

import (
	"github.com/gin-gonic/gin"
)

func (rc *RefController) CreateRef(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if _, err := rc.refServ.TokenValidation(token); err != nil {

	}
}
