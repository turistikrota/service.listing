package post

import (
	"context"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
)

type I18nDetail struct {
	Locale string
	Slug   string
}

type Repository interface {
	Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error)
	Update(ctx context.Context, entity *Entity) *i18np.Error
	Delete(ctx context.Context, postUUID string) *i18np.Error
	Restore(ctx context.Context, postUUID string) *i18np.Error
	Disable(ctx context.Context, postUUID string) *i18np.Error
	Enable(ctx context.Context, postUUID string) *i18np.Error
	MarkValid(ctx context.Context, postUUID string) *i18np.Error
	MarkInvalid(ctx context.Context, postUUID string) *i18np.Error
	ReOrder(ctx context.Context, postUUID string, order int) *i18np.Error
	View(ctx context.Context, detail I18nDetail) (*Entity, *i18np.Error)
	AdminView(ctx context.Context, postUUID string) (*Entity, *i18np.Error)
	FilterByOwner(ctx context.Context, ownerNickName string, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)
	ListMy(ctx context.Context, ownerUUID string, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)
}
