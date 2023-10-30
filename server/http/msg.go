package http

type successMessages struct {
	PostCreated string
	PostUpdated string
}

type errorMessages struct {
	RequiredAuth      string
	CurrentUserAccess string
	AdminRoute        string
}

type messages struct {
	Success successMessages
	Error   errorMessages
}

var Messages = messages{
	Success: successMessages{
		PostCreated: "http_success_post_created",
		PostUpdated: "http_success_post_updated",
	},
	Error: errorMessages{
		RequiredAuth:      "http_error_required_auth",
		CurrentUserAccess: "http_error_current_user_access",
		AdminRoute:        "http_error_admin_route",
	},
}
