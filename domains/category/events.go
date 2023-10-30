package category

import (
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.post/domains/post"
)

type ValidationSuccessEvent struct {
	PostUUID string          `json:"postUUID"`
	Post     *post.Entity    `json:"entity"`
	User     UserDetailEvent `json:"user"`
}

type ValidationFailedEvent struct {
	PostUUID string             `json:"postUUID"`
	Post     *post.Entity       `json:"entity"`
	Errors   []*ValidationError `json:"errors"`
	User     UserDetailEvent    `json:"user"`
}

type UserDetailEvent struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type ValidationError struct {
	Field   string  `json:"field"`
	Message string  `json:"message"`
	Params  i18np.P `json:"params"`
}
