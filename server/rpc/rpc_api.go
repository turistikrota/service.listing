package rpc

import (
	"context"

	"github.com/turistikrota/service.listing/app/query"
	"github.com/turistikrota/service.listing/domains/listing"
	protos "github.com/turistikrota/service.listing/protos"
)

func (s srv) GetEntity(ctx context.Context, req *protos.GetEntityRequest) (*protos.Entity, error) {
	res, err := s.app.Queries.ListingAdminView(ctx, query.ListingAdminViewQuery{
		ListingUUID: req.Uuid,
	})
	if err != nil {
		return nil, err
	}
	images := make([]*protos.Image, len(res.Images))
	for i, img := range res.Images {
		var order int32
		if img.Order != nil {
			order = int32(*img.Order)
		}
		images[i] = &protos.Image{
			Url:   img.Url,
			Order: order,
		}
	}
	entity := &protos.Entity{
		Uuid:         res.UUID,
		BusinessUuid: res.Business.UUID,
		Images:       images,
		BusinessName: res.Business.NickName,
		CityName:     res.Location.City,
		DistrictName: res.Location.Street,
		CountryName:  res.Location.Country,
	}
	locale := listing.LocaleTR
	if req.Locale == listing.LocaleEN.String() {
		locale = listing.LocaleEN
	}
	meta := res.Meta[locale]
	entity.Title = meta.Title
	entity.Description = meta.Description
	entity.Slug = meta.Slug
	return entity, nil
}
