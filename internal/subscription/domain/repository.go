package domain

import "context"

type Repository interface {
	Get(context.Context, int, string) (Subscription, error)
	List(context.Context, int) ([]Subscription, error)
	Save(context.Context, Subscription) error
	Delete(context.Context, int, string) error
}
