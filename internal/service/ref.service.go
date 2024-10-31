package service

import "ReferralAPI/internal/repository"

type RefService struct {
	refRepo *repository.RefRepository
}

func NewRefService(repo *repository.RefRepository) *RefService {
	return &RefService{refRepo: repo}
}
