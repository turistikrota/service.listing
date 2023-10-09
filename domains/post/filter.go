package post

import "time"

type FilterEntity struct {
	Locale      string
	Query       string
	Price       *FilterPrice
	People      *FilterPeople
	Coordinates []float64
	Distance    float64
	Features    []*FilterFeature
	Type        Type
	Categories  []string
}

type FilterPrice struct {
	Min                 float64
	Max                 float64
	StartDate           *time.Time
	EndDate             *time.Time
	Currencies          []string
	IsCurrencyProtected bool
}

type FilterPeople struct {
	Adult int
	Kid   int
	Baby  int
}

type FilterFeature struct {
	CategoryInputUUID string
	Value             interface{}
}
