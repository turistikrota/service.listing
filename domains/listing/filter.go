package listing

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FilterEntity struct {
	Locale       string            `json:"-"`
	Query        string            `json:"query,omitempty" validate:"omitempty,max=100"`
	Price        *FilterPrice      `json:"price,omitempty" validate:"omitempty"`
	Validation   *FilterValidation `json:"validation,omitempty" validate:"omitempty"`
	Coordinates  []float64         `json:"coordinates,omitempty" validate:"omitempty,min=2,max=2"`
	Distance     *float64          `json:"distance,omitempty" validate:"omitempty,gt=6,lt=16"`
	Features     []*FilterFeature  `json:"features,omitempty" validate:"omitempty,dive"`
	Categories   []string          `json:"categories,omitempty" validate:"omitempty,dive,object_id"`
	Sort         Sort              `json:"sort,omitempty" validate:"omitempty,oneof=most_recent nearest"`
	Order        Order             `json:"order,omitempty" validate:"omitempty,oneof=asc desc"`
	StartDate    *time.Time        `json:"-"`
	EndDate      *time.Time        `json:"-"`
	StartDateStr string            `json:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDateStr   string            `json:"end_date" validate:"omitempty,datetime=2006-01-02"`
}

func (e *FilterEntity) Parse() {
	if e.StartDateStr != "" {
		t, _ := time.Parse("2006-01-02", e.StartDateStr)
		e.StartDate = &t
	}
	if e.EndDateStr != "" {
		t, _ := time.Parse("2006-01-02", e.EndDateStr)
		e.EndDate = &t
	}
}

type FilterPrice struct {
	Min          float64    `json:"min" validate:"omitempty,gt=0"`
	Max          float64    `json:"max" validate:"omitempty,gt=0"`
	StartDateStr string     `json:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDateStr   string     `json:"end_date" validate:"omitempty,datetime=2006-01-02"`
	StartDate    *time.Time `json:"-"`
	EndDate      *time.Time `json:"-"`
}

type FilterValidation struct {
	Adult     int   `json:"adult" validate:"omitempty,gt=0"`
	Kid       int   `json:"kid" validate:"omitempty,gt=0"`
	Baby      int   `json:"baby" validate:"omitempty,gt=0"`
	Family    *bool `json:"family" validate:"omitempty"`
	Pet       *bool `json:"pet" validate:"omitempty"`
	Smoke     *bool `json:"smoke" validate:"omitempty"`
	Alcohol   *bool `json:"alcohol" validate:"omitempty"`
	Party     *bool `json:"party" validate:"omitempty"`
	Unmarried *bool `json:"unmarried" validate:"omitempty"`
	Guest     *bool `json:"guest" validate:"omitempty"`
}

type FilterFeature struct {
	CategoryInputUUID string
	Value             interface{}
}

type (
	Sort  string
	Order string
)

const (
	SortByMostRecent Sort = "most_recent"
	SortByNearest    Sort = "nearest"
)

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

func (s Sort) IsValid() bool {
	return s == SortByMostRecent ||
		s == SortByNearest
}

func (p *FilterPrice) Parse() {
	if p.StartDateStr != "" {
		t, _ := time.Parse("2006-01-02", p.StartDateStr)
		p.StartDate = &t
	}
	if p.EndDateStr != "" {
		t, _ := time.Parse("2006-01-02", p.EndDateStr)
		p.EndDate = &t
	}
}

func (o Order) IsValid() bool {
	return o == OrderAsc ||
		o == OrderDesc
}

func (e *FilterEntity) GetPerfectDistance() float64 {
	if e.Distance == nil {
		return 100
	}
	distances := map[float64]float64{
		7:  500,
		8:  300,
		9:  200,
		10: 100,
		11: 50,
		12: 20,
		13: 10,
		14: 5,
		15: 3,
	}
	if distance, ok := distances[*e.Distance]; ok {
		return distance
	}
	return 10
}

func (r *repo) filterToBson(filter FilterEntity, nickName string) bson.M {
	list := make([]bson.M, 0)
	if nickName != "" {
		list = r.filterByBusiness(list, nickName)
	}
	list = r.filterByLocation(list, filter)
	list = r.filterByCategory(list, filter)
	list = r.filterByQuery(list, filter)
	list = r.filterByPrice(list, filter)
	list = r.filterByValidation(list, filter)
	listLen := len(list)
	if listLen == 0 {
		return bson.M{}
	}
	if listLen == 1 {
		return list[0]
	}
	return bson.M{
		"$and": list,
	}
}

func (r *repo) filterByLocation(list []bson.M, filter FilterEntity) []bson.M {
	if filter.Coordinates != nil && len(filter.Coordinates) == 2 {
		distance := filter.GetPerfectDistance()
		radius := distance / 6378.1
		list = append(list, bson.M{
			locationField(locationFields.Coordinates): bson.M{
				"$geoWithin": bson.M{
					"$centerSphere": []interface{}{
						filter.Coordinates,
						radius,
					},
				},
			},
		})
	}
	return list
}

func (r *repo) filterByCategory(list []bson.M, filter FilterEntity) []bson.M {
	if filter.Categories != nil && len(filter.Categories) > 0 {
		list = append(list, bson.M{
			fields.CategoryUUIDs: bson.M{
				"$in": filter.Categories,
			},
		})
	}
	return list
}

func (r *repo) businessFilter(nickName string) bson.M {
	return bson.M{
		businessField(businessFields.NickName): nickName,
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
		fields.IsActive: true,
		fields.IsValid:  true,
	}
}

func (r *repo) baseFilter() bson.M {
	return bson.M{
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
		fields.IsActive: true,
		fields.IsValid:  true,
	}
}

