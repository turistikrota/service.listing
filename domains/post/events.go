package post

import (
	"github.com/cilloparch/cillop/events"
	"github.com/turistikrota/service.post/config"
	"github.com/turistikrota/service.post/domains/booking"
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
		PostUUID     string        `json:"post_uuid"`
		BusinessUUID string        `json:"business_uuid"`
		BusinessName string        `json:"business_name"`
		TotalPrice   float64       `json:"total_price"`
		PricePerDays []PricePerDay `json:"price_per_days"`
	}
	BookingValidationFailEvent struct {
		BookingUUID  string                     `json:"booking_uuid"`
		PostUUID     string                     `json:"post_uuid"`
		BusinessUUID string                     `json:"business_uuid"`
		BusinessName string                     `json:"business_name"`
		Errors       []*booking.ValidationError `json:"errors"`
	}
)

type postEvents struct {
	publisher events.Publisher
	topics    config.Topics
}

type EventConfig struct {
	Topics    config.Topics
	Publisher events.Publisher
}

func NewEvents(cnf EventConfig) Events {
	return &postEvents{
		publisher: cnf.Publisher,
		topics:    cnf.Topics,
	}
}

func (e *postEvents) Created(event CreatedEvent) {
	_ = e.publisher.Publish(e.topics.Post.Created, event)
}

func (e *postEvents) Updated(event UpdatedEvent) {
	_ = e.publisher.Publish(e.topics.Post.Updated, event)
}

func (e *postEvents) Deleted(event DeletedEvent) {
	_ = e.publisher.Publish(e.topics.Post.Deleted, event)
}

func (e *postEvents) Disabled(event DisabledEvent) {
	_ = e.publisher.Publish(e.topics.Post.Disabled, event)
}

func (e *postEvents) Enabled(event EnabledEvent) {
	_ = e.publisher.Publish(e.topics.Post.Enabled, event)
}

func (e *postEvents) ReOrder(event ReOrderEvent) {
	_ = e.publisher.Publish(e.topics.Post.ReOrdered, event)
}

func (e *postEvents) Restore(event RestoreEvent) {
	_ = e.publisher.Publish(e.topics.Post.Restored, event)
}

func (e *postEvents) BookingValidationSuccess(event BookingValidationSuccessEvent) {
	_ = e.publisher.Publish(e.topics.Booking.ValidationSuccess, event)
}

func (e *postEvents) BookingValidationFail(event BookingValidationFailEvent) {
	_ = e.publisher.Publish(e.topics.Booking.ValidationFail, event)
}
