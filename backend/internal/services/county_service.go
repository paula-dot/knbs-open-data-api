package services

import (
	"context"
	"fmt"

	"github.com/paula-dot/knbs-open-data-api/backend/internal/database"
)

type CountyService interface {
	GetAllCounties(ctx context.Context) ([]database.County, error)
	GetCountyByID(ctx context.Context, id int32) (database.County, error)
}

type countyService struct {
	db database.Querier
}

func NewCountyService(db database.Querier) CountyService {
	return &countyService{
		db: db,
	}
}

func (s *countyService) GetAllCounties(ctx context.Context) ([]database.County, error) {
	counties, err := s.db.ListCounties(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get counties: %w", err)
	}
	return counties, nil
}

func (s *countyService) GetCountyByID(ctx context.Context, id int32) (database.County, error) {
	// sqlc generated method is likely GetCounty(ctx, id) — delegate to it
	county, err := s.db.GetCounty(ctx, id)
	if err != nil {
		return database.County{}, fmt.Errorf("failed to get county by ID: %w", err)
	}
	return county, nil
}
