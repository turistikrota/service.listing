package post

type messages struct {
	InvalidType                 string
	InvalidBusiness             string
	InvalidMeta                 string
	MetaMinLength               string
	InvalidImages               string
	ImagesMinLength             string
	InvalidCategories           string
	CategoriesMinLength         string
	InvalidPriceDate            string
	PriceDateConflict           string
	Failed                      string
	InvalidUUID                 string
	NotFound                    string
	InvalidPeople               string
	MinAdult                    string
	ValidateBookingNotAvailable string
	ValidateBookingAdult        string
	ValidateBookingKid          string
	ValidateBookingBaby         string
	ValidateBookingNotFound     string
}

var i18nMessages = messages{
	InvalidType:                 "post_invalid_type",
	InvalidBusiness:             "post_invalid_business",
	InvalidMeta:                 "post_invalid_meta",
	MetaMinLength:               "post_meta_min_length",
	InvalidImages:               "post_invalid_images",
	ImagesMinLength:             "post_images_min_length",
	InvalidCategories:           "post_invalid_categories",
	CategoriesMinLength:         "post_categories_min_length",
	InvalidPriceDate:            "post_invalid_price_date",
	PriceDateConflict:           "post_price_date_conflict",
	Failed:                      "post_failed",
	InvalidUUID:                 "post_invalid_uuid",
	NotFound:                    "post_not_found",
	InvalidPeople:               "post_invalid_people",
	MinAdult:                    "post_min_adult",
	ValidateBookingNotAvailable: "post_validate_booking_not_available",
	ValidateBookingAdult:        "post_validate_booking_adult",
	ValidateBookingKid:          "post_validate_booking_kid",
	ValidateBookingBaby:         "post_validate_booking_baby",
	ValidateBookingNotFound:     "post_validate_booking_not_found",
}
