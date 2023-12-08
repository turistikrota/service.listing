package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.listing/domains/account"
	"github.com/turistikrota/service.listing/domains/listing"
)

type ListingEnableCmd struct {
	Account          account.Entity `json:"-"`
	ListingUUID      string         `json:"-"`
	BusinessNickName string         `json:"-"`
}

type ListingEnableRes struct{}

type ListingEnableHandler cqrs.HandlerFunc[ListingEnableCmd, *ListingEnableRes]

func NewListingEnableHandler(repo listing.Repository, events listing.Events) ListingEnableHandler {
	return func(ctx context.Context, cmd ListingEnableCmd) (*ListingEnableRes, *i18np.Error) {
		err := repo.Enable(ctx, cmd.ListingUUID)
		if err != nil {
			return nil, err
		}
		events.Enabled(listing.EnabledEvent{
			UUID: cmd.ListingUUID,
			Account: listing.AccountEvent{
				UUID: cmd.Account.UUID,
				Name: cmd.Account.Name,
			},
			BusinessNickName: cmd.BusinessNickName,
		})
		return &ListingEnableRes{}, nil
	}
}
