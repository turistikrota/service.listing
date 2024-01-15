package listing

import (
	"time"

	"github.com/cilloparch/cillop/formats"
)

type ListingPriceValidationDto struct {
	Price     float64 `json:"price" validate:"required,gt=0"`
	StartDate string  `json:"startDate" validate:"required,datetime=2006-01-02"`
	EndDate   string  `json:"endDate" validate:"required,datetime=2006-01-02"`
}

type ListDto struct {
	UUID     string          `json:"uuid" bson:"_id,omitempty"`
	Business Business        `json:"business" bson:"business"`
	Images   []Image         `json:"images" bson:"images"`
	Meta     map[Locale]Meta `json:"meta" bson:"meta"`
	Prices   []Price         `json:"prices" bson:"prices"`
	Currency Currency        `json:"currency" bson:"currency"`
	Location Location        `json:"location" bson:"location"`
}

type DetailDto struct {
	UUID          string          `json:"uuid" bson:"_id,omitempty"`
	Business      Business        `json:"business" bson:"business"`
	Images        []Image         `json:"images" bson:"images"`
	Meta          map[Locale]Meta `json:"meta" bson:"meta"`
	CategoryUUIDs []string        `json:"categoryUUIDs" bson:"categoryUUIDs"`
	Features      []Feature       `json:"features" bson:"features"`
	Prices        []Price         `json:"prices" bson:"prices"`
	Location      Location        `json:"location" bson:"location"`
	Boosts        []Boost         `json:"boosts" bson:"boosts"`
	Currency      Currency        `json:"currency" bson:"currency"`
	Validation    Validation      `json:"validation" bson:"validation"`
	CreatedAt     time.Time       `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time       `json:"updatedAt" bson:"updated_at"`
}

type AdminDetailDto struct {
	*Entity
}

type AdminListDto struct {
	UUID          string          `json:"uuid" bson:"_id,omitempty"`
	Business      Business        `json:"business" bson:"business"`
	Images        []Image         `json:"images" bson:"images"`
	Meta          map[Locale]Meta `json:"meta" bson:"meta"`
	CategoryUUIDs []string        `json:"categoryUUIDs" bson:"categoryUUIDs"`
	Features      []Feature       `json:"features" bson:"features"`
	Prices        []Price         `json:"prices" bson:"prices"`
	Location      Location        `json:"location" bson:"location"`
	Boosts        []Boost         `json:"boosts" bson:"boosts"`
	Validation    Validation      `json:"validation" bson:"validation"`
	Currency      Currency        `json:"currency" bson:"currency"`
	Order         *int            `json:"order" bson:"order"`
	IsActive      bool            `json:"isActive" bson:"is_active"`
	IsDeleted     bool            `json:"isDeleted" bson:"is_deleted"`
	IsValid       bool            `json:"isValid" bson:"is_valid"`
	CreatedAt     time.Time       `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time       `json:"updatedAt" bson:"updated_at"`
}

type BusinessListDto struct {
	UUID      string          `json:"uuid" bson:"_id,omitempty"`
	Images    []Image         `json:"images" bson:"images"`
	Meta      map[Locale]Meta `json:"meta" bson:"meta"`
	Location  Location        `json:"location" bson:"location"`
	Boosts    []Boost         `json:"boosts" bson:"boosts"`
	Currency  Currency        `json:"currency" bson:"currency"`
	Order     *int            `json:"order" bson:"order"`
	IsActive  bool            `json:"isActive" bson:"is_active"`
	IsDeleted bool            `json:"isDeleted" bson:"is_deleted"`
	IsValid   bool            `json:"isValid" bson:"is_valid"`
	CreatedAt time.Time       `json:"createdAt" bson:"created_at"`
}

type PricePerDay struct {
	Date  time.Time `json:"date"`
	Price float64   `json:"price"`
}

func (d *ListingPriceValidationDto) ToEntity() Price {
	start, _ := time.Parse(formats.DateYYYYMMDD, d.StartDate)
	end, _ := time.Parse(formats.DateYYYYMMDD, d.EndDate)
	return Price{
		Price:     d.Price,
		StartDate: start,
		EndDate:   end,
	}
}

func (e *Entity) ToList() *ListDto {
	return &ListDto{
		UUID:     e.UUID,
		Business: e.Business,
		Images:   e.Images,
		Meta:     e.Meta,
		Prices:   e.Prices,
		Location: e.Location,
		Currency: e.Currency,
	}
}

func (e *Entity) ToDetail() *DetailDto {
	return &DetailDto{
		UUID:          e.UUID,
		Business:      e.Business,
		Images:        e.Images,
		Meta:          e.Meta,
		CategoryUUIDs: e.CategoryUUIDs,
		Features:      e.Features,
		Prices:        e.Prices,
		Location:      e.Location,
		Boosts:        e.Boosts,
		Currency:      e.Currency,
		Validation:    *e.Validation,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
	}
}

func (e *Entity) ToAdminDetail() *AdminDetailDto {
	return &AdminDetailDto{
		Entity: e,
	}
}

func (e *Entity) ToAdminList() *AdminListDto {
	return &AdminListDto{
		UUID:          e.UUID,
		Business:      e.Business,
		Images:        e.Images,
		Meta:          e.Meta,
		CategoryUUIDs: e.CategoryUUIDs,
		Features:      e.Features,
		Prices:        e.Prices,
		Location:      e.Location,
		Boosts:        e.Boosts,
		Currency:      e.Currency,
		Validation:    *e.Validation,
		Order:         e.Order,
		IsActive:      e.IsActive,
		IsDeleted:     e.IsDeleted,
		IsValid:       e.IsValid,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
	}
}

func (e *Entity) ToBusinessList() *BusinessListDto {
	return &BusinessListDto{
		UUID:      e.UUID,
		Images:    e.Images,
		Meta:      e.Meta,
		Currency:  e.Currency,
		Location:  e.Location,
		Boosts:    e.Boosts,
		Order:     e.Order,
		IsActive:  e.IsActive,
		IsDeleted: e.IsDeleted,
		IsValid:   e.IsValid,
		CreatedAt: e.CreatedAt,
	}
}
