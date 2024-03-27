package handler

import (
	"context"
	"time"

	"github.com/alemelomeza/poceda/internal/action/domain"
	saveaction "github.com/alemelomeza/poceda/internal/action/usecase/save_action"
	deletesubscription "github.com/alemelomeza/poceda/internal/subscription/usecase/delete_subscription"
	"github.com/alemelomeza/poceda/pkg/events"
	"github.com/influxdata/influxdb-client-go/v2/domain"
)

type SubscriptionDeletedHandler struct {
	saveActionUC saveaction.SaveActionUsecase
}

func NewSubscriptionDeletedHandler(saveActionUC saveaction.SaveActionUsecase) *SubscriptionDeletedHandler {
	return &SubscriptionDeletedHandler{
		saveActionUC: saveActionUC,
	}
}

func (h *SubscriptionDeletedHandler) Handle(event events.Event) {
	i := event.GetPayload()
	if input, ok := i.(deletesubscription.InputDTO); ok {
		action := domain.Action{
			CaseID:     input.CaseID,
			LDAP:       input.LDAP,
			ActionType: "unsubscribe",
			SLA:        time.Time{},
			Comment:    "",
			CreatedAt:  time.Now().String(),
		}
		h.saveActionUC.Execute(context.Background(), action)
	}
}
