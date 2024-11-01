package controllers

import (
	"ReferralAPI/models"
	cerr "ReferralAPI/pkg/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get referral token
// @Description Retrieves the referral token associated with the specified email address
// @Tags referral
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param email query string true "User email"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Router /referral/token [get]
func (rc *RefController) GetToken(c *gin.Context) {
	token := c.GetHeader("Authorization")

	_, err := rc.refServ.TokenValidation(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: cerr.ErrNotAuthorized.Error()})
		return
	}

	email := c.Query("email")
	refToken, err := rc.refServ.GetToken(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "got referral token", Data: refToken})
}
