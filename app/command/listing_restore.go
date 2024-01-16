package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.listing/domains/listing"
)

type ListingRestoreCmd struct {
	ListingUUID string `json:"-"`
}

type ListingRestoreRes struct{}

type ListingRestoreHandler cqrs.HandlerFunc[ListingRestoreCmd, *ListingRestoreRes]

func NewListingRestoreHandler(repo listing.Repository, events listing.Events) ListingRestoreHandler {
	return func(ctx context.Context, cmd ListingRestoreCmd) (*ListingRestoreRes, *i18np.Error) {
		res, err := repo.AdminView(ctx, cmd.ListingUUID)
		if err != nil {
			return nil, err
		}
		err = repo.Restore(ctx, cmd.ListingUUID)
		if err != nil {
			return nil, err
		}
		events.Restore(listing.RestoreEvent{
			UUID:             cmd.ListingUUID,
			BusinessNickName: res.Business.NickName,
		})
		return &ListingRestoreRes{}, nil
	}
}
