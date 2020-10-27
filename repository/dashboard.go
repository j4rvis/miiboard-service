package repository

import (
	"database/sql"
	"miiboard-service/model"
)

// DashboardRepository ...
type DashboardRepository interface {
	GetByID(id int64) (model.Dashboard, error)
	GetAll() ([]model.Dashboard, error)
	Save(model.Dashboard) (model.Dashboard, error)
	Delete(model.Dashboard) error
}

type dashboardRepository struct {
	db *sql.DB
}

/* func (repo dashboardRepository) GetByID(id int64) (model.Dashboard, error) {

}
*/
