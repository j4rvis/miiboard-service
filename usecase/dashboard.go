package usecase

import (
	"miiboard-service/model"
	"miiboard-service/repository"
)

type DashboardUsecase interface {
	GetByID(id int64) (model.Dashboard, error)
}

type dashboardUseCases struct {
	dashboards repository.DashboardRepository
}

func NewDashboardUseCase(repository repository.DashboardRepository) DashboardUsecase {
	return &dashboardUseCases{repository}
}

func (d DashboardUsecase) GetByID(id int64) (model.Dashboard, error) {

}
