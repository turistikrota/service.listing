package command

type PostDetailCmd struct {
	PostUUID string `json:"postUUID" params:"post_uuid" validate:"required,object_id"`
}
