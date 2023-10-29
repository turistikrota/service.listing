package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/account"
	"github.com/turistikrota/service.post/domains/post"
)

type PostCreateCommand struct {
	Account       account.Entity                `json:"-"`
	Owner         post.Owner                    `json:"-"`
	Images        []post.Image                  `json:"images" validate:"min=1,max=10,dive,required"`
	Meta          map[post.Locale]post.Meta     `json:"meta" validate:"required,dive,required"`
	CategoryUUIDs []string                      `json:"categoryUUIDs" validate:"required,min=1,max=30,dive,required,object_id"`
	Features      []post.Feature                `json:"features" validate:"required,min=1,max=30,dive,required"`
	Prices        []post.PostPriceValidationDto `json:"prices" validate:"required,min=1,max=100,dive,required"`
	Location      post.Location                 `json:"location" validate:"required"`
	Boosts        []post.Boost                  `json:"boosts" validate:"omitempty,min=1,max=10,dive,required"`
	Type          post.Type                     `json:"type" validate:"required"`
	People        post.People                   `json:"people" validate:"required"`
	Count         *int                          `json:"count" validate:"required,min=1,max=100,numeric"`
}

type PostCreateRes struct {
	UUID string `json:"uuid"`
}

type PostCreateHandler cqrs.HandlerFunc[PostCreateCommand, *PostCreateRes]

func NewPostCreateHandler(factory post.Factory, repo post.Repository, events post.Events) PostCreateHandler {
	return func(ctx context.Context, cmd PostCreateCommand) (*PostCreateRes, *i18np.Error) {
		e := factory.New(post.NewConfig{
			Owner:         cmd.Owner,
			Images:        cmd.Images,
			Meta:          cmd.Meta,
			CategoryUUIDs: cmd.CategoryUUIDs,
			Features:      cmd.Features,
			Prices:        cmd.Prices,
			Location:      cmd.Location,
			Boosts:        cmd.Boosts,
			People:        cmd.People,
			Type:          cmd.Type,
			Count:         cmd.Count,
			ForCreate:     true,
		})
		err := factory.Validate(*e)
		if err != nil {
			return nil, err
		}
		saved, _err := repo.Create(ctx, e)
		if _err != nil {
			return nil, _err
		}
		events.Created(post.CreatedEvent{
			Entity: saved,
			Account: post.AccountEvent{
				UUID: cmd.Account.UUID,
				Name: cmd.Account.Name,
			},
		})
		return &PostCreateRes{
			UUID: saved.UUID,
		}, nil
	}
}