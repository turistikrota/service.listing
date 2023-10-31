package command

type PostDetailCmd struct {
	PostUUID string `json:"postUUID" params:"uuid" validate:"required,object_id"`
}
