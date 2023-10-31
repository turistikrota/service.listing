package post

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FilterEntity struct {
	Locale      string           `json:"-"`
	Query       string           `json:"query,omitempty" validate:"omitempty,max=100"`
	Price       *FilterPrice     `json:"price,omitempty" validate:"omitempty"`
	People      *FilterPeople    `json:"people,omitempty" validate:"omitempty"`
	Coordinates []float64        `json:"coordinates,omitempty" validate:"omitempty,min=2,max=2"`
	Distance    *float64         `json:"distance,omitempty" validate:"omitempty,gt=6,lt=16"`
	Features    []*FilterFeature `json:"features,omitempty" validate:"omitempty,dive"`
	Types       []Type           `json:"types,omitempty" validate:"omitempty"`
	Categories  []string         `json:"categories,omitempty" validate:"omitempty,dive,object_id"`
	Sort        Sort             `json:"sort,omitempty" validate:"omitempty,oneof=most_recent nearest"`
	Order       Order            `json:"order,omitempty" validate:"omitempty,oneof=asc desc"`
}

type FilterPrice struct {
	Min       float64    `json:"min" validate:"omitempty,gt=0"`
	Max       float64    `json:"max" validate:"omitempty,gt=0"`
	StartDate *time.Time `json:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDate   *time.Time `json:"end_date" validate:"omitempty,datetime=2006-01-02"`
}

type FilterPeople struct {
	Adult int `json:"adult" validate:"omitempty,gt=0"`
	Kid   int `json:"kid" validate:"omitempty,gt=0"`
	Baby  int `json:"baby" validate:"omitempty,gt=0"`
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
		list = r.filterByOwner(list, nickName)
	}
	list = r.filterByTypes(list, filter)
	list = r.filterByLocation(list, filter)
	list = r.filterByCategory(list, filter)
	list = r.filterByQuery(list, filter)
	list = r.filterByPrice(list, filter)
	list = r.filterByPeople(list, filter)
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

func (r *repo) filterByTypes(list []bson.M, filter FilterEntity) []bson.M {
	if len(filter.Types) > 0 {
		list = append(list, bson.M{
			fields.Type: bson.M{
				"$in": filter.Types,
			},
		})
	}
	return list
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

func (r *repo) ownerFilter(nickName string) bson.M {
	return bson.M{
		ownerField(ownerFields.NickName): nickName,
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

func (r *repo) filterByOwner(list []bson.M, nickName string) []bson.M {
	return append(list, r.ownerFilter(nickName))
}

func (r *repo) filterByQuery(list []bson.M, filter FilterEntity) []bson.M {
	if filter.Query != "" {
		list = append(list, bson.M{
			"$or": []bson.M{
				{
					metaField(metaFields.Locale, metaFields.Title): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					metaField(metaFields.Locale, metaFields.Description): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
			},
		})
	}
	return list
}

func (r *repo) filterByPeople(list []bson.M, filter FilterEntity) []bson.M {
	if filter.People != nil {
		peopleFilters := make([]bson.M, 0)
		if filter.People.Adult != 0 {
			peopleFilters = append(peopleFilters, bson.M{
				"$and": []bson.M{
					{
						peopleField(peopleFields.MinAdult): bson.M{
							"$lte": filter.People.Adult,
						},
					},
					{
						peopleField(peopleFields.MaxAdult): bson.M{
							"$gte": filter.People.Adult,
						},
					},
				},
			})
		}
		if filter.People.Kid != 0 {
			peopleFilters = append(peopleFilters, bson.M{
				"$and": []bson.M{
					{
						peopleField(peopleFields.MinKid): bson.M{
							"$lte": filter.People.Kid,
						},
					},
					{
						peopleField(peopleFields.MaxKid): bson.M{
							"$gte": filter.People.Kid,
						},
					},
				},
			})
		}
		if filter.People.Baby != 0 {
			peopleFilters = append(peopleFilters, bson.M{
				"$and": []bson.M{
					{
						peopleField(peopleFields.MinBaby): bson.M{
							"$lte": filter.People.Baby,
						},
					},
					{
						peopleField(peopleFields.MaxBaby): bson.M{
							"$gte": filter.People.Baby,
						},
					},
				},
			})
		}
		if len(peopleFilters) > 0 {
			if len(peopleFilters) == 1 {
				list = append(list, peopleFilters[0])
			} else {
				list = append(list, bson.M{
					"$and": peopleFilters,
				})
			}
		}
	}
	return list
}

func (r *repo) filterByPrice(list []bson.M, filter FilterEntity) []bson.M {
	if filter.Price != nil {
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
					"$lte": filter.Price.StartDate,
				},
			})
		}
		if filter.Price.EndDate != nil {
			priceFilters = append(priceFilters, bson.M{
				priceField(priceFields.EndDate): bson.M{
					"$gte": filter.Price.EndDate,
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
