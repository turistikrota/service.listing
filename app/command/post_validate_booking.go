package command

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/booking"
	"github.com/turistikrota/service.post/domains/post"
)

type PostValidateBookingCmd struct {
	BookingUUID string          `json:"booking_uuid"`
	PostUUID    string          `json:"post_uuid"`
	People      *booking.People `json:"people"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
}

type PostValidateBookingRes struct{}

type PostValidateBookingHandler cqrs.HandlerFunc[PostValidateBookingCmd, *PostValidateBookingRes]

func NewPostValidateBookingHandler(factory post.Factory, repo post.Repository, events post.Events) PostValidateBookingHandler {

	validateDateRange := func(cmd PostValidateBookingCmd, p *post.Entity) ([]post.PricePerDay, *float64, *i18np.Error) {
		days := make([]post.PricePerDay, 0)
		var totalPrice float64
		for i := cmd.StartDate; i.Before(cmd.EndDate); i = i.AddDate(0, 0, 1) {
			if !p.IsAvailable(i) {
				return nil, nil, i18np.NewError("post.validation.booking.not_available", i18np.P{
					"date": i.Format("2006-01-02"),
				})
			}
			price := p.GetPrice(i)
			days = append(days, post.PricePerDay{
				Date:  i,
				Price: price,
			})
			totalPrice += price
		}
		return days, &totalPrice, nil
	}

	validatePeople := func(cmd PostValidateBookingCmd, p *post.Entity) *i18np.Error {
		if p.Validation.MinAdult != nil && cmd.People.Adult < *p.Validation.MinAdult || cmd.People.Adult > *p.Validation.MaxAdult {
			return factory.Errors.ValidateBookingAdult(*p.Validation.MinAdult, *p.Validation.MaxAdult)
		}
		if p.Validation.MinKid != nil && cmd.People.Kid < *p.Validation.MinKid || cmd.People.Kid > *p.Validation.MaxKid {
			return factory.Errors.ValidateBookingKid(*p.Validation.MinKid, *p.Validation.MaxKid)
		}
		if p.Validation.MinBaby != nil && cmd.People.Baby < *p.Validation.MinBaby || cmd.People.Baby > *p.Validation.MaxBaby {
			return factory.Errors.ValidateBookingBaby(*p.Validation.MinBaby, *p.Validation.MaxBaby)
		}
		return nil
	}

	failEvent := func(field string, err *i18np.Error, p *post.Entity, cmd PostValidateBookingCmd) (*PostValidateBookingRes, *i18np.Error) {
		errors := make([]*booking.ValidationError, 0)
		errors = append(errors, &booking.ValidationError{
			Field:   field,
			Message: err.Key,
			Params:  *err.Params,
		})
		event := post.BookingValidationFailEvent{
			BookingUUID: cmd.BookingUUID,
			PostUUID:    cmd.PostUUID,
			Errors:      errors,
		}
		if p != nil {
			event.OwnerName = p.Owner.NickName
			event.OwnerUUID = p.Owner.UUID
		}
		events.BookingValidationFail(event)
		return nil, err
	}

	return func(ctx context.Context, cmd PostValidateBookingCmd) (*PostValidateBookingRes, *i18np.Error) {
		p, exists, err := repo.GetByUUID(ctx, cmd.PostUUID)
		if !exists {
			return failEvent("post", factory.Errors.ValidateBookingNotFound(), p, cmd)
		}
		if err != nil {
			return nil, err
		}
		dates, totalPrice, error := validateDateRange(cmd, p)
		if error != nil {
			return failEvent("date", error, p, cmd)
		}
		_err := validatePeople(cmd, p)
		if _err != nil {
			return failEvent("people", _err, p, cmd)
		}
		events.BookingValidationSuccess(post.BookingValidationSuccessEvent{
			BookingUUID:  cmd.BookingUUID,
			PostUUID:     cmd.PostUUID,
			OwnerUUID:    p.Owner.UUID,
			OwnerName:    p.Owner.NickName,
			TotalPrice:   *totalPrice,
			PricePerDays: dates,
		})
		return &PostValidateBookingRes{}, nil
	}
}
