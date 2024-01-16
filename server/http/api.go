package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.listing/app/command"
	"github.com/turistikrota/service.listing/app/query"
	"github.com/turistikrota/service.listing/domains/account"
	"github.com/turistikrota/service.listing/domains/listing"
	"github.com/turistikrota/service.listing/pkg/utils"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_business"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
)

func (h srv) ListingCreate(ctx *fiber.Ctx) error {
	cmd := command.ListingCreateCmd{}
	h.parseBody(ctx, &cmd)
	a := current_account.Parse(ctx)
	o := current_business.Parse(ctx)
	cmd.Account = account.Entity{
		UUID: current_user.Parse(ctx).UUID,
		Name: a.Name,
	}
	cmd.Business = listing.Business{
		UUID:     o.UUID,
		NickName: o.NickName,
	}
	res, err := h.app.Commands.ListingCreate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.ListingCreated, res)
}

func (h srv) ListingUpdate(ctx *fiber.Ctx) error {
	detail := command.ListingDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.ListingUpdateCmd{}
	cmd.ListingUUID = detail.ListingUUID
	h.parseBody(ctx, &cmd)
	a := current_account.Parse(ctx)
	o := current_business.Parse(ctx)
	cmd.Account = account.Entity{
		UUID: current_user.Parse(ctx).UUID,
		Name: a.Name,
	}
	cmd.Business = listing.Business{
		UUID:     o.UUID,
		NickName: o.NickName,
	}
	res, err := h.app.Commands.ListingUpdate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.ListingUpdated, res)
}

func (h srv) ListingEnable(ctx *fiber.Ctx) error {
	detail := command.ListingDetailCmd{}
	a := current_account.Parse(ctx)
	h.parseParams(ctx, &detail)
	cmd := command.ListingEnableCmd{}
	cmd.ListingUUID = detail.ListingUUID
	cmd.Account = account.Entity{
		UUID: current_user.Parse(ctx).UUID,
		Name: a.Name,
	}
	cmd.BusinessNickName = current_business.Parse(ctx).NickName
	res, err := h.app.Commands.ListingEnable(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ListingDisable(ctx *fiber.Ctx) error {
	detail := command.ListingDetailCmd{}
	a := current_account.Parse(ctx)
	h.parseParams(ctx, &detail)
	cmd := command.ListingDisableCmd{}
	cmd.ListingUUID = detail.ListingUUID
	cmd.Account = account.Entity{
		UUID: current_user.Parse(ctx).UUID,
		Name: a.Name,
	}
	cmd.BusinessNickName = current_business.Parse(ctx).NickName
	res, err := h.app.Commands.ListingDisable(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ListingDelete(ctx *fiber.Ctx) error {
	detail := command.ListingDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.ListingDeleteCmd{}
	cmd.ListingUUID = detail.ListingUUID
	res, err := h.app.Commands.ListingDelete(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ListingReOrder(ctx *fiber.Ctx) error {
	detail := command.ListingDetailCmd{}
	a := current_account.Parse(ctx)
	h.parseParams(ctx, &detail)
	cmd := command.ListingReOrderCmd{}
	cmd.ListingUUID = detail.ListingUUID
	cmd.Account = account.Entity{
		UUID: current_user.Parse(ctx).UUID,
		Name: a.Name,
	}
	cmd.BusinessNickName = current_business.Parse(ctx).NickName
	h.parseBody(ctx, &cmd)
	res, err := h.app.Commands.ListingReOrder(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ListingRestore(ctx *fiber.Ctx) error {
	detail := command.ListingDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.ListingRestoreCmd{}
	cmd.ListingUUID = detail.ListingUUID
	res, err := h.app.Commands.ListingRestore(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ListingView(ctx *fiber.Ctx) error {
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	query := query.ListingViewQuery{}
	query.Locale = l
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.ListingView(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ListingViewAdmin(ctx *fiber.Ctx) error {
	query := query.ListingAdminViewQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.ListingAdminView(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ListingViewBusiness(ctx *fiber.Ctx) error {
	query := query.ListingBusinessViewQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.ListingBusinessView(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ListingFilterByBusiness(ctx *fiber.Ctx) error {
	pagination := utils.Pagination{}
	h.parseQuery(ctx, &pagination)
	filter := listing.FilterEntity{}
	h.parseBody(ctx, &filter)
	query := query.ListingFilterByBusinessQuery{}
	query.Pagination = &pagination
	query.FilterEntity = filter
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.ListingFilterByBusiness(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ListingFilter(ctx *fiber.Ctx) error {
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	pagination := utils.Pagination{}
	h.parseQuery(ctx, &pagination)
	filter := listing.FilterEntity{}
	filter.Locale = l
	h.parseBody(ctx, &filter)
	query := query.ListingFilterQuery{}
	query.Pagination = &pagination
	query.FilterEntity = filter
	res, err := h.app.Queries.ListingFilter(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ListingAdminFilter(ctx *fiber.Ctx) error {
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	pagination := utils.Pagination{}
	h.parseQuery(ctx, &pagination)
	filter := listing.FilterEntity{}
	filter.Locale = l
	h.parseBody(ctx, &filter)
	query := query.ListingAdminFilterQuery{}
	query.Pagination = &pagination
	query.FilterEntity = filter
	res, err := h.app.Queries.ListingAdminFilter(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ListingListMy(ctx *fiber.Ctx) error {
	pagination := utils.Pagination{}
	h.parseQuery(ctx, &pagination)
	business := current_business.Parse(ctx)
	query := query.ListingListMyQuery{}
	query.Pagination = &pagination
	query.BusinessUUID = business.UUID
	res, err := h.app.Queries.ListingListMy(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}
