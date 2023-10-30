package app

import (
	"github.com/turistikrota/service.post/app/command"
	"github.com/turistikrota/service.post/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	PostCreate              command.PostCreateHandler
	PostUpdate              command.PostUpdateHandler
	PostValidated           command.PostValidatedHandler
	PostUpdateOwnerNickName command.PostUpdateOwnerNickNameHandler
	PostEnable              command.PostEnableHandler
	PostDisable             command.PostDisableHandler
	PostDelete              command.PostDeleteHandler
	PostRestore             command.PostRestoreHandler
	PostReOrder             command.PostReOrderHandler
}

type Queries struct {
	PostView          query.PostViewHandler
	PostAdminView     query.PostAdminViewHandler
	PostFilterByOwner query.PostFilterByOwnerHandler
	PostListMy        query.PostListMyHandler
}
