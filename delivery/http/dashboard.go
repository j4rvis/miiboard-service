package delivery

import (
	"miiboard-service/usecase"
	"net/http"
)

// DashboardHandler ...
type DashboardHandler struct {
	dashboardUseCase usecase.DashboardUseCase
}

// NewDashboardHandler ...
func NewDashboardHandler(duc usecase.DashboardUseCase) *DashboardHandler {
	return &DashboardHandler{duc}
}

// GetDashboard ...
func (h *DashboardHandler) GetDashboard(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("Hello"))
	return
}
