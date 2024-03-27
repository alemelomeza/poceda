package rest

import (
	"encoding/json"
	"net/http"

	deletesubscription "github.com/alemelomeza/poceda/internal/subscription/usecase/delete_subscription"
	getsubscription "github.com/alemelomeza/poceda/internal/subscription/usecase/get_subscription"
	listsubscriptions "github.com/alemelomeza/poceda/internal/subscription/usecase/list_subscriptions"
	savesubscription "github.com/alemelomeza/poceda/internal/subscription/usecase/save_subscription"
)

type Handler struct {
	DeleteUC deletesubscription.DeleteSubcriptionUsecase
	GetUC    getsubscription.GetSubscriptionUsecase
	ListUC   listsubscriptions.ListSubscriptionsUsecase
	SaveUC   savesubscription.SaveSubscriptionUsecase
}

func New(
	deleteUC deletesubscription.DeleteSubcriptionUsecase,
	getUC getsubscription.GetSubscriptionUsecase,
	listUC listsubscriptions.ListSubscriptionsUsecase,
	saveUC savesubscription.SaveSubscriptionUsecase,
) *Handler {
	return &Handler{
		DeleteUC: deleteUC,
		GetUC:    getUC,
		ListUC:   listUC,
		SaveUC:   saveUC,
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.DeleteUC.Execute(r.Context(), 0, "")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	output, err := h.GetUC.Execute(r.Context(), 0, "")
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
	var input savesubscription.InputDTO
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
