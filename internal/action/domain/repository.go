package domain

import "context"

type Repository interface {
	List(context.Context, int) ([]Action, error)
	Save(context.Context, Action) error
}