func (r *repo) filterByBusiness(list []bson.M, nickName string) []bson.M {
	return append(list, r.businessFilter(nickName))
}

func (r *repo) filterByQuery(list []bson.M, filter FilterEntity) []bson.M {
	if filter.Query != "" {
		list = append(list, bson.M{
			"$or": []bson.M{
				{
					metaField(filter.Locale, metaFields.Title): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					metaField(filter.Locale, metaFields.Description): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
			},
		})
	}
	return list
}

func (r *repo) filterByValidation(list []bson.M, filter FilterEntity) []bson.M {
	if filter.Validation != nil {
		validationFilters := make([]bson.M, 0)
		if filter.Validation.Adult != 0 {
			validationFilters = append(validationFilters, bson.M{
				"$and": []bson.M{
					{
						validationField(validationFields.MinAdult): bson.M{
							"$lte": filter.Validation.Adult,
						},
					},
					{
						validationField(validationFields.MaxAdult): bson.M{
							"$gte": filter.Validation.Adult,
						},
					},
				},
			})
		}
		if filter.Validation.Kid != 0 {
			validationFilters = append(validationFilters, bson.M{
				"$and": []bson.M{
					{
						validationField(validationFields.MinKid): bson.M{
							"$lte": filter.Validation.Kid,
						},
					},
					{
						validationField(validationFields.MaxKid): bson.M{
							"$gte": filter.Validation.Kid,
						},
					},
				},
			})
		}
		if filter.Validation.Baby != 0 {
			validationFilters = append(validationFilters, bson.M{
				"$and": []bson.M{
					{
						validationField(validationFields.MinBaby): bson.M{
							"$lte": filter.Validation.Baby,
						},
					},
					{
						validationField(validationFields.MaxBaby): bson.M{
							"$gte": filter.Validation.Baby,
						},
					},
				},
			})
		}
		if filter.Validation.Family != nil {
			validationFilters = append(validationFilters, bson.M{
				validationField(validationFields.OnlyFamily): !*filter.Validation.Family,
			})
		}
		if filter.Validation.Pet != nil {
			validationFilters = append(validationFilters, bson.M{
				validationField(validationFields.NoPet): !*filter.Validation.Pet,
			})
		}
		if filter.Validation.Smoke != nil {
			validationFilters = append(validationFilters, bson.M{
				validationField(validationFields.NoSmoke): !*filter.Validation.Smoke,
			})
		}
		if filter.Validation.Alcohol != nil {
			validationFilters = append(validationFilters, bson.M{
				validationField(validationFields.NoAlcohol): !*filter.Validation.Alcohol,
			})
		}
		if filter.Validation.Party != nil {
			validationFilters = append(validationFilters, bson.M{
				validationField(validationFields.NoParty): !*filter.Validation.Party,
			})
		}
		if filter.Validation.Unmarried != nil {
			validationFilters = append(validationFilters, bson.M{
				validationField(validationFields.NoUnmarried): !*filter.Validation.Unmarried,
			})
		}
		if filter.Validation.Guest != nil {
			validationFilters = append(validationFilters, bson.M{
				validationField(validationFields.NoGuest): !*filter.Validation.Guest,
			})
		}
		if filter.StartDate != nil && filter.EndDate != nil && filter.StartDate.Before(*filter.EndDate) {
			totalDate := filter.EndDate.Sub(*filter.StartDate).Hours() / 24
			validationFilters = append(validationFilters, bson.M{
				"$and": []bson.M{
					{
						validationField(validationFields.MinDate): bson.M{
							"$lte": totalDate,
						},
					},
					{
						validationField(validationFields.MaxDate): bson.M{
							"$gte": totalDate,
						},
					},
				},
			})
		}

		if len(validationFilters) > 0 {
			if len(validationFilters) == 1 {
				list = append(list, validationFilters[0])
			} else {
				list = append(list, bson.M{
					"$and": validationFilters,
				})
			}
		}
	}
	return list
}

func (r *repo) filterByPrice(list []bson.M, filter FilterEntity) []bson.M {
	if filter.Price != nil {
		filter.Price.Parse()
		priceFilters := make([]bson.M, 0)
		if filter.Price.Min != 0 {
			priceFilters = append(priceFilters, bson.M{
				priceField(priceFields.Price): bson.M{
					"$gte": filter.Price.Min,
				},
			})
		}
		if filter.Price.Max != 0 {
			priceFilters = append(priceFilters, bson.M{
				priceField(priceFields.Price): bson.M{
					"$lte": filter.Price.Max,
				},
			})
		}
		if filter.Price.StartDate != nil {
			priceFilters = append(priceFilters, bson.M{
				priceField(priceFields.StartDate): bson.M{
					"$gte": filter.Price.StartDate,
				},
			})
		}
		if filter.Price.EndDate != nil {
			priceFilters = append(priceFilters, bson.M{
				priceField(priceFields.EndDate): bson.M{
					"$lte": filter.Price.EndDate,
				},
			})
		}
		filterLen := len(priceFilters)
		if filterLen > 0 {
			if filterLen == 1 {
				list = append(list, priceFilters[0])
			} else {
				list = append(list, bson.M{
					"$and": priceFilters,
				})
			}
		}
	}
	return list
}

func (r *repo) sort(opts *options.FindOptions, filter FilterEntity) *options.FindOptions {
	order := -1
	if filter.Order == OrderAsc {
		order = 1
	}
	field := fields.UpdatedAt
	switch filter.Sort {
	case SortByMostRecent:
		field = fields.UpdatedAt
	case SortByNearest:
		field = locationField(locationFields.Coordinates)
	}
	opts.SetSort(bson.D{{Key: field, Value: order}})
	return opts
}
