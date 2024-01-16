package service

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/validation"
	"github.com/turistikrota/service.listing/app"
	"github.com/turistikrota/service.listing/app/command"
	"github.com/turistikrota/service.listing/app/query"
	"github.com/turistikrota/service.listing/config"
	"github.com/turistikrota/service.listing/domains/listing"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Validator   *validation.Validator
	MongoDB     *mongo.DB
	CacheSrv    cache.Service
}

func NewApplication(cnf Config) app.Application {
	listingFactory := listing.NewFactory()
	listingRepo := listing.NewRepo(cnf.MongoDB.GetCollection(cnf.App.DB.Listing.Collection), listingFactory)
	listingEvents := listing.NewEvents(listing.EventConfig{
		Topics:    cnf.App.Topics,
		Publisher: cnf.EventEngine,
	})

	return app.Application{
		Commands: app.Commands{
			ListingCreate:                 command.NewListingCreateHandler(listingFactory, listingRepo, listingEvents),
			ListingUpdate:                 command.NewListingUpdateHandler(listingFactory, listingRepo, listingEvents),
			ListingValidated:              command.NewListingValidatedHandler(listingRepo),
			ListingUpdateBusinessNickName: command.NewListingUpdateBusinessNickNameHandler(),
			ListingEnable:                 command.NewListingEnableHandler(listingRepo, listingEvents),
			ListingDisable:                command.NewListingDisableHandler(listingRepo, listingEvents),
			ListingDelete:                 command.NewListingDeleteHandler(listingRepo, listingEvents),
			ListingRestore:                command.NewListingRestoreHandler(listingRepo, listingEvents),
			ListingReOrder:                command.NewListingReOrderHandler(listingRepo, listingEvents),
			BookingValidate:               command.NewListingValidateBookingHandler(listingFactory, listingRepo, listingEvents),
		},
		Queries: app.Queries{
			ListingView:             query.NewListingViewHandler(listingRepo, cnf.CacheSrv),
			ListingAdminView:        query.NewListingAdminViewHandler(listingRepo),
			ListingBusinessView:     query.NewListingBusinessViewHandler(listingRepo),
			ListingFilterByBusiness: query.NewListingFilterByBusinessHandler(listingRepo),
			ListingFilter:           query.NewListingFilterHandler(listingRepo),
			ListingListMy:           query.NewListingListMyHandler(listingRepo),
		},
	}
}
