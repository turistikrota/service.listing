package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/account"
	"github.com/turistikrota/service.post/domains/post"
)

type PostReOrderCmd struct {
	Account  account.Entity `json:"-"`
	PostUUID string         `json:"-"`
	Order    *int           `json:"order" validate:"required,min=1,max=100,numeric"`
}

type PostReOrderRes struct{}

type PostReOrderHandler cqrs.HandlerFunc[PostReOrderCmd, *PostReOrderRes]

func NewPostReOrderHandler(factory post.Factory, repo post.Repository, events post.Events) PostReOrderHandler {
	return func(ctx context.Context, cmd PostReOrderCmd) (*PostReOrderRes, *i18np.Error) {
		err := repo.ReOrder(ctx, cmd.PostUUID, *cmd.Order)
		if err != nil {
			return nil, err
		}
		events.ReOrder(post.ReOrderEvent{
			UUID: cmd.PostUUID,
			Account: post.AccountEvent{
				UUID: cmd.Account.UUID,
				Name: cmd.Account.Name,
			},
			NewOrder: *cmd.Order,
		})
		return &PostReOrderRes{}, nil
	}
}
