package repository

import (
	"database/sql"
	"errors"
	"miiboard-service/model"
	"time"
)

// DashboardRepository ...
type DashboardRepository interface {
	GetByID(id string) (model.Dashboard, error)
	// GetAll() ([]model.Dashboard, error)
	// Save(model.Dashboard) (int64, error)
	// Delete(model.Dashboard) error
}

type dashboardRepository struct {
	db *sql.DB
}

// NewDashboardRepository ...
func NewDashboardRepository(db *sql.DB) DashboardRepository {
	return &dashboardRepository{db}
}

func (repo dashboardRepository) GetByID(id string) (model.Dashboard, error) {
	rows, err := repo.db.Query("SELECT id, title, created_at, updated_at FROM dashboards WHERE id=?", id)
	if err != nil {
		return model.Dashboard{}, err
	}
	defer rows.Close()

	if rows.Next() {
		var (
			id        string
			title     string
			createdAt time.Time
			updatedAt time.Time
		)
		err = rows.Scan(&id, &title, &createdAt, &updatedAt)
		if err != nil {
			return model.Dashboard{}, err
		}
		return model.Dashboard{
			ID:        id,
			Title:     title,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}, nil
	}
	return model.Dashboard{}, errors.New("No dashboard found")
}
