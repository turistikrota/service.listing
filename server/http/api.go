package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.post/app/command"
	"github.com/turistikrota/service.post/app/query"
	"github.com/turistikrota/service.post/domains/account"
	"github.com/turistikrota/service.post/domains/post"
	"github.com/turistikrota/service.post/pkg/utils"
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
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
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
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
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
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
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
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
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
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
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
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) PostRestore(ctx *fiber.Ctx) error {
	detail := command.PostDetailCmd{}
	a := current_account.Parse(ctx)
	h.parseParams(ctx, &detail)
	cmd := command.PostRestoreCmd{}
	cmd.PostUUID = detail.PostUUID
	cmd.Account = account.Entity{
		UUID: a.ID,
		Name: a.Name,
	}
	res, err := h.app.Commands.PostRestore(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) PostView(ctx *fiber.Ctx) error {
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	query := query.PostViewQuery{}
	query.Locale = l
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.PostView(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) PostViewAdmin(ctx *fiber.Ctx) error {
	query := query.PostAdminViewQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.PostAdminView(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) PostFilterByOwner(ctx *fiber.Ctx) error {
	pagination := utils.Pagination{}
	h.parseQuery(ctx, &pagination)
	filter := post.FilterEntity{}
	h.parseBody(ctx, &filter)
	query := query.PostFilterByOwnerQuery{}
	query.Pagination = &pagination
	query.FilterEntity = filter
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.PostFilterByOwner(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) PostListMy(ctx *fiber.Ctx) error {
	pagination := utils.Pagination{}
	h.parseQuery(ctx, &pagination)
	owner := current_owner.Parse(ctx)
	query := query.PostListMyQuery{}
	query.Pagination = &pagination
	query.OwnerUUID = owner.UUID
	res, err := h.app.Queries.PostListMy(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}
