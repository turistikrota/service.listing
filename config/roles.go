package config

import "github.com/turistikrota/service.shared/base_roles"

type listingRoles struct {
	Super   string
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

type businessRoles struct {
	Super string
}

type roles struct {
	base_roles.Roles
	Listing  listingRoles
	Business businessRoles
}

var Roles = roles{
	Roles: base_roles.BaseRoles,
	Listing: listingRoles{
		Super:   "listing.super",
		Create:  "listing.create",
		Update:  "listing.update",
		Delete:  "listing.delete",
		Enable:  "listing.enable",
		Disable: "listing.disable",
		ReOrder: "listing.re_order",
		Restore: "listing.restore",
		List:    "listing.list",
		View:    "listing.view",
	},
	Business: businessRoles{
		Super: "business.super",
	},
}
