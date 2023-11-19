package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.post/domains/post"
	"github.com/turistikrota/service.post/pkg/utils"
)

type PostFilterByBusinessQuery struct {
	*utils.Pagination
	post.FilterEntity
	NickName string `json:"-" params:"nickName" validate:"required"`
}

type PostFilterByBusinessRes struct {
	*list.Result[*post.ListDto]
}

type PostFilterByBusinessHandler cqrs.HandlerFunc[PostFilterByBusinessQuery, *PostFilterByBusinessRes]

func NewPostFilterByBusinessHandler(repo post.Repository) PostFilterByBusinessHandler {
	return func(ctx context.Context, query PostFilterByBusinessQuery) (*PostFilterByBusinessRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		res, err := repo.FilterByBusiness(ctx, query.NickName, query.FilterEntity, list.Config{
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
		return &PostFilterByBusinessRes{
			Result: result,
		}, nil
	}
}
