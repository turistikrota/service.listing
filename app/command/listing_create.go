package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.listing/domains/account"
	"github.com/turistikrota/service.listing/domains/listing"
)

type ListingMetaRequest struct {
	TR *listing.Meta `json:"tr" validate:"required,dive"`
	EN *listing.Meta `json:"en" validate:"required,dive"`
}

type ListingCreateCmd struct {
	Account       account.Entity                      `json:"-"`
	Business      listing.Business                    `json:"-"`
	Images        []listing.Image                     `json:"images" validate:"min=1,max=10,dive,required"`
	Meta          *ListingMetaRequest                 `json:"meta" validate:"required,dive"`
	CategoryUUIDs []string                            `json:"categoryUUIDs" validate:"required,min=1,max=30,dive,required,object_id"`
	Features      []listing.Feature                   `json:"features" validate:"required,min=0,max=30,dive,required"`
	Prices        []listing.ListingPriceValidationDto `json:"prices" validate:"required,min=1,max=100,dive,required"`
	Location      *listing.Location                   `json:"location" validate:"required,dive"`
	Boosts        []listing.Boost                     `json:"boosts" validate:"omitempty,min=0,max=10,dive,required"`
	Validation    *listing.Validation                 `json:"validation" validate:"required,dive"`
	Currency      listing.Currency                    `json:"currency" validate:"required,oneof=TRY USD EUR"`
}

type ListingCreateRes struct {
	UUID string `json:"uuid"`
}

type ListingCreateHandler cqrs.HandlerFunc[ListingCreateCmd, *ListingCreateRes]

func NewListingCreateHandler(factory listing.Factory, repo listing.Repository, events listing.Events) ListingCreateHandler {
	return func(ctx context.Context, cmd ListingCreateCmd) (*ListingCreateRes, *i18np.Error) {
		e := factory.New(listing.NewConfig{
			Business:      cmd.Business,
			Images:        cmd.Images,
			Meta:          factory.CreateSlugs(cmd.Meta.TR, cmd.Meta.EN),
			CategoryUUIDs: cmd.CategoryUUIDs,
			Features:      cmd.Features,
			Prices:        cmd.Prices,
			Location:      *cmd.Location,
			Boosts:        cmd.Boosts,
			Validation:    cmd.Validation,
			Currency:      cmd.Currency,
			ForCreate:     true,
		})
		err := factory.Validate(*e)
		if err != nil {
			return nil, err
		}
		saved, _err := repo.Create(ctx, e)
		if _err != nil {
			return nil, _err
		}
		events.Created(listing.CreatedEvent{
			Entity: saved,
			Account: listing.AccountEvent{
				UUID: cmd.Account.UUID,
				Name: cmd.Account.Name,
			},
			BusinessNickName: cmd.Business.NickName,
		})
		return &ListingCreateRes{
			UUID: saved.UUID,
		}, nil
	}
}
