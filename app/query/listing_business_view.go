package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.listing/domains/listing"
)

type ListingBusinessViewQuery struct {
	ListingUUID string `json:"uuid" params:"uuid" validate:"required,object_id"`
}

type ListingBusinessViewRes struct {
	*listing.BusinessDetailDto
}

type ListingBusinessViewHandler cqrs.HandlerFunc[ListingBusinessViewQuery, *ListingBusinessViewRes]

func NewListingBusinessViewHandler(repo listing.Repository) ListingBusinessViewHandler {
	return func(ctx context.Context, query ListingBusinessViewQuery) (*ListingBusinessViewRes, *i18np.Error) {
		res, err := repo.BusinessView(ctx, query.ListingUUID)
		if err != nil {
			return nil, err
		}
		return &ListingBusinessViewRes{
			BusinessDetailDto: res.ToBusinessDetail(),
		}, nil
	}
}
