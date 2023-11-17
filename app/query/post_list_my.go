package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.post/domains/post"
	"github.com/turistikrota/service.post/pkg/utils"
)

type PostListMyQuery struct {
	*utils.Pagination
	OwnerUUID string
}

type PostListMyRes struct {
	*list.Result[*post.OwnerListDto]
}

type PostListMyHandler cqrs.HandlerFunc[PostListMyQuery, *PostListMyRes]

func NewPostListMyHandler(repo post.Repository) PostListMyHandler {
	return func(ctx context.Context, query PostListMyQuery) (*PostListMyRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		res, err := repo.ListMy(ctx, query.OwnerUUID, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		li := make([]*post.OwnerListDto, len(res.List))
		for i, v := range res.List {
			li[i] = v.ToOwnerList()
		}
		result := &list.Result[*post.OwnerListDto]{
			IsNext:        res.IsNext,
			IsPrev:        res.IsPrev,
			FilteredTotal: res.FilteredTotal,
			Total:         res.Total,
			Page:          res.Page,
			List:          li,
		}
		return &PostListMyRes{
			Result: result,
		}, nil
	}
}
