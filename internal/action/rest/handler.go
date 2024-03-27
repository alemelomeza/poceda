package rest

import (
	"encoding/json"
	"net/http"

	listactions "github.com/alemelomeza/poceda/internal/action/usecase/list_actions"
	saveaction "github.com/alemelomeza/poceda/internal/action/usecase/save_action"
)

type Handler struct {
	ListUC listactions.ListActionsUsecase
	SaveUC saveaction.SaveActionUsecase
}

func New(
	listUC listactions.ListActionsUsecase,
	saveUC saveaction.SaveActionUsecase,
) *Handler {
	return &Handler{
		ListUC: listUC,
		SaveUC: saveUC,
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	output, err := h.ListUC.Execute(r.Context(), 0)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
	var input saveaction.RequestDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = h.SaveUC.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
