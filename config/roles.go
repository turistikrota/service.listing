package config

import "github.com/turistikrota/service.shared/base_roles"

type postRoles struct {
	Create  string
	Update  string
	Delete  string
	Enable  string
	Disable string
	ReOrder string
	Restore string
	List    string
	View    string
}

type roles struct {
	base_roles.Roles
	Post postRoles
}

var Roles = roles{
	Roles: base_roles.BaseRoles,
	Post: postRoles{
		Create:  "post.create",
		Update:  "post.update",
		Delete:  "post.delete",
		Enable:  "post.enable",
		Disable: "post.disable",
		ReOrder: "post.re_order",
		Restore: "post.restore",
		List:    "post.list",
		View:    "post.view",
	},
}
