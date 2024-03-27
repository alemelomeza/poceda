package listsubscriptions

import (
	"context"

	"github.com/alemelomeza/poceda/internal/subscription/domain"
)

type OutputDTO struct {
	CaseID    int    `json:"case_id"`
	LDAP      string `json:"ldap"`
	CreatedAt string `json:"created_at"`
}

type ListSubscriptionsUsecase interface {
	Execute(context.Context, int) ([]OutputDTO, error)
}

type usecase struct {
	repo domain.Repository
}

func New(repo domain.Repository) ListSubscriptionsUsecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Execute(ctx context.Context, input int) ([]OutputDTO, error) {
	subscriptons, err := uc.repo.List(ctx, input)
	if err != nil {
		return nil, err
	}
	var response []OutputDTO
	for _, subscription := range subscriptons {
		response = append(response, OutputDTO{
			CaseID:    subscription.CaseID,
			LDAP:      subscription.LDAP,
			CreatedAt: subscription.CreatedAt.String(),
		})
	}
	return response, nil
}
