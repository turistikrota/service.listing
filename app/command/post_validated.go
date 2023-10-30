package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/account"
	"github.com/turistikrota/service.post/domains/post"
)

type PostValidatedCmd struct {
	New     *post.Entity   `json:"-"`
	Account account.Entity `json:"-"`
}

type PostValidatedRes struct{}

type PostValidatedHandler cqrs.HandlerFunc[PostValidatedCmd, *PostValidatedRes]

func NewPostValidatedHandler(factory post.Factory, repo post.Repository, events post.Events) PostValidatedHandler {
	return func(ctx context.Context, cmd PostValidatedCmd) (*PostValidatedRes, *i18np.Error) {
		err := repo.Update(ctx, cmd.New)
		if err != nil {
			return nil, err
		}
		return &PostValidatedRes{}, nil
	}
}
