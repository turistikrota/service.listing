package post

import (
	"time"

	"github.com/cilloparch/cillop/formats"
	"github.com/cilloparch/cillop/i18np"
)

type Errors interface {
	InvalidType() *i18np.Error
	InvalidOwner() *i18np.Error
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

type postErrors struct{}

func newPostErrors() Errors {
	return &postErrors{}
}

func (e *postErrors) InvalidType() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidType)
}

func (e *postErrors) InvalidOwner() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidOwner)
}

func (e *postErrors) InvalidMeta() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidMeta)
}

func (e *postErrors) MetaMinLength() *i18np.Error {
	return i18np.NewError(i18nMessages.MetaMinLength)
}

func (e *postErrors) InvalidImages() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidImages)
}

func (e *postErrors) ImagesMinLength() *i18np.Error {
	return i18np.NewError(i18nMessages.ImagesMinLength)
}

func (e *postErrors) InvalidCategories() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidCategories)
}

func (e *postErrors) CategoriesMinLength() *i18np.Error {
	return i18np.NewError(i18nMessages.CategoriesMinLength)
}

func (e *postErrors) InvalidPriceDate() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidPriceDate)
}

func (e *postErrors) PriceDateConflict(start time.Time, end time.Time) *i18np.Error {
	return i18np.NewError(i18nMessages.PriceDateConflict, i18np.P{
		"Start": start.Format(formats.DateYYYYMMDD),
		"End":   end.Format(formats.DateYYYYMMDD),
	})
}

func (e *postErrors) Failed(operation string) *i18np.Error {
	return i18np.NewError(i18nMessages.Failed, i18np.P{
		"Operation": operation,
	})
}

func (e *postErrors) InvalidUUID() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidUUID)
}

func (e *postErrors) NotFound() *i18np.Error {
	return i18np.NewError(i18nMessages.NotFound)
}

func (e *postErrors) InvalidPeople() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidPeople)
}

func (e *postErrors) MinAdult() *i18np.Error {
	return i18np.NewError(i18nMessages.MinAdult, i18np.P{
		"Min": 1,
	})
}

func (e *postErrors) ValidateBookingNotAvailable(date time.Time) *i18np.Error {
	return i18np.NewError(i18nMessages.ValidateBookingNotAvailable, i18np.P{
		"Date": date.Format(formats.DateYYYYMMDD),
	})
}

func (e *postErrors) ValidateBookingAdult(min int, max int) *i18np.Error {
	return i18np.NewError(i18nMessages.ValidateBookingAdult, i18np.P{
		"Min": min,
		"Max": max,
	})
}

func (e *postErrors) ValidateBookingKid(min int, max int) *i18np.Error {
	return i18np.NewError(i18nMessages.ValidateBookingKid, i18np.P{
		"Min": min,
		"Max": max,
	})
}

func (e *postErrors) ValidateBookingBaby(min int, max int) *i18np.Error {
	return i18np.NewError(i18nMessages.ValidateBookingBaby, i18np.P{
		"Min": min,
		"Max": max,
	})
}

func (e *postErrors) ValidateBookingNotFound() *i18np.Error {
	return i18np.NewError(i18nMessages.ValidateBookingNotFound)
}
