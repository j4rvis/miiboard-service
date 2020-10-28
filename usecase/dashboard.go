package usecase

import (
	"miiboard-service/model"
	"miiboard-service/repository"
)

// DashboardUseCase ...
type DashboardUseCase interface {
	GetByID(id string) (model.Dashboard, error)
}

// DashboardUseCase ...
type dashboardUseCase struct {
	repo repository.DashboardRepository
}

// NewDashboardUseCase ...
func NewDashboardUseCase(repository repository.DashboardRepository) DashboardUseCase {
	return dashboardUseCase{repository}
}

// GetByID ...
func (d dashboardUseCase) GetByID(id string) (model.Dashboard, error) {
	dashboard, err := d.repo.GetByID(id)
	if err != nil {
		return model.Dashboard{}, err
	}
	return dashboard, nil
}
