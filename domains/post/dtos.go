package post

import (
	"time"

	"github.com/cilloparch/cillop/formats"
)

type PostPriceValidationDto struct {
	Price     float64 `json:"price" validate:"required,gt=0"`
	StartDate string  `json:"startDate" validate:"required,datetime=2006-01-02"`
	EndDate   string  `json:"endDate" validate:"required,datetime=2006-01-02"`
}

type ListDto struct {
	UUID          string          `json:"uuid" bson:"_id,omitempty"`
	Owner         Owner           `json:"owner" bson:"owner"`
	Images        []Image         `json:"images" bson:"images"`
	Meta          map[Locale]Meta `json:"meta" bson:"meta"`
	CategoryUUIDs []string        `json:"categoryUUIDs" bson:"categoryUUIDs"`
	Features      []Feature       `json:"features" bson:"features"`
	Prices        []Price         `json:"prices" bson:"prices"`
	Location      Location        `json:"location" bson:"location"`
	Boosts        []Boost         `json:"boosts" bson:"boosts"`
	People        People          `json:"people" bson:"people"`
	Type          Type            `json:"type" bson:"type"`
	Count         *int            `json:"count" bson:"count"`
}

type DetailDto struct {
	UUID          string          `json:"uuid" bson:"_id,omitempty"`
	Owner         Owner           `json:"owner" bson:"owner"`
	Images        []Image         `json:"images" bson:"images"`
	Meta          map[Locale]Meta `json:"meta" bson:"meta"`
	CategoryUUIDs []string        `json:"categoryUUIDs" bson:"categoryUUIDs"`
	Features      []Feature       `json:"features" bson:"features"`
	Prices        []Price         `json:"prices" bson:"prices"`
	Location      Location        `json:"location" bson:"location"`
	Boosts        []Boost         `json:"boosts" bson:"boosts"`
	People        People          `json:"people" bson:"people"`
	Type          Type            `json:"type" bson:"type"`
	Count         *int            `json:"count" bson:"count"`
	CreatedAt     time.Time       `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time       `json:"updatedAt" bson:"updated_at"`
}

type AdminDetailDto struct{}

type AdminListDto struct {
	UUID          string          `json:"uuid" bson:"_id,omitempty"`
	Owner         Owner           `json:"owner" bson:"owner"`
	Images        []Image         `json:"images" bson:"images"`
	Meta          map[Locale]Meta `json:"meta" bson:"meta"`
	CategoryUUIDs []string        `json:"categoryUUIDs" bson:"categoryUUIDs"`
	Features      []Feature       `json:"features" bson:"features"`
	Prices        []Price         `json:"prices" bson:"prices"`
	Location      Location        `json:"location" bson:"location"`
	Boosts        []Boost         `json:"boosts" bson:"boosts"`
	People        People          `json:"people" bson:"people"`
	Type          Type            `json:"type" bson:"type"`
	Count         *int            `json:"count" bson:"count"`
	Order         *int            `json:"order" bson:"order"`
	IsActive      bool            `json:"isActive" bson:"is_active"`
	IsDeleted     bool            `json:"isDeleted" bson:"is_deleted"`
	IsValid       bool            `json:"isValid" bson:"is_valid"`
	CreatedAt     time.Time       `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time       `json:"updatedAt" bson:"updated_at"`
}

func (d *PostPriceValidationDto) ToEntity() Price {
	start, _ := time.Parse(formats.DateYYYYMMDD, d.StartDate)
	end, _ := time.Parse(formats.DateYYYYMMDD, d.EndDate)
	return Price{
		Price:     d.Price,
		StartDate: start,
		EndDate:   end,
	}
}
