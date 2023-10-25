package post

import (
	"time"

	"github.com/cilloparch/cillop/formats"
)

type PostPriceValidationDto struct {
	Price     float64 `json:"price" validate:"required,min=0"`
	StartDate string  `json:"startDate" validate:"required,datetime=2006-01-02"`
	EndDate   string  `json:"endDate" validate:"required,datetime=2006-01-02"`
	Currency  string  `json:"currency" validate:"required,oneof=TRY USD EUR"`
}

func (d *PostPriceValidationDto) ToEntity() Price {
	start, _ := time.Parse(formats.DateYYYYMMDD, d.StartDate)
	end, _ := time.Parse(formats.DateYYYYMMDD, d.EndDate)
	return Price{
		Price:     d.Price,
		StartDate: start,
		EndDate:   end,
		Currency: Currency(d.Currency),
	}
}
