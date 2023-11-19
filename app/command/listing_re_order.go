package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.listing/domains/account"
	"github.com/turistikrota/service.listing/domains/listing"
)

type ListingReOrderCmd struct {
	Account     account.Entity `json:"-"`
	ListingUUID string         `json:"-"`
	Order       *int           `json:"order" validate:"required,min=-1,number"`
}

type ListingReOrderRes struct{}

type ListingReOrderHandler cqrs.HandlerFunc[ListingReOrderCmd, *ListingReOrderRes]

func NewListingReOrderHandler(repo listing.Repository, events listing.Events) ListingReOrderHandler {
	return func(ctx context.Context, cmd ListingReOrderCmd) (*ListingReOrderRes, *i18np.Error) {
		err := repo.ReOrder(ctx, cmd.ListingUUID, *cmd.Order)
		if err != nil {
			return nil, err
		}
		events.ReOrder(listing.ReOrderEvent{
			UUID: cmd.ListingUUID,
			Account: listing.AccountEvent{
				UUID: cmd.Account.UUID,
				Name: cmd.Account.Name,
			},
			NewOrder: *cmd.Order,
		})
		return &ListingReOrderRes{}, nil
	}
}
