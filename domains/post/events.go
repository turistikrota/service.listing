package post

import (
	"github.com/cilloparch/cillop/events"
	"github.com/turistikrota/service.post/config"
)

type Events interface {
	Created(event CreatedEvent)
	Updated(event UpdatedEvent)
	Deleted(event DeletedEvent)
	Disabled(event DisabledEvent)
	Enabled(event EnabledEvent)
	ReOrder(event ReOrderEvent)
	Restore(event RestoreEvent)
}

type (
	CreatedEvent struct {
		Entity *Entity   `json:"entity"`
		User   UserEvent `json:"user"`
	}
	UpdatedEvent struct {
		Old    *Entity   `json:"old"`
		Entity *Entity   `json:"entity"`
		User   UserEvent `json:"user"`
	}
	DeletedEvent struct {
		UUID string    `json:"uuid"`
		User UserEvent `json:"user"`
	}
	DisabledEvent struct {
		UUID string    `json:"uuid"`
		User UserEvent `json:"user"`
	}
	EnabledEvent struct {
		UUID string    `json:"uuid"`
		User UserEvent `json:"user"`
	}
	ReOrderEvent struct {
		UUID string    `json:"uuid"`
		User UserEvent `json:"user"`
	}
	RestoreEvent struct {
		UUID string    `json:"uuid"`
		User UserEvent `json:"user"`
	}
	UserEvent struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
		Code string `json:"code"`
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
