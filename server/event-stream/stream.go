package event_stream

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/server"
	"github.com/turistikrota/service.post/app"
	"github.com/turistikrota/service.post/config"
)

type srv struct {
	app    app.Application
	topics config.Topics
	engine events.Engine
}

type Config struct {
	App    app.Application
	Engine events.Engine
	Topics config.Topics
}

func New(config Config) server.Server {
	return srv{
		app:    config.App,
		engine: config.Engine,
		topics: config.Topics,
	}
}

func (s srv) Listen() error {
	err := s.engine.Subscribe(s.topics.Category.PostValidationSuccess, s.OnPostValidationSuccess)
	if err != nil {
		return err
	}
	err = s.engine.Subscribe(s.topics.Booking.ValidationStart, s.OnBookingValidationStart)
	return err
}
