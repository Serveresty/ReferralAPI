package controllers

import (
	"ReferralAPI/models"
	cerr "ReferralAPI/pkg/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Delete a referral code
// @Description Deletes the referral code for the authenticated user
// @Tags referral
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Router /referral/delete-ref [delete]
func (rc *RefController) DeleteRef(c *gin.Context) {
	token := c.GetHeader("Authorization")

	claims, err := rc.refServ.TokenValidation(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: cerr.ErrNotAuthorized.Error()})
		return
	}

	err = rc.refServ.DeleteRefToken(claims.Subject)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "referral code removed"})
}
