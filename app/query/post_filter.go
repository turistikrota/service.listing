package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.post/domains/post"
	"github.com/turistikrota/service.post/pkg/utils"
)

type PostFilterQuery struct {
	*utils.Pagination
	post.FilterEntity
}

type PostFilterRes struct {
	*list.Result[*post.ListDto]
}

type PostFilterHandler cqrs.HandlerFunc[PostFilterQuery, *PostFilterRes]

func NewPostFilterHandler(repo post.Repository) PostFilterHandler {
	return func(ctx context.Context, query PostFilterQuery) (*PostFilterRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		res, err := repo.Filter(ctx, query.FilterEntity, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		li := make([]*post.ListDto, len(res.List))
		for i, v := range res.List {
			li[i] = v.ToList()
		}
		result := &list.Result[*post.ListDto]{
			IsNext:        res.IsNext,
			IsPrev:        res.IsPrev,
			FilteredTotal: res.FilteredTotal,
			Total:         res.Total,
			Page:          res.Page,
			List:          li,
		}
		return &PostFilterRes{
			Result: result,
		}, nil
	}
}
