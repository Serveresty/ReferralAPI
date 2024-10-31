package controllers

import "ReferralAPI/internal/service"

type RefController struct {
	refServ *service.RefService
}

func NewRefController(serv *service.RefService) *RefController {
	return &RefController{refServ: serv}
}
