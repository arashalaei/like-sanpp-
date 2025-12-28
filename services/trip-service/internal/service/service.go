package service

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
)

type Service struct {
	repo repository.InmemRepo
}

func NewService(repo repository.InmemRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateTrip(ctx context.Context, fare domain.RideFareModel) (string, error) {
	return "NOT IMPLEMENTED!", nil
}
