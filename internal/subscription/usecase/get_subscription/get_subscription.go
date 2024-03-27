package getsubscription

import (
	"context"

	"github.com/alemelomeza/poceda/internal/subscription/domain"
)

type InputDTO struct {
	CaseID int    `json:"case_id"`
	LDAP   string `json:"ldap"`
}

type OutputDTO struct {
	CaseID    int    `json:"case_id"`
	LDAP      string `json:"ldap"`
	CreatedAt string `json:"created_at"`
}

type GetSubscriptionUsecase interface {
	Execute(context.Context, int, string) (OutputDTO, error)
}

type usecase struct {
	repo domain.Repository
}

func New(repo domain.Repository) GetSubscriptionUsecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Execute(ctx context.Context, caseID int, ldap string) (OutputDTO, error) {
	output, err := uc.repo.Get(ctx, caseID, ldap)
	if err != nil {
		return OutputDTO{}, err
	}
	response := OutputDTO{
		CaseID:    output.CaseID,
		LDAP:      output.LDAP,
		CreatedAt: output.CreatedAt.String(),
	}
	return response, nil
}
