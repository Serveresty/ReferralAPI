package controllers

import (
	"ReferralAPI/models"
	cerr "ReferralAPI/pkg/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a referral code
// @Description Generates a referral code for the authenticated user
// @Tags referral
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Router /referral/create-ref [post]
func (rc *RefController) CreateRef(c *gin.Context) {
	token := c.GetHeader("Authorization")

	claims, err := rc.refServ.TokenValidation(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: cerr.ErrNotAuthorized.Error()})
		return
	}

	refCode, err := rc.refServ.CreateRefToken(claims.Subject)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, models.Response{Message: "referral code created", Data: refCode})
}
