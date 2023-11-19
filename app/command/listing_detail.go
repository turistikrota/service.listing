package command

type ListingDetailCmd struct {
	ListingUUID string `json:"listingUUID" params:"uuid" validate:"required,object_id"`
}
