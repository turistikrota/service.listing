package post

import (
	"time"

	"github.com/cilloparch/cillop/i18np"
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
	People        People
	Type          Type
	Count         *int
	Order         *int
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
		People:        cnf.People,
		Location:      cnf.Location,
		Boosts:        cnf.Boosts,
		Type:          cnf.Type,
		Order:         cnf.Order,
		IsActive:      false,
		IsDeleted:     false,
		IsValid:       false,
		UpdatedAt:     t,
		Count:         cnf.Count,
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
		f.validateType,
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
	if *e.People.MinAdult == 0 && *e.People.MaxAdult == 0 && *e.People.MinKid == 0 && *e.People.MaxKid == 0 && *e.People.MinBaby == 0 && *e.People.MaxBaby == 0 {
		return f.Errors.InvalidPeople()
	}
	if *e.People.MinAdult == 0 {
		return f.Errors.MinAdult()
	}
	return nil
}

func (f Factory) validateType(e Entity) *i18np.Error {
	list := []Type{
		TypeEstate,
		TypeCar,
		TypeBoat,
		TypeMotorcycle,
		TypeOther,
	}
	for _, v := range list {
		if v == e.Type {
			return nil
		}
	}
	return f.Errors.InvalidType()
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
