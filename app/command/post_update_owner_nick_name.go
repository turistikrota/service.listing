package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type PostUpdateOwnerNickNameCommand struct{}

type PostUpdateOwnerNickNameRes struct{}

type PostUpdateOwnerNickNameHandler cqrs.HandlerFunc[PostUpdateOwnerNickNameCommand, *PostUpdateOwnerNickNameRes]

func NewPostUpdateOwnerNickNameHandler() PostUpdateOwnerNickNameHandler {
	return func(ctx context.Context, cmd PostUpdateOwnerNickNameCommand) (*PostUpdateOwnerNickNameRes, *i18np.Error) {
		return &PostUpdateOwnerNickNameRes{}, nil
	}
}
