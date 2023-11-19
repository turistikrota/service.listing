package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.listing/domains/listing"
)

type ListingViewQuery struct {
	Locale string `json:"-"`
	Slug   string `json:"slug" params:"slug" validate:"required,slug"`
}

type ListingViewRes struct {
	*listing.DetailDto
}

type ListingViewHandler cqrs.HandlerFunc[ListingViewQuery, *ListingViewRes]

func NewListingViewHandler(repo listing.Repository, cacheSrv cache.Service) ListingViewHandler {
	cache := cache.New[*listing.Entity](cacheSrv)

	createCacheEntity := func() *listing.Entity {
		return &listing.Entity{}
	}
	return func(ctx context.Context, query ListingViewQuery) (*ListingViewRes, *i18np.Error) {
		cacheHandler := func() (*listing.Entity, *i18np.Error) {
			return repo.View(ctx, listing.I18nDetail{
				Locale: query.Locale,
				Slug:   query.Slug,
			})
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("listing:slug:%s:locale:%s", query.Slug, query.Locale))
		if err != nil {
			return nil, err
		}
		return &ListingViewRes{
			DetailDto: res.ToDetail(),
		}, nil
	}
}
