package post

import (
	"time"

	"github.com/cilloparch/cillop/i18np"
	"github.com/ssibrahimbas/slug"
)

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newPostErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

type NewConfig struct {
	Owner         Owner
	Images        []Image
	Meta          map[Locale]Meta
	CategoryUUIDs []string
	Features      []Feature
	Prices        []PostPriceValidationDto
	Location      Location
	Boosts        []Boost
	Validation    *Validation
	ForCreate     bool
}

func (f Factory) New(cnf NewConfig) *Entity {
	t := time.Now()
	prices := make([]Price, len(cnf.Prices))
	for i, p := range cnf.Prices {
		prices[i] = p.ToEntity()
	}
	e := &Entity{
		Owner:         cnf.Owner,
		Images:        cnf.Images,
		Meta:          cnf.Meta,
		CategoryUUIDs: cnf.CategoryUUIDs,
		Features:      cnf.Features,
		Prices:        prices,
		Validation:    cnf.Validation,
		Location:      cnf.Location,
		Boosts:        cnf.Boosts,
		IsActive:      false,
		IsDeleted:     false,
		IsValid:       false,
		UpdatedAt:     t,
	}
	if cnf.ForCreate {
		e.CreatedAt = t
	}
	return e
}

type validator func(e Entity) *i18np.Error

func (f Factory) Validate(entity Entity) *i18np.Error {
	validators := []validator{
		f.validateOwner,
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

func (f Factory) validateOwner(e Entity) *i18np.Error {
	if e.Owner == (Owner{}) {
		return f.Errors.InvalidOwner()
	}
	if e.Owner.UUID == "" {
		return f.Errors.InvalidOwner()
	}
	if e.Owner.NickName == "" {
		return f.Errors.InvalidOwner()
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

func (f Factory) CreateSlugs(trMeta *Meta, enMeta *Meta) map[Locale]Meta {
	trMeta.Slug = slug.New(trMeta.Title, slug.TR)
	enMeta.Slug = slug.New(enMeta.Title, slug.EN)
	return map[Locale]Meta{
		LocaleTR: *trMeta,
		LocaleEN: *enMeta,
	}
}
