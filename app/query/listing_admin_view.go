package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.listing/domains/listing"
)

type ListingAdminViewQuery struct {
	ListingUUID string `json:"uuid" params:"uuid" validate:"required,object_id"`
}

type ListingAdminViewRes struct {
	*listing.AdminDetailDto
}

type ListingAdminViewHandler cqrs.HandlerFunc[ListingAdminViewQuery, *ListingAdminViewRes]

func NewListingAdminViewHandler(repo listing.Repository) ListingAdminViewHandler {
	return func(ctx context.Context, query ListingAdminViewQuery) (*ListingAdminViewRes, *i18np.Error) {
		res, err := repo.AdminView(ctx, query.ListingUUID)
		if err != nil {
			return nil, err
		}
		return &ListingAdminViewRes{
			AdminDetailDto: res.ToAdminDetail(),
		}, nil
	}
}
