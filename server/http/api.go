package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.post/app/command"
	"github.com/turistikrota/service.post/domains/account"
	"github.com/turistikrota/service.post/domains/post"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_owner"
)

func (h srv) PostCreate(ctx *fiber.Ctx) error {
	cmd := command.PostCreateCmd{}
	h.parseBody(ctx, &cmd)
	a := current_account.Parse(ctx)
	o := current_owner.Parse(ctx)
	cmd.Account = account.Entity{
		UUID: a.ID,
		Name: a.Name,
	}
	cmd.Owner = post.Owner{
		UUID:     o.UUID,
		NickName: o.NickName,
	}
	res, err := h.app.Commands.PostCreate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.PostCreated, res)
}

func (h srv) PostUpdate(ctx *fiber.Ctx) error {
	detail := command.PostDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.PostUpdateCmd{}
	cmd.PostUUID = detail.PostUUID
	h.parseBody(ctx, &cmd)
	a := current_account.Parse(ctx)
	o := current_owner.Parse(ctx)
	cmd.Account = account.Entity{
		UUID: a.ID,
		Name: a.Name,
	}
	cmd.Owner = post.Owner{
		UUID:     o.UUID,
		NickName: o.NickName,
	}
	res, err := h.app.Commands.PostUpdate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.PostUpdated, res)
}

func (h srv) PostEnable(ctx *fiber.Ctx) error {
	detail := command.PostDetailCmd{}
	a := current_account.Parse(ctx)
	h.parseParams(ctx, &detail)
	cmd := command.PostEnableCmd{}
	cmd.PostUUID = detail.PostUUID
	cmd.Account = account.Entity{
		UUID: a.ID,
		Name: a.Name,
	}
	res, err := h.app.Commands.PostEnable(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) PostDisable(ctx *fiber.Ctx) error {
	detail := command.PostDetailCmd{}
	a := current_account.Parse(ctx)
	h.parseParams(ctx, &detail)
	cmd := command.PostDisableCmd{}
	cmd.PostUUID = detail.PostUUID
	cmd.Account = account.Entity{
		UUID: a.ID,
		Name: a.Name,
	}
	res, err := h.app.Commands.PostDisable(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) PostDelete(ctx *fiber.Ctx) error {
	detail := command.PostDetailCmd{}
	a := current_account.Parse(ctx)
	h.parseParams(ctx, &detail)
	cmd := command.PostDeleteCmd{}
	cmd.PostUUID = detail.PostUUID
	cmd.Account = account.Entity{
		UUID: a.ID,
		Name: a.Name,
	}
	res, err := h.app.Commands.PostDelete(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) PostReOrder(ctx *fiber.Ctx) error {
	detail := command.PostDetailCmd{}
	a := current_account.Parse(ctx)
	h.parseParams(ctx, &detail)
	cmd := command.PostReOrderCmd{}
	cmd.PostUUID = detail.PostUUID
	cmd.Account = account.Entity{
		UUID: a.ID,
		Name: a.Name,
	}
	h.parseBody(ctx, &cmd)
	res, err := h.app.Commands.PostReOrder(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}
