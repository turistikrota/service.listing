package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/account"
	"github.com/turistikrota/service.post/domains/post"
)

type PostDisableCmd struct {
	Account  account.Entity `json:"-"`
	PostUUID string         `json:"-"`
}

type PostDisableRes struct{}

type PostDisableHandler cqrs.HandlerFunc[PostDisableCmd, *PostDisableRes]

func NewPostDisableHandler(factory post.Factory, repo post.Repository, events post.Events) PostDisableHandler {
	return func(ctx context.Context, cmd PostDisableCmd) (*PostDisableRes, *i18np.Error) {
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
