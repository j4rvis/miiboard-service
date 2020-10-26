package repository

import "miiboard-service/model"

// DashboardRepository ...
type DashboardRepository interface {
	GetByID(id int64) (model.Dashboard, error)
	GetAll() ([]model.Dashboard, error)
	Save(model.Dashboard) (model.Dashboard, error)
	Delete(model.Dashboard) error
}
