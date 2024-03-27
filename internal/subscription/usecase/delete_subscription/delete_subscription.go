package deletesubscription

import (
	"context"

	"github.com/alemelomeza/poceda/internal/shared/event"
	"github.com/alemelomeza/poceda/internal/subscription/domain"
	"github.com/alemelomeza/poceda/pkg/events"
)

type InputDTO struct {
	CaseID int    `json:"case_id"`
	LDAP   string `json:"ldap"`
}

type DeleteSubcriptionUsecase interface {
	Execute(context.Context, int, string) error
}

type usecase struct {
	repo                domain.Repository
	dispatcher          events.EventDispatcher
	subscriptionDeleted event.SubscriptionDeleted
}

func New(
	repo domain.Repository,
	dispatcher events.EventDispatcher,
	subscriptionDeleted event.SubscriptionDeleted,
) DeleteSubcriptionUsecase {
	return &usecase{
		repo:                repo,
		dispatcher:          dispatcher,
		subscriptionDeleted: subscriptionDeleted,
	}
}

func (uc *usecase) Execute(ctx context.Context, caseID int, ldap string) error {
	err := uc.repo.Delete(ctx, caseID, ldap)
	if err != nil {
		return err
	}
	uc.subscriptionDeleted.SetPayload(InputDTO{
		CaseID: caseID,
		LDAP:   ldap,
	})
	uc.dispatcher.Dispatch(&uc.subscriptionDeleted)
	return nil
}
