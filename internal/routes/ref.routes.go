package routes

import "ReferralAPI/internal/controllers"

type RefRouteController struct {
	refCont *controllers.RefController
}

func NewRefRouteController(cont *controllers.RefController) *RefRouteController {
	return &RefRouteController{refCont: cont}
}
