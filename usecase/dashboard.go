package usecase

import (
	"miiboard-service/model"
	"miiboard-service/repository"
)

// DashboardUseCase ...
type DashboardUseCase interface {
	GetByID(id int64) (model.Dashboard, error)
}

type dashboardUseCase struct {
	dashboards repository.DashboardRepository
}

// NewDashboardUseCase ...
func NewDashboardUseCase(repository repository.DashboardRepository) DashboardUseCase {
	return &dashboardUseCase{repository}
}

func (d *dashboardUseCase) GetByID(id int64) (model.Dashboard, error) {

}
