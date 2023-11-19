package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.listing/domains/listing"
	"github.com/turistikrota/service.listing/pkg/utils"
)

type ListingListMyQuery struct {
	*utils.Pagination
	BusinessUUID string
}

type ListingListMyRes struct {
	*list.Result[*listing.BusinessListDto]
}

type ListingListMyHandler cqrs.HandlerFunc[ListingListMyQuery, *ListingListMyRes]

func NewListingListMyHandler(repo listing.Repository) ListingListMyHandler {
	return func(ctx context.Context, query ListingListMyQuery) (*ListingListMyRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		res, err := repo.ListMy(ctx, query.BusinessUUID, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		li := make([]*listing.BusinessListDto, len(res.List))
		for i, v := range res.List {
			li[i] = v.ToBusinessList()
		}
		result := &list.Result[*listing.BusinessListDto]{
			IsNext:        res.IsNext,
			IsPrev:        res.IsPrev,
			FilteredTotal: res.FilteredTotal,
			Total:         res.Total,
			Page:          res.Page,
			List:          li,
		}
		return &ListingListMyRes{
			Result: result,
		}, nil
	}
}
