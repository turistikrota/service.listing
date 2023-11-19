package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type ListingUpdateBusinessNickNameCmd struct{}

type ListingUpdateBusinessNickNameRes struct{}

type ListingUpdateBusinessNickNameHandler cqrs.HandlerFunc[ListingUpdateBusinessNickNameCmd, *ListingUpdateBusinessNickNameRes]

func NewListingUpdateBusinessNickNameHandler() ListingUpdateBusinessNickNameHandler {
	return func(ctx context.Context, cmd ListingUpdateBusinessNickNameCmd) (*ListingUpdateBusinessNickNameRes, *i18np.Error) {
		return &ListingUpdateBusinessNickNameRes{}, nil
	}
}
