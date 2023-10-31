package http

import (
	"strings"
	"time"

	"github.com/cilloparch/cillop/helpers/http"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/server"
	"github.com/cilloparch/cillop/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/timeout"

	"github.com/turistikrota/service.post/app"
	"github.com/turistikrota/service.post/config"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	httpServer "github.com/turistikrota/service.shared/server/http"
	"github.com/turistikrota/service.shared/server/http/auth"
	"github.com/turistikrota/service.shared/server/http/auth/claim_guard"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_owner"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.shared/server/http/auth/device_uuid"
	"github.com/turistikrota/service.shared/server/http/auth/required_access"
)

type srv struct {
	config      config.App
	app         app.Application
	i18n        *i18np.I18n
	validator   validation.Validator
	tknSrv      token.Service
	sessionSrv  session.Service
	httpHeaders config.HttpHeaders
}

type Config struct {
	Env         config.App
	App         app.Application
	I18n        *i18np.I18n
	Validator   validation.Validator
	HttpHeaders config.HttpHeaders
	TokenSrv    token.Service
	SessionSrv  session.Service
}

func New(config Config) server.Server {
	return srv{
		config:      config.Env,
		app:         config.App,
		i18n:        config.I18n,
		validator:   config.Validator,
		tknSrv:      config.TokenSrv,
		sessionSrv:  config.SessionSrv,
		httpHeaders: config.HttpHeaders,
	}
}

func (h srv) Listen() error {
	return http.RunServer(http.Config{
		Host:        h.config.Http.Host,
		Port:        h.config.Http.Port,
		I18n:        h.i18n,
		AcceptLangs: []string{},
		CreateHandler: func(router fiber.Router) fiber.Router {
			router.Use(h.cors(), h.deviceUUID())

			// Owner routes
			owner := router.Group("/owner", h.rateLimit(), h.currentUserAccess(), h.requiredAccess(), h.currentAccountAccess())
			owner.Post("/", h.currentOwnerAccess(config.Roles.Post.Super, config.Roles.Post.Create), h.wrapWithTimeout(h.PostCreate))
			owner.Put("/:uuid", h.currentOwnerAccess(config.Roles.Post.Super, config.Roles.Post.Update), h.wrapWithTimeout(h.PostUpdate))
			owner.Put("/:uuid/enable", h.currentOwnerAccess(config.Roles.Post.Super, config.Roles.Post.Enable), h.wrapWithTimeout(h.PostEnable))
			owner.Put("/:uuid/disable", h.currentOwnerAccess(config.Roles.Post.Super, config.Roles.Post.Disable), h.wrapWithTimeout(h.PostDisable))
			owner.Put("/:uuid/restore", h.currentOwnerAccess(config.Roles.Post.Super, config.Roles.Post.Restore), h.wrapWithTimeout(h.PostRestore))
			owner.Delete("/:uuid", h.currentOwnerAccess(config.Roles.Post.Super, config.Roles.Post.Delete), h.wrapWithTimeout(h.PostDelete))
			owner.Put("/:uuid/re-order", h.currentOwnerAccess(config.Roles.Post.Super, config.Roles.Post.ReOrder), h.wrapWithTimeout(h.PostReOrder))
			owner.Get("/", h.currentOwnerAccess(config.Roles.Post.Super, config.Roles.Post.List), h.wrapWithTimeout(h.PostListMy))
			owner.Get("/:uuid", h.currentOwnerAccess(config.Roles.Post.Super, config.Roles.Post.View), h.wrapWithTimeout(h.PostViewAdmin))

			// Admin routes
			admin := router.Group("/admin", h.currentUserAccess(), h.requiredAccess(), h.adminRoute())
			admin.Get("/:uuid", h.wrapWithTimeout(h.PostViewAdmin))

			// Public routes
			router.Get("/:slug", h.rateLimit(), h.wrapWithTimeout(h.PostView))
			router.Post("/filter/:nickName", h.rateLimit(), h.wrapWithTimeout(h.PostFilterByOwner))
			router.Post("/filter", h.rateLimit(), h.wrapWithTimeout(h.PostFilter))

			return router
		},
	})
}

func (h srv) currentOwnerAccess(roles ...string) fiber.Handler {
	return current_owner.New(current_owner.Config{
		I18n:         h.i18n,
		Roles:        roles,
		RequiredKey:  Messages.Error.RequiredOwnerSelect,
		ForbiddenKey: Messages.Error.ForbiddenOwnerSelect,
	})
}

func (h srv) currentAccountAccess() fiber.Handler {
	return current_account.New(current_account.Config{
		I18n:         h.i18n,
		RequiredKey:  Messages.Error.RequiredAccountSelect,
		ForbiddenKey: Messages.Error.ForbiddenAccountSelect,
	})
}

func (h srv) parseBody(c *fiber.Ctx, d interface{}) {
	http.ParseBody(c, h.validator, *h.i18n, d)
}

func (h srv) parseParams(c *fiber.Ctx, d interface{}) {
	http.ParseParams(c, h.validator, *h.i18n, d)
}

func (h srv) parseQuery(c *fiber.Ctx, d interface{}) {
	http.ParseQuery(c, h.validator, *h.i18n, d)
}

func (h srv) currentUserAccess() fiber.Handler {
	return current_user.New(current_user.Config{
		TokenSrv:   h.tknSrv,
		SessionSrv: h.sessionSrv,
		I18n:       h.i18n,
		MsgKey:     Messages.Error.CurrentUserAccess,
		HeaderKey:  httpServer.Headers.Authorization,
		CookieKey:  auth.Cookies.AccessToken,
		UseCookie:  true,
		UseBearer:  true,
		IsRefresh:  false,
		IsAccess:   true,
	})
}

func (h srv) rateLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        50,
		Expiration: 1 * time.Minute,
	})
}

func (h srv) deviceUUID() fiber.Handler {
	return device_uuid.New(device_uuid.Config{
		Domain: h.httpHeaders.Domain,
	})
}

func (h srv) requiredAccess() fiber.Handler {
	return required_access.New(required_access.Config{
		I18n:   *h.i18n,
		MsgKey: Messages.Error.RequiredAuth,
	})
}

func (h srv) adminRoute(extra ...string) fiber.Handler {
	claims := []string{config.Roles.Admin}
	if len(extra) > 0 {
		claims = append(claims, extra...)
	}
	return claim_guard.New(claim_guard.Config{
		Claims: claims,
		I18n:   *h.i18n,
		MsgKey: Messages.Error.AdminRoute,
	})
}

func (h srv) cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowMethods:     h.httpHeaders.AllowedMethods,
		AllowHeaders:     h.httpHeaders.AllowedHeaders,
		AllowCredentials: h.httpHeaders.AllowCredentials,
		AllowOriginsFunc: func(origin string) bool {
			origins := strings.Split(h.httpHeaders.AllowedOrigins, ",")
			for _, o := range origins {
				if strings.Contains(origin, o) {
					return true
				}
			}
			return false
		},
	})
}

func (h srv) wrapWithTimeout(fn fiber.Handler) fiber.Handler {
	return timeout.NewWithContext(fn, 10*time.Second)
}
