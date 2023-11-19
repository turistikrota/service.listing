package event_stream

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/turistikrota/service.listing/app/command"
	"github.com/turistikrota/service.listing/domains/account"
	"github.com/turistikrota/service.listing/domains/category"
)

func (s srv) OnListingValidationSuccess(data []byte) {
	fmt.Println("OnListingValidationSuccess")
	e := category.ValidationSuccessEvent{}
	err := json.Unmarshal(data, &e)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, _ = s.app.Commands.ListingValidated(context.Background(), command.ListingValidatedCmd{
		New: e.Listing,
		Account: account.Entity{
			UUID: e.User.UUID,
			Name: e.User.Name,
		},
	})
}

func (s srv) OnBookingValidationStart(data []byte) {
	fmt.Println("OnBookingValidationStart")
	cmd := command.ListingValidateBookingCmd{}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, _ = s.app.Commands.BookingValidate(context.Background(), cmd)
}
