package query

import (
	"context"
	"fmt"

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
	*list.Result[*post.AdminListDto]
}

type PostListMyHandler cqrs.HandlerFunc[PostListMyQuery, *PostListMyRes]

func NewPostListMyHandler(repo post.Repository) PostListMyHandler {
	return func(ctx context.Context, query PostListMyQuery) (*PostListMyRes, *i18np.Error) {
		fmt.Println("list my inside 1")
		query.Default()
		fmt.Println("list my inside 2")
		offset := (*query.Page - 1) * *query.Limit
		fmt.Println("list my inside 3")
		res, err := repo.ListMy(ctx, query.OwnerUUID, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		fmt.Println("list my inside 4")
		if err != nil {
			return nil, err
		}
		fmt.Println("list my inside 5")
		li := make([]*post.AdminListDto, len(res.List))
		for i, v := range res.List {
			li[i] = v.ToAdminList()
		}
		fmt.Println("list my inside 6")
		result := &list.Result[*post.AdminListDto]{
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
