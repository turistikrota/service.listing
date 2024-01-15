package command

import (
	"context"
	"fmt"
	"time"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.listing/domains/booking"
	"github.com/turistikrota/service.listing/domains/listing"
)

type ListingValidateBookingCmd struct {
	BookingUUID string          `json:"booking_uuid"`
	ListingUUID string          `json:"listing_uuid"`
	People      *booking.People `json:"people"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
}

type ListingValidateBookingRes struct{}

type ListingValidateBookingHandler cqrs.HandlerFunc[ListingValidateBookingCmd, *ListingValidateBookingRes]

func NewListingValidateBookingHandler(factory listing.Factory, repo listing.Repository, events listing.Events) ListingValidateBookingHandler {

	validateDateRange := func(cmd ListingValidateBookingCmd, p *listing.Entity) ([]listing.PricePerDay, *float64, *i18np.Error) {
		days := make([]listing.PricePerDay, 0)
		var totalPrice float64
		for i := cmd.StartDate; i.Before(cmd.EndDate); i = i.AddDate(0, 0, 1) {
			if !p.IsAvailable(i) {
				return nil, nil, i18np.NewError("listing.validation.booking.not_available", i18np.P{
					"date": i.Format("2006-01-02"),
				})
			}
			price := p.GetPrice(i)
			days = append(days, listing.PricePerDay{
				Date:  i,
				Price: price,
			})
			totalPrice += price
		}
		return days, &totalPrice, nil
	}

	validatePeople := func(cmd ListingValidateBookingCmd, p *listing.Entity) *i18np.Error {
		if p.Validation.MinAdult != nil && p.Validation.MaxAdult != nil && cmd.People.Adult < *p.Validation.MinAdult || cmd.People.Adult > *p.Validation.MaxAdult {
			return factory.Errors.ValidateBookingAdult(*p.Validation.MinAdult, *p.Validation.MaxAdult)
		}
		if p.Validation.MinKid != nil && p.Validation.MaxKid != nil && cmd.People.Kid < *p.Validation.MinKid || cmd.People.Kid > *p.Validation.MaxKid {
			return factory.Errors.ValidateBookingKid(*p.Validation.MinKid, *p.Validation.MaxKid)
		}
		if p.Validation.MinBaby != nil && p.Validation.MaxBaby != nil && cmd.People.Baby < *p.Validation.MinBaby || cmd.People.Baby > *p.Validation.MaxBaby {
			return factory.Errors.ValidateBookingBaby(*p.Validation.MinBaby, *p.Validation.MaxBaby)
		}
		return nil
	}

	failEvent := func(field string, err *i18np.Error, p *listing.Entity, cmd ListingValidateBookingCmd) (*ListingValidateBookingRes, *i18np.Error) {
		errors := make([]*booking.ValidationError, 0)
		errors = append(errors, &booking.ValidationError{
			Field:   field,
			Message: err.Key,
			Params:  *err.Params,
		})
		event := listing.BookingValidationFailEvent{
			BookingUUID: cmd.BookingUUID,
			ListingUUID: cmd.ListingUUID,
			Errors:      errors,
		}
		if p != nil {
			event.BusinessName = p.Business.NickName
			event.BusinessUUID = p.Business.UUID
		}
		fmt.Printf("%+v\n", event)
		events.BookingValidationFail(event)
		return nil, err
	}

	return func(ctx context.Context, cmd ListingValidateBookingCmd) (*ListingValidateBookingRes, *i18np.Error) {
		p, exists, err := repo.GetByUUID(ctx, cmd.ListingUUID)
		if err != nil {
			return nil, err
		}
		if !exists {
			return failEvent("listing", factory.Errors.ValidateBookingNotFound(), p, cmd)
		}
		dates, totalPrice, error := validateDateRange(cmd, p)
		if error != nil {
			return failEvent("date", error, p, cmd)
		}
		_err := validatePeople(cmd, p)
		if _err != nil {
			return failEvent("people", _err, p, cmd)
		}
		events.BookingValidationSuccess(listing.BookingValidationSuccessEvent{
			BookingUUID:  cmd.BookingUUID,
			ListingUUID:  cmd.ListingUUID,
			BusinessUUID: p.Business.UUID,
			BusinessName: p.Business.NickName,
			TotalPrice:   *totalPrice,
			PricePerDays: dates,
			Currency:     p.Currency,
		})
		return &ListingValidateBookingRes{}, nil
	}
}
