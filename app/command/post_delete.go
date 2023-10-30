package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/account"
	"github.com/turistikrota/service.post/domains/post"
)

type PostDeleteCmd struct {
	Account  account.Entity `json:"-"`
	PostUUID string         `json:"-"`
}

type PostDeleteRes struct{}

type PostDeleteHandler cqrs.HandlerFunc[PostDeleteCmd, *PostDeleteRes]

func NewPostDeleteHandler(factory post.Factory, repo post.Repository, events post.Events) PostDeleteHandler {
	return func(ctx context.Context, cmd PostDeleteCmd) (*PostDeleteRes, *i18np.Error) {
		err := repo.Delete(ctx, cmd.PostUUID)
		if err != nil {
			return nil, err
		}
		events.Deleted(post.DeletedEvent{
			UUID: cmd.PostUUID,
			Account: post.AccountEvent{
				UUID: cmd.Account.UUID,
				Name: cmd.Account.Name,
			},
		})
		return &PostDeleteRes{}, nil
	}
}
