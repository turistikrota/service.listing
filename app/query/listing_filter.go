package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.listing/domains/listing"
	"github.com/turistikrota/service.listing/pkg/utils"
)

type ListingFilterQuery struct {
	*utils.Pagination
	listing.FilterEntity
}

type ListingFilterRes struct {
	*list.Result[*listing.ListDto]
}

type ListingFilterHandler cqrs.HandlerFunc[ListingFilterQuery, *ListingFilterRes]

func NewListingFilterHandler(repo listing.Repository) ListingFilterHandler {
	return func(ctx context.Context, query ListingFilterQuery) (*ListingFilterRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		res, err := repo.Filter(ctx, query.FilterEntity, list.Config{
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
		return &ListingFilterRes{
			Result: result,
		}, nil
	}
}
