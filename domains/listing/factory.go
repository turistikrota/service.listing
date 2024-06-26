package listing

import (
	"time"

	"github.com/9ssi7/nanoid"
	"github.com/cilloparch/cillop/i18np"
	"github.com/google/uuid"
	"github.com/ssibrahimbas/slug"
	"github.com/turistikrota/service.listing/domains/payment"
)

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newListingErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

type NewConfig struct {
	Business             Business
	Images               []Image
	Meta                 map[Locale]Meta
	CategoryUUIDs        []string
	Features             []Feature
	Prices               []ListingPriceValidationDto
	ExtraPaymentChannels []payment.Channel
	Location             Location
	Boosts               []Boost
	Validation           *Validation
	ForCreate            bool
	Currency             Currency
}

func (f Factory) New(cnf NewConfig) *Entity {
	t := time.Now()
	prices := make([]Price, len(cnf.Prices))
	for i, p := range cnf.Prices {
		prices[i] = p.ToEntity()
	}
	e := &Entity{
		Business:             cnf.Business,
		Images:               cnf.Images,
		Meta:                 cnf.Meta,
		CategoryUUIDs:        cnf.CategoryUUIDs,
		Features:             cnf.Features,
		Prices:               prices,
		Validation:           cnf.Validation,
		Location:             cnf.Location,
		ExtraPaymentChannels: cnf.ExtraPaymentChannels,
		Boosts:               cnf.Boosts,
		Currency:             cnf.Currency,
		IsActive:             false,
		IsDeleted:            false,
		IsValid:              false,
		UpdatedAt:            t,
	}
	if cnf.ForCreate {
		order := 0
		e.CreatedAt = t
		e.Order = &order
	}
	return e
}

type validator func(e Entity) *i18np.Error

func (f Factory) Validate(entity Entity) *i18np.Error {
	validators := []validator{
		f.validateBusiness,
		f.validatePrices,
		f.validatePeople,
		f.validateMeta,
	}
	for _, v := range validators {
		if err := v(entity); err != nil {
			return err
		}
	}
	return nil
}

func (f Factory) validateBusiness(e Entity) *i18np.Error {
	if e.Business == (Business{}) {
		return f.Errors.InvalidBusiness()
	}
	if e.Business.UUID == "" {
		return f.Errors.InvalidBusiness()
	}
	if e.Business.NickName == "" {
		return f.Errors.InvalidBusiness()
	}
	return nil
}

func (f Factory) validatePeople(e Entity) *i18np.Error {
	if *e.Validation.MinAdult == 0 && *e.Validation.MaxAdult == 0 && *e.Validation.MinKid == 0 && *e.Validation.MaxKid == 0 && *e.Validation.MinBaby == 0 && *e.Validation.MaxBaby == 0 {
		return f.Errors.InvalidPeople()
	}
	if *e.Validation.MinAdult == 0 {
		return f.Errors.MinAdult()
	}
	return nil
}

func (f Factory) validatePrices(e Entity) *i18np.Error {
	for i, p := range e.Prices {
		if p.StartDate.After(p.EndDate) {
			return f.Errors.InvalidPriceDate()
		}
		for j, p2 := range e.Prices {
			if i == j {
				continue
			}
			if p.StartDate.Before(p2.EndDate) && p.EndDate.After(p2.StartDate) {
				return f.Errors.PriceDateConflict(p2.StartDate, p2.EndDate)
			}
		}
	}
	return nil
}

func (f Factory) validateMeta(e Entity) *i18np.Error {
	if _, ok := e.Meta[LocaleEN]; !ok {
		return f.Errors.InvalidMeta()
	}
	if _, ok := e.Meta[LocaleTR]; !ok {
		return f.Errors.InvalidMeta()
	}
	return nil
}

func (f Factory) CreateSlugs(trMeta *Meta, enMeta *Meta, olds ...Meta) map[Locale]Meta {
	if len(olds) > 0 {
		if trMeta.Title == olds[0].Title && enMeta.Title == olds[1].Title {
			return map[Locale]Meta{
				LocaleTR: *trMeta,
				LocaleEN: *enMeta,
			}
		}
	}
	uniqueSuffix, err := nanoid.New()
	if err != nil {
		uniqueSuffix = uuid.New().String()
	}
	trMeta.Slug = slug.New(trMeta.Title+"-"+uniqueSuffix, slug.TR)
	enMeta.Slug = slug.New(enMeta.Title+"-"+uniqueSuffix, slug.EN)
	return map[Locale]Meta{
		LocaleTR: *trMeta,
		LocaleEN: *enMeta,
	}
}
