package saveaction

import (
	"context"

	"github.com/alemelomeza/poceda/internal/action/domain"
)

type InputDTO struct {
	CaseID     int    `json:"case_id"`
	LDAP       string `json:"ldap"`
	ActionType string `json:"action"`
	SLA        string `json:"sla,omitempty"`
	Comment    string `json:"comment"`
}

type SaveActionUsecase interface {
	Execute(context.Context, InputDTO) error
}

type usecase struct {
	repo domain.Repository
}

func New(repo domain.Repository) SaveActionUsecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Execute(ctx context.Context, input InputDTO) error {
	return nil
}
