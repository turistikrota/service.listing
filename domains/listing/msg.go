package listing

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
	InvalidType:                 "listing_invalid_type",
	InvalidBusiness:             "listing_invalid_business",
	InvalidMeta:                 "listing_invalid_meta",
	MetaMinLength:               "listing_meta_min_length",
	InvalidImages:               "listing_invalid_images",
	ImagesMinLength:             "listing_images_min_length",
	InvalidCategories:           "listing_invalid_categories",
	CategoriesMinLength:         "listing_categories_min_length",
	InvalidPriceDate:            "listing_invalid_price_date",
	PriceDateConflict:           "listing_price_date_conflict",
	Failed:                      "listing_failed",
	InvalidUUID:                 "listing_invalid_uuid",
	NotFound:                    "listing_not_found",
	InvalidPeople:               "listing_invalid_people",
	MinAdult:                    "listing_min_adult",
	ValidateBookingNotAvailable: "listing_validate_booking_not_available",
	ValidateBookingAdult:        "listing_validate_booking_adult",
	ValidateBookingKid:          "listing_validate_booking_kid",
	ValidateBookingBaby:         "listing_validate_booking_baby",
	ValidateBookingNotFound:     "listing_validate_booking_not_found",
}
