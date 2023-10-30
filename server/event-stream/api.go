package event_stream

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/turistikrota/service.post/app/command"
	"github.com/turistikrota/service.post/domains/account"
	"github.com/turistikrota/service.post/domains/category"
)

func (s srv) OnPostValidationSuccess(data []byte) {
	fmt.Println("OnPostValidationSuccess")
	e := category.ValidationSuccessEvent{}
	err := json.Unmarshal(data, &e)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, _ = s.app.Commands.PostValidated(context.Background(), command.PostValidatedCmd{
		New: e.Post,
		Account: account.Entity{
			UUID: e.User.UUID,
			Name: e.User.Name,
		},
	})
}
