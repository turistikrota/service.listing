package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type PostUpdateOwnerNickNameCmd struct{}

type PostUpdateOwnerNickNameRes struct{}

type PostUpdateOwnerNickNameHandler cqrs.HandlerFunc[PostUpdateOwnerNickNameCmd, *PostUpdateOwnerNickNameRes]

func NewPostUpdateOwnerNickNameHandler() PostUpdateOwnerNickNameHandler {
	return func(ctx context.Context, cmd PostUpdateOwnerNickNameCmd) (*PostUpdateOwnerNickNameRes, *i18np.Error) {
		return &PostUpdateOwnerNickNameRes{}, nil
	}
}
