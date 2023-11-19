package service

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/validation"
	"github.com/turistikrota/service.post/app"
	"github.com/turistikrota/service.post/app/command"
	"github.com/turistikrota/service.post/app/query"
	"github.com/turistikrota/service.post/config"
	"github.com/turistikrota/service.post/domains/post"
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
	postFactory := post.NewFactory()
	postRepo := post.NewRepo(cnf.MongoDB.GetCollection(cnf.App.DB.Post.Collection), postFactory)
	postEvents := post.NewEvents(post.EventConfig{
		Topics:    cnf.App.Topics,
		Publisher: cnf.EventEngine,
	})

	return app.Application{
		Commands: app.Commands{
			PostCreate:                 command.NewPostCreateHandler(postFactory, postRepo, postEvents),
			PostUpdate:                 command.NewPostUpdateHandler(postFactory, postRepo, postEvents),
			PostValidated:              command.NewPostValidatedHandler(postRepo),
			PostUpdateBusinessNickName: command.NewPostUpdateBusinessNickNameHandler(),
			PostEnable:                 command.NewPostEnableHandler(postRepo, postEvents),
			PostDisable:                command.NewPostDisableHandler(postRepo, postEvents),
			PostDelete:                 command.NewPostDeleteHandler(postRepo, postEvents),
			PostRestore:                command.NewPostRestoreHandler(postRepo, postEvents),
			PostReOrder:                command.NewPostReOrderHandler(postRepo, postEvents),
			BookingValidate:            command.NewPostValidateBookingHandler(postFactory, postRepo, postEvents),
		},
		Queries: app.Queries{
			PostView:             query.NewPostViewHandler(postRepo, cnf.CacheSrv),
			PostAdminView:        query.NewPostAdminViewHandler(postRepo),
			PostFilterByBusiness: query.NewPostFilterByBusinessHandler(postRepo),
			PostFilter:           query.NewPostFilterHandler(postRepo),
			PostListMy:           query.NewPostListMyHandler(postRepo),
		},
	}
}
