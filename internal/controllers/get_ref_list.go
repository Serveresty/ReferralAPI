package controllers

import (
	"ReferralAPI/models"
	cerr "ReferralAPI/pkg/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get referral list
// @Description Retrieves the list of referrals for a given user ID
// @Tags referral
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "User ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Router /ref-list/{:id} [get]
func (rc *RefController) GetRefList(c *gin.Context) {
	token := c.GetHeader("Authorization")

	_, err := rc.refServ.TokenValidation(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: cerr.ErrNotAuthorized.Error()})
		return
	}

	id := c.Param("id")
	users, err := rc.refServ.GetRefList(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "got referrals list", Data: users})
}
