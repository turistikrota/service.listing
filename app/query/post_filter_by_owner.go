package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.post/domains/post"
	"github.com/turistikrota/service.post/pkg/utils"
)

type PostFilterByOwnerQuery struct {
	*utils.Pagination
	post.FilterEntity
	NickName string `json:"-" params:"nickName" validate:"required"`
}

type PostFilterByOwnerRes struct {
	*list.Result[*post.ListDto]
}

type PostFilterByOwnerHandler cqrs.HandlerFunc[PostFilterByOwnerQuery, *PostFilterByOwnerRes]

func NewPostFilterByOwnerHandler(repo post.Repository) PostFilterByOwnerHandler {
	return func(ctx context.Context, query PostFilterByOwnerQuery) (*PostFilterByOwnerRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		res, err := repo.FilterByOwner(ctx, query.NickName, query.FilterEntity, list.Config{
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
		return &PostFilterByOwnerRes{
			Result: result,
		}, nil
	}
}
