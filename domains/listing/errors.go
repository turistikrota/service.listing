package listing

import (
	"time"

	"github.com/cilloparch/cillop/formats"
	"github.com/cilloparch/cillop/i18np"
)

type Errors interface {
	InvalidType() *i18np.Error
	InvalidBusiness() *i18np.Error
	InvalidMeta() *i18np.Error
	InvalidImages() *i18np.Error
	InvalidCategories() *i18np.Error
	MetaMinLength() *i18np.Error
	ImagesMinLength() *i18np.Error
	CategoriesMinLength() *i18np.Error
	InvalidPriceDate() *i18np.Error
	PriceDateConflict(time.Time, time.Time) *i18np.Error
	Failed(string) *i18np.Error
	InvalidUUID() *i18np.Error
	InvalidPeople() *i18np.Error
	NotFound() *i18np.Error
	MinAdult() *i18np.Error
	ValidateBookingNotAvailable(time.Time) *i18np.Error
	ValidateBookingAdult(int, int) *i18np.Error
	ValidateBookingKid(int, int) *i18np.Error
	ValidateBookingBaby(int, int) *i18np.Error
	ValidateBookingNotFound() *i18np.Error
}

type listingErrors struct{}

func newListingErrors() Errors {
	return &listingErrors{}
}

func (e *listingErrors) InvalidType() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidType)
}

func (e *listingErrors) InvalidBusiness() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidBusiness)
}

func (e *listingErrors) InvalidMeta() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidMeta)
}

func (e *listingErrors) MetaMinLength() *i18np.Error {
	return i18np.NewError(i18nMessages.MetaMinLength)
}

func (e *listingErrors) InvalidImages() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidImages)
}

func (e *listingErrors) ImagesMinLength() *i18np.Error {
	return i18np.NewError(i18nMessages.ImagesMinLength)
}

func (e *listingErrors) InvalidCategories() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidCategories)
}

func (e *listingErrors) CategoriesMinLength() *i18np.Error {
	return i18np.NewError(i18nMessages.CategoriesMinLength)
}

func (e *listingErrors) InvalidPriceDate() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidPriceDate)
}

func (e *listingErrors) PriceDateConflict(start time.Time, end time.Time) *i18np.Error {
	return i18np.NewError(i18nMessages.PriceDateConflict, i18np.P{
		"Start": start.Format(formats.DateYYYYMMDD),
		"End":   end.Format(formats.DateYYYYMMDD),
	})
}

func (e *listingErrors) Failed(operation string) *i18np.Error {
	return i18np.NewError(i18nMessages.Failed, i18np.P{
		"Operation": operation,
	})
}

func (e *listingErrors) InvalidUUID() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidUUID)
}

func (e *listingErrors) NotFound() *i18np.Error {
	return i18np.NewError(i18nMessages.NotFound)
}

func (e *listingErrors) InvalidPeople() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidPeople)
}

func (e *listingErrors) MinAdult() *i18np.Error {
	return i18np.NewError(i18nMessages.MinAdult, i18np.P{
		"Min": 1,
	})
}

func (e *listingErrors) ValidateBookingNotAvailable(date time.Time) *i18np.Error {
	return i18np.NewError(i18nMessages.ValidateBookingNotAvailable, i18np.P{
		"Date": date.Format(formats.DateYYYYMMDD),
	})
}

func (e *listingErrors) ValidateBookingAdult(min int, max int) *i18np.Error {
	return i18np.NewError(i18nMessages.ValidateBookingAdult, i18np.P{
		"Min": min,
		"Max": max,
	})
}

func (e *listingErrors) ValidateBookingKid(min int, max int) *i18np.Error {
	return i18np.NewError(i18nMessages.ValidateBookingKid, i18np.P{
		"Min": min,
		"Max": max,
	})
}

func (e *listingErrors) ValidateBookingBaby(min int, max int) *i18np.Error {
	return i18np.NewError(i18nMessages.ValidateBookingBaby, i18np.P{
		"Min": min,
		"Max": max,
	})
}

func (e *listingErrors) ValidateBookingNotFound() *i18np.Error {
	return i18np.NewError(i18nMessages.ValidateBookingNotFound)
}
