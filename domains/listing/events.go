package listing

import (
	"github.com/cilloparch/cillop/events"
	"github.com/turistikrota/service.listing/config"
	"github.com/turistikrota/service.listing/domains/booking"
)

type Events interface {
	Created(event CreatedEvent)
	Updated(event UpdatedEvent)
	Deleted(event DeletedEvent)
	Disabled(event DisabledEvent)
	Enabled(event EnabledEvent)
	ReOrder(event ReOrderEvent)
	Restore(event RestoreEvent)
	BookingValidationSuccess(event BookingValidationSuccessEvent)
	BookingValidationFail(event BookingValidationFailEvent)
}

type (
	CreatedEvent struct {
		Entity  *Entity      `json:"entity"`
		Account AccountEvent `json:"account"`
	}
	UpdatedEvent struct {
		Old     *Entity      `json:"old"`
		Entity  *Entity      `json:"entity"`
		Account AccountEvent `json:"account"`
	}
	DeletedEvent struct {
		UUID    string       `json:"uuid"`
		Account AccountEvent `json:"account"`
	}
	DisabledEvent struct {
		UUID    string       `json:"uuid"`
		Account AccountEvent `json:"account"`
	}
	EnabledEvent struct {
		UUID    string       `json:"uuid"`
		Account AccountEvent `json:"account"`
	}
	ReOrderEvent struct {
		UUID     string       `json:"uuid"`
		Account  AccountEvent `json:"account"`
		NewOrder int          `json:"new_order"`
	}
	RestoreEvent struct {
		UUID    string       `json:"uuid"`
		Account AccountEvent `json:"account"`
	}
	AccountEvent struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}
	BookingValidationSuccessEvent struct {
		BookingUUID  string        `json:"booking_uuid"`
		ListingUUID  string        `json:"listing_uuid"`
		BusinessUUID string        `json:"business_uuid"`
		BusinessName string        `json:"business_name"`
		TotalPrice   float64       `json:"total_price"`
		PricePerDays []PricePerDay `json:"price_per_days"`
	}
	BookingValidationFailEvent struct {
		BookingUUID  string                     `json:"booking_uuid"`
		ListingUUID  string                     `json:"listing_uuid"`
		BusinessUUID string                     `json:"business_uuid"`
		BusinessName string                     `json:"business_name"`
		Errors       []*booking.ValidationError `json:"errors"`
	}
)

type listingEvents struct {
	publisher events.Publisher
	topics    config.Topics
}

type EventConfig struct {
	Topics    config.Topics
	Publisher events.Publisher
}

func NewEvents(cnf EventConfig) Events {
	return &listingEvents{
		publisher: cnf.Publisher,
		topics:    cnf.Topics,
	}
}

func (e *listingEvents) Created(event CreatedEvent) {
	_ = e.publisher.Publish(e.topics.Listing.Created, event)
}

func (e *listingEvents) Updated(event UpdatedEvent) {
	_ = e.publisher.Publish(e.topics.Listing.Updated, event)
}

func (e *listingEvents) Deleted(event DeletedEvent) {
	_ = e.publisher.Publish(e.topics.Listing.Deleted, event)
}

func (e *listingEvents) Disabled(event DisabledEvent) {
	_ = e.publisher.Publish(e.topics.Listing.Disabled, event)
}

func (e *listingEvents) Enabled(event EnabledEvent) {
	_ = e.publisher.Publish(e.topics.Listing.Enabled, event)
}

func (e *listingEvents) ReOrder(event ReOrderEvent) {
	_ = e.publisher.Publish(e.topics.Listing.ReOrdered, event)
}

func (e *listingEvents) Restore(event RestoreEvent) {
	_ = e.publisher.Publish(e.topics.Listing.Restored, event)
}

func (e *listingEvents) BookingValidationSuccess(event BookingValidationSuccessEvent) {
	_ = e.publisher.Publish(e.topics.Booking.ValidationSuccess, event)
}

func (e *listingEvents) BookingValidationFail(event BookingValidationFailEvent) {
	_ = e.publisher.Publish(e.topics.Booking.ValidationFail, event)
}
