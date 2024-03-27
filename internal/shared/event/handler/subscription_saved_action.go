package handler

import (
	"context"
	"time"

	"github.com/alemelomeza/poceda/internal/action/domain"
	saveaction "github.com/alemelomeza/poceda/internal/action/usecase/save_action"
	savesubscription "github.com/alemelomeza/poceda/internal/subscription/usecase/save_subscription"
	"github.com/alemelomeza/poceda/pkg/events"
)

type SubscriptionSavedHandler struct {
	saveActionUC saveaction.SaveActionUsecase
}

func NewSubscriptionSavedHandler(saveActionUC saveaction.SaveActionUsecase) *SubscriptionSavedHandler {
	return &SubscriptionSavedHandler{
		saveActionUC: saveActionUC,
	}
}

func (h *SubscriptionSavedHandler) Handle(event events.Event) {
	i := event.GetPayload()
	if input, ok := i.(savesubscription.InputDTO); ok {
		sla, _ := time.Parse(time.RFC3339, input.SLA)
		action := domain.Action{
			CaseID:     input.CaseID,
			LDAP:       input.LDAP,
			ActionType: "subscribe",
			SLA:        sla,
			Comment:    input.Comment,
			CreatedAt:  time.Now().String(),
		}
		h.saveActionUC.Execute(context.Background(), action)
	}
}
