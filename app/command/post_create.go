package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/account"
	"github.com/turistikrota/service.post/domains/post"
)

type PostMetaRequest struct {
	TR *post.Meta `json:"tr" validate:"required,dive"`
	EN *post.Meta `json:"en" validate:"required,dive"`
}

type PostCreateCmd struct {
	Account       account.Entity                `json:"-"`
	Owner         post.Owner                    `json:"-"`
	Images        []post.Image                  `json:"images" validate:"min=1,max=10,dive,required"`
	Meta          *PostMetaRequest              `json:"meta" validate:"required,dive"`
	CategoryUUIDs []string                      `json:"categoryUUIDs" validate:"required,min=1,max=30,dive,required,object_id"`
	Features      []post.Feature                `json:"features" validate:"required,min=0,max=30,dive,required"`
	Prices        []post.PostPriceValidationDto `json:"prices" validate:"required,min=1,max=100,dive,required"`
	Location      *post.Location                `json:"location" validate:"required,dive"`
	Boosts        []post.Boost                  `json:"boosts" validate:"omitempty,min=0,max=10,dive,required"`
	Type          post.Type                     `json:"type" validate:"required"`
	People        *post.People                  `json:"people" validate:"required,dive"`
	Count         *int                          `json:"count" validate:"required,min=1,max=100,numeric"`
	Order         *int                          `json:"order" validate:"required,min=0,max=1000,numeric"`
}

type PostCreateRes struct {
	UUID string `json:"uuid"`
}

type PostCreateHandler cqrs.HandlerFunc[PostCreateCmd, *PostCreateRes]

func NewPostCreateHandler(factory post.Factory, repo post.Repository, events post.Events) PostCreateHandler {
	return func(ctx context.Context, cmd PostCreateCmd) (*PostCreateRes, *i18np.Error) {
		meta := map[post.Locale]post.Meta{
			post.LocaleTR: *cmd.Meta.TR,
			post.LocaleEN: *cmd.Meta.EN,
		}
		e := factory.New(post.NewConfig{
			Owner:         cmd.Owner,
			Images:        cmd.Images,
			Meta:          meta,
			CategoryUUIDs: cmd.CategoryUUIDs,
			Features:      cmd.Features,
			Prices:        cmd.Prices,
			Location:      *cmd.Location,
			Boosts:        cmd.Boosts,
			People:        *cmd.People,
			Type:          cmd.Type,
			Count:         cmd.Count,
			Order:         cmd.Order,
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
