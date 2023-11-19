package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type PostUpdateBusinessNickNameCmd struct{}

type PostUpdateBusinessNickNameRes struct{}

type PostUpdateBusinessNickNameHandler cqrs.HandlerFunc[PostUpdateBusinessNickNameCmd, *PostUpdateBusinessNickNameRes]

func NewPostUpdateBusinessNickNameHandler() PostUpdateBusinessNickNameHandler {
	return func(ctx context.Context, cmd PostUpdateBusinessNickNameCmd) (*PostUpdateBusinessNickNameRes, *i18np.Error) {
		return &PostUpdateBusinessNickNameRes{}, nil
	}
}
