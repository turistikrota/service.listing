package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/account"
	"github.com/turistikrota/service.post/domains/post"
)

type PostDisableCommand struct {
	Account  account.Entity `json:"-"`
	PostUUID string         `json:"-"`
}

type PostDisableRes struct{}

type PostDisableHandler cqrs.HandlerFunc[PostDisableCommand, *PostDisableRes]

func NewPostDisableHandler(factory post.Factory, repo post.Repository, events post.Events) PostDisableHandler {
	return func(ctx context.Context, cmd PostDisableCommand) (*PostDisableRes, *i18np.Error) {
		err := repo.Disable(ctx, cmd.PostUUID)
		if err != nil {
			return nil, err
		}
		events.Disabled(post.DisabledEvent{
			UUID: cmd.PostUUID,
			Account: post.AccountEvent{
				UUID: cmd.Account.UUID,
				Name: cmd.Account.Name,
			},
		})
		return &PostDisableRes{}, nil
	}
}
