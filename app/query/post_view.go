package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/post"
)

type PostViewQuery struct {
	Locale string `json:"-"`
	Slug   string `json:"slug" params:"slug" validate:"required,slug"`
}

type PostViewRes struct {
	*post.DetailDto
}

type PostViewHandler cqrs.HandlerFunc[PostViewQuery, *PostViewRes]

func NewPostViewHandler(repo post.Repository, cacheSrv cache.Service) PostViewHandler {
	cache := cache.New[*post.Entity](cacheSrv)

	createCacheEntity := func() *post.Entity {
		return &post.Entity{}
	}
	return func(ctx context.Context, query PostViewQuery) (*PostViewRes, *i18np.Error) {
		cacheHandler := func() (*post.Entity, *i18np.Error) {
			return repo.View(ctx, post.I18nDetail{
				Locale: query.Locale,
				Slug:   query.Slug,
			})
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("post:slug:%s:locale:%s", query.Slug, query.Locale))
		if err != nil {
			return nil, err
		}
		return &PostViewRes{
			DetailDto: res.ToDetail(),
		}, nil
	}
}
