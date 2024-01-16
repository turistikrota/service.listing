package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.listing/domains/listing"
	"github.com/turistikrota/service.listing/pkg/utils"
)

type ListingAdminFilterQuery struct {
	*utils.Pagination
	listing.FilterEntity
}

type ListingAdminFilterRes struct {
	*list.Result[*listing.ListDto]
}

type ListingAdminFilterHandler cqrs.HandlerFunc[ListingAdminFilterQuery, *ListingAdminFilterRes]

func NewListingAdminFilterHandler(repo listing.Repository) ListingAdminFilterHandler {
	return func(ctx context.Context, query ListingAdminFilterQuery) (*ListingAdminFilterRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		res, err := repo.AdminFilter(ctx, query.FilterEntity, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		li := make([]*listing.ListDto, len(res.List))
		for i, v := range res.List {
			li[i] = v.ToList()
		}
		result := &list.Result[*listing.ListDto]{
			IsNext:        res.IsNext,
			IsPrev:        res.IsPrev,
			FilteredTotal: res.FilteredTotal,
			Total:         res.Total,
			Page:          res.Page,
			List:          li,
		}
		return &ListingAdminFilterRes{
			Result: result,
		}, nil
	}
}
