package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.listing/domains/account"
	"github.com/turistikrota/service.listing/domains/listing"
)

type ListingValidatedCmd struct {
	New     *listing.Entity `json:"-"`
	Account account.Entity  `json:"-"`
}

type ListingValidatedRes struct{}

type ListingValidatedHandler cqrs.HandlerFunc[ListingValidatedCmd, *ListingValidatedRes]

func NewListingValidatedHandler(repo listing.Repository) ListingValidatedHandler {
	return func(ctx context.Context, cmd ListingValidatedCmd) (*ListingValidatedRes, *i18np.Error) {
		err := repo.Update(ctx, cmd.New)
		if err != nil {
			return nil, err
		}
		return &ListingValidatedRes{}, nil
	}
}
