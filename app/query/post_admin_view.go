package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/post"
)

type PostAdminViewQuery struct {
	PostUUID string `json:"uuid" params:"uuid" validate:"required,object_id"`
}

type PostAdminViewRes struct {
	*post.AdminDetailDto
}

type PostAdminViewHandler cqrs.HandlerFunc[PostAdminViewQuery, *PostAdminViewRes]

func NewPostAdminViewHandler(repo post.Repository) PostAdminViewHandler {
	return func(ctx context.Context, query PostAdminViewQuery) (*PostAdminViewRes, *i18np.Error) {
		res, err := repo.AdminView(ctx, query.PostUUID)
		if err != nil {
			return nil, err
		}
		return &PostAdminViewRes{
			AdminDetailDto: res.ToAdminDetail(),
		}, nil
	}
}
