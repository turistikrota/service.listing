package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/account"
	"github.com/turistikrota/service.post/domains/post"
)

type PostEnableCmd struct {
	Account  account.Entity `json:"-"`
	PostUUID string         `json:"-"`
}

type PostEnableRes struct{}

type PostEnableHandler cqrs.HandlerFunc[PostEnableCmd, *PostEnableRes]

func NewPostEnableHandler(factory post.Factory, repo post.Repository, events post.Events) PostEnableHandler {
	return func(ctx context.Context, cmd PostEnableCmd) (*PostEnableRes, *i18np.Error) {
		err := repo.Enable(ctx, cmd.PostUUID)
		if err != nil {
			return nil, err
		}
		events.Enabled(post.EnabledEvent{
			UUID: cmd.PostUUID,
			Account: post.AccountEvent{
				UUID: cmd.Account.UUID,
				Name: cmd.Account.Name,
			},
		})
		return &PostEnableRes{}, nil
	}
}
