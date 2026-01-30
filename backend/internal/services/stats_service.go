package services

import (
	"context"

	"github.com/paula-dot/knbs-open-data-api/backend/internal/database"
)

type StatsService interface {
	GetIndicators(ctx context.Context) ([]database.Indicator, error)
	GetData(ctx context.Context, indicatorCode string, year int32) ([]database.GetDataByIndicatorRow, error)
}

type statsService struct {
	db database.Querier
}

func NewStatsService(db database.Querier) StatsService {
	return &statsService{db: db}
}

func (s *statsService) GetIndicators(ctx context.Context) ([]database.Indicator, error) {
	return s.db.ListIndicators(ctx)
}

func (s *statsService) GetData(ctx context.Context, indicatorCode string, year int32) ([]database.GetDataByIndicatorRow, error) {
	params := database.GetDataByIndicatorParams{
		Code: indicatorCode,
		Year: year,
	}
	return s.db.GetDataByIndicator(ctx, params)
}
