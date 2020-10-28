package delivery

import (
	"encoding/json"
	"fmt"
	"miiboard-service/usecase"
	"net/http"

	"github.com/go-chi/chi"
)

// DashboardHandler ...
type DashboardHandler struct {
	duc usecase.DashboardUseCase
}

// NewDashboardHandler ...
func NewDashboardHandler(duc usecase.DashboardUseCase) *DashboardHandler {
	return &DashboardHandler{duc}
}

// GetDashboard ...
func (h *DashboardHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		errStr := fmt.Sprintf("error getting dashboard, no id given")
		w.Write([]byte(errStr))
		return
	}
	dashboard, err := h.duc.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errStr := fmt.Sprintf("error %v", err)
		w.Write([]byte(errStr))
		return
	}
	result, err := json.Marshal(dashboard)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errStr := fmt.Sprintf("error marshalling response %v", err)
		w.Write([]byte(errStr))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
