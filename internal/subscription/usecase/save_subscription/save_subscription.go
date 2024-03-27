package savesubscription

import (
	"context"
	"time"

	"github.com/alemelomeza/poceda/internal/shared/event"
	"github.com/alemelomeza/poceda/internal/subscription/domain"
	"github.com/alemelomeza/poceda/pkg/events"
)

type InputDTO struct {
	CaseID  int    `json:"case_id"`
	LDAP    string `json:"ldap"`
	SLA     string `json:"sla,omitempty"`
	Comment string `json:"comment"`
}

type SaveSubscriptionUsecase interface {
	Execute(context.Context, InputDTO) error
}

type usecase struct {
	repo              domain.Repository
	dispatcher        events.EventDispatcher
	subscriptionSaved event.SubscriptionSaved
}

func New(
	repo domain.Repository,
	dispatcher events.EventDispatcher,
	subscriptionSaved event.SubscriptionSaved,
) SaveSubscriptionUsecase {
	return &usecase{
		repo:              repo,
		dispatcher:        dispatcher,
		subscriptionSaved: subscriptionSaved,
	}
}

func (uc *usecase) Execute(ctx context.Context, input InputDTO) error {
	subscription := domain.Subscription{
		CaseID:    input.CaseID,
		LDAP:      input.LDAP,
		CreatedAt: time.Now(),
	}
	err := uc.repo.Save(ctx, subscription)
	if err != nil {
		return err
	}
	uc.subscriptionSaved.SetPayload(input)
	uc.dispatcher.Dispatch(&uc.subscriptionSaved)
	return nil
}
