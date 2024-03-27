package listactions

import (
	"context"

	"github.com/alemelomeza/poceda/internal/action/domain"
)

type ResponseDTO struct {
	CaseID     int    `json:"case_id"`
	LDAP       string `json:"ldap"`
	ActionType string `json:"action"`
	SLA        string `json:"sla,omitempty"`
	Comment    string `json:"comment"`
	CreatedAt  string `json:"created_at"`
}

type ListActionsUsecase interface {
	Execute(context.Context, int) ([]ResponseDTO, error)
}

type usecase struct {
	repo domain.Repository
}

func New(repo domain.Repository) ListActionsUsecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Execute(ctx context.Context, caseID int) ([]ResponseDTO, error) {
	actions, err := uc.repo.List(ctx, caseID)
	if err != nil {
		return nil, err
	}
	var response []ResponseDTO
	for _, action := range actions {
		response = append(response, ResponseDTO{
			CaseID:     action.CaseID,
			LDAP:       action.LDAP,
			ActionType: action.ActionType,
			SLA:        action.SLA.String(),
			Comment:    action.Comment,
			CreatedAt:  action.CreatedAt,
		})
	}
	return response, nil
}
