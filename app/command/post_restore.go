package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/account"
	"github.com/turistikrota/service.post/domains/post"
)

type PostRestoreCommand struct {
	Account  account.Entity `json:"-"`
	PostUUID string         `json:"-"`
}

type PostRestoreRes struct{}

type PostRestoreHandler cqrs.HandlerFunc[PostRestoreCommand, *PostRestoreRes]

func NewPostRestoreHandler(factory post.Factory, repo post.Repository, events post.Events) PostRestoreHandler {
	return func(ctx context.Context, cmd PostRestoreCommand) (*PostRestoreRes, *i18np.Error) {
		err := repo.Restore(ctx, cmd.PostUUID)
		if err != nil {
			return nil, err
		}
		events.Restore(post.RestoreEvent{
			UUID: cmd.PostUUID,
			Account: post.AccountEvent{
				UUID: cmd.Account.UUID,
				Name: cmd.Account.Name,
			},
		})
		return &PostRestoreRes{}, nil
	}
}
