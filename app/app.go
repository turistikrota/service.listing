package app

import (
	"github.com/turistikrota/service.listing/app/command"
	"github.com/turistikrota/service.listing/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ListingCreate                 command.ListingCreateHandler
	ListingUpdate                 command.ListingUpdateHandler
	ListingValidated              command.ListingValidatedHandler
	ListingUpdateBusinessNickName command.ListingUpdateBusinessNickNameHandler
	ListingEnable                 command.ListingEnableHandler
	ListingDisable                command.ListingDisableHandler
	ListingDelete                 command.ListingDeleteHandler
	ListingRestore                command.ListingRestoreHandler
	ListingReOrder                command.ListingReOrderHandler
	BookingValidate               command.ListingValidateBookingHandler
}

type Queries struct {
	ListingView             query.ListingViewHandler
	ListingAdminView        query.ListingAdminViewHandler
	ListingAdminFilter      query.ListingAdminFilterHandler
	ListingBusinessView     query.ListingBusinessViewHandler
	ListingFilterByBusiness query.ListingFilterByBusinessHandler
	ListingFilter           query.ListingFilterHandler
	ListingListMy           query.ListingListMyHandler
}
