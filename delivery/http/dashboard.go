package delivery

import (
	"net/http"

	"github.com/go-chi/chi"
)

// DashboardHandler ...
type DashboardHandler interface {
	Routes() *chi.Mux
	getDashboard(http.ResponseWriter, *http.Request)
	getAllDashboards(http.ResponseWriter, *http.Request)
	addDashboard(http.ResponseWriter, *http.Request)
	updateDashboard(http.ResponseWriter, *http.Request)
	deleteDashboard(http.ResponseWriter, *http.Request)
}
