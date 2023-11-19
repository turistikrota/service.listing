package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.listing/domains/account"
	"github.com/turistikrota/service.listing/domains/listing"
)

type ListingDisableCmd struct {
	Account     account.Entity `json:"-"`
	ListingUUID string         `json:"-"`
}

type ListingDisableRes struct{}

type ListingDisableHandler cqrs.HandlerFunc[ListingDisableCmd, *ListingDisableRes]

func NewListingDisableHandler(repo listing.Repository, events listing.Events) ListingDisableHandler {
	return func(ctx context.Context, cmd ListingDisableCmd) (*ListingDisableRes, *i18np.Error) {
		err := repo.Disable(ctx, cmd.ListingUUID)
		if err != nil {
			return nil, err
		}
		events.Disabled(listing.DisabledEvent{
			UUID: cmd.ListingUUID,
			Account: listing.AccountEvent{
				UUID: cmd.Account.UUID,
				Name: cmd.Account.Name,
			},
		})
		return &ListingDisableRes{}, nil
	}
}
