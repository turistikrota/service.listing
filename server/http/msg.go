package http

type successMessages struct {
	PostCreated string
	PostUpdated string
	Ok          string
}

type errorMessages struct {
	RequiredAuth            string
	CurrentUserAccess       string
	AdminRoute              string
	RequiredBusinessSelect  string
	ForbiddenBusinessSelect string
	RequiredAccountSelect   string
	ForbiddenAccountSelect  string
}

type messages struct {
	Success successMessages
	Error   errorMessages
}

var Messages = messages{
	Success: successMessages{
		PostCreated: "http_success_post_created",
		PostUpdated: "http_success_post_updated",
		Ok:          "http_success_ok",
	},
	Error: errorMessages{
		RequiredAuth:            "http_error_required_auth",
		CurrentUserAccess:       "http_error_current_user_access",
		AdminRoute:              "http_error_admin_route",
		RequiredBusinessSelect:  "http_error_required_business_select",
		ForbiddenBusinessSelect: "http_error_forbidden_business_select",
		RequiredAccountSelect:   "http_error_required_account_select",
		ForbiddenAccountSelect:  "http_error_forbidden_account_select",
	},
}
