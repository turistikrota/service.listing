package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.listing/domains/account"
	"github.com/turistikrota/service.listing/domains/listing"
	"github.com/turistikrota/service.listing/domains/payment"
)

type ListingUpdateCmd struct {
	Account              account.Entity                      `json:"-"`
	Business             listing.Business                    `json:"-"`
	ListingUUID          string                              `json:"-"`
	Images               []listing.Image                     `json:"images" validate:"min=1,max=10,dive,required"`
	Meta                 *ListingMetaRequest                 `json:"meta" validate:"required,dive"`
	CategoryUUIDs        []string                            `json:"categoryUUIDs" validate:"required,min=1,max=30,dive,required,object_id"`
	ExtraPaymentChannels []payment.Channel                   `json:"extraPaymentChannels" bson:"extra_payment_channels" validate:"required,min=1,max=30,dive,required,oneof=at_the_door"`
	Features             []listing.Feature                   `json:"features" validate:"required,min=0,max=30,dive,required"`
	Prices               []listing.ListingPriceValidationDto `json:"prices" validate:"required,min=1,max=100,dive,required"`
	Location             *listing.Location                   `json:"location" validate:"required,dive"`
	Boosts               []listing.Boost                     `json:"boosts" validate:"omitempty,min=0,max=10,dive,required"`
	Validation           *listing.Validation                 `json:"validation" validate:"required,dive"`
	Currency             listing.Currency                    `json:"currency" validate:"required,oneof=TRY USD EUR"`
}

type ListingUpdateRes struct {
}

type ListingUpdateHandler cqrs.HandlerFunc[ListingUpdateCmd, *ListingUpdateRes]

func NewListingUpdateHandler(factory listing.Factory, repo listing.Repository, events listing.Events) ListingUpdateHandler {
	return func(ctx context.Context, cmd ListingUpdateCmd) (*ListingUpdateRes, *i18np.Error) {
		e := factory.New(listing.NewConfig{
			Business:             cmd.Business,
			Images:               cmd.Images,
			Meta:                 factory.CreateSlugs(cmd.Meta.TR, cmd.Meta.EN),
			CategoryUUIDs:        cmd.CategoryUUIDs,
			Features:             cmd.Features,
			ExtraPaymentChannels: cmd.ExtraPaymentChannels,
			Prices:               cmd.Prices,
			Location:             *cmd.Location,
			Boosts:               cmd.Boosts,
			Validation:           cmd.Validation,
			Currency:             cmd.Currency,
			ForCreate:            false,
		})
		err := factory.Validate(*e)
		if err != nil {
			return nil, err
		}
		old, _err := repo.AdminView(ctx, cmd.ListingUUID)
		if _err != nil {
			return nil, _err
		}
		e.CreatedAt = old.CreatedAt
		e.UUID = old.UUID
		e.Order = old.Order
		events.Updated(listing.UpdatedEvent{
			Old:    old,
			Entity: e,
			Account: listing.AccountEvent{
				UUID: cmd.Account.UUID,
				Name: cmd.Account.Name,
			},
			BusinessNickName: cmd.Business.NickName,
		})
		return &ListingUpdateRes{}, nil
	}
}
