package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/account"
	"github.com/turistikrota/service.post/domains/post"
)

type PostUpdateCmd struct {
	Account       account.Entity                `json:"-"`
	Owner         post.Owner                    `json:"-"`
	PostUUID      string                        `json:"-"`
	Images        []post.Image                  `json:"images" validate:"min=1,max=10,dive,required"`
	Meta          *PostMetaRequest              `json:"meta" validate:"required,dive"`
	CategoryUUIDs []string                      `json:"categoryUUIDs" validate:"required,min=1,max=30,dive,required,object_id"`
	Features      []post.Feature                `json:"features" validate:"required,min=0,max=30,dive,required"`
	Prices        []post.PostPriceValidationDto `json:"prices" validate:"required,min=1,max=100,dive,required"`
	Location      *post.Location                `json:"location" validate:"required,dive"`
	Boosts        []post.Boost                  `json:"boosts" validate:"omitempty,min=0,max=10,dive,required"`
	Validation    *post.Validation              `json:"validation" validate:"required,dive"`
}

type PostUpdateRes struct {
}

type PostUpdateHandler cqrs.HandlerFunc[PostUpdateCmd, *PostUpdateRes]

func NewPostUpdateHandler(factory post.Factory, repo post.Repository, events post.Events) PostUpdateHandler {
	return func(ctx context.Context, cmd PostUpdateCmd) (*PostUpdateRes, *i18np.Error) {
		e := factory.New(post.NewConfig{
			Owner:         cmd.Owner,
			Images:        cmd.Images,
			Meta:          factory.CreateSlugs(cmd.Meta.TR, cmd.Meta.EN),
			CategoryUUIDs: cmd.CategoryUUIDs,
			Features:      cmd.Features,
			Prices:        cmd.Prices,
			Location:      *cmd.Location,
			Boosts:        cmd.Boosts,
			Validation:    cmd.Validation,
			ForCreate:     false,
		})
		err := factory.Validate(*e)
		if err != nil {
			return nil, err
		}
		old, _err := repo.AdminView(ctx, cmd.PostUUID)
		if _err != nil {
			return nil, _err
		}
		e.CreatedAt = old.CreatedAt
		e.UUID = old.UUID
		e.Order = old.Order
		events.Updated(post.UpdatedEvent{
			Old:    old,
			Entity: e,
			Account: post.AccountEvent{
				UUID: cmd.Account.UUID,
				Name: cmd.Account.Name,
			},
		})
		return &PostUpdateRes{}, nil
	}
}
