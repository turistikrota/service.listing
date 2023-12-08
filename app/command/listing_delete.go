package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.listing/domains/account"
	"github.com/turistikrota/service.listing/domains/listing"
)

type ListingDeleteCmd struct {
	Account          account.Entity `json:"-"`
	ListingUUID      string         `json:"-"`
	BusinessNickName string         `json:"-"`
}

type ListingDeleteRes struct{}

type ListingDeleteHandler cqrs.HandlerFunc[ListingDeleteCmd, *ListingDeleteRes]

func NewListingDeleteHandler(repo listing.Repository, events listing.Events) ListingDeleteHandler {
	return func(ctx context.Context, cmd ListingDeleteCmd) (*ListingDeleteRes, *i18np.Error) {
		err := repo.Delete(ctx, cmd.ListingUUID)
		if err != nil {
			return nil, err
		}
		events.Deleted(listing.DeletedEvent{
			UUID: cmd.ListingUUID,
			Account: listing.AccountEvent{
				UUID: cmd.Account.UUID,
				Name: cmd.Account.Name,
			},
			BusinessNickName: cmd.BusinessNickName,
		})
		return &ListingDeleteRes{}, nil
	}
}
