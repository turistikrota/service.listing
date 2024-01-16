package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.listing/domains/listing"
)

type ListingDeleteCmd struct {
	ListingUUID string `json:"-"`
}

type ListingDeleteRes struct{}

type ListingDeleteHandler cqrs.HandlerFunc[ListingDeleteCmd, *ListingDeleteRes]

func NewListingDeleteHandler(repo listing.Repository, events listing.Events) ListingDeleteHandler {
	return func(ctx context.Context, cmd ListingDeleteCmd) (*ListingDeleteRes, *i18np.Error) {
		res, _, err := repo.GetByUUID(ctx, cmd.ListingUUID)
		if err != nil {
			return nil, err
		}
		err = repo.Delete(ctx, cmd.ListingUUID)
		if err != nil {
			return nil, err
		}
		events.Deleted(listing.DeletedEvent{
			UUID:             cmd.ListingUUID,
			BusinessNickName: res.Business.NickName,
		})
		return &ListingDeleteRes{}, nil
	}
}
