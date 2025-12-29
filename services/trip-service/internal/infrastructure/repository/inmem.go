package repository

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
)

type Inmemrepo struct{}

func NewInmemRepository() Inmemrepo {
	return Inmemrepo{}
}

func (r *Inmemrepo) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	return nil, nil
}
