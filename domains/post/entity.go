package post

import "time"

type Entity struct {
	UUID          string          `json:"uuid"`
	Owner         Owner           `json:"owner"`
	Images        []Image         `json:"images"`
	Meta          map[Locale]Meta `json:"meta"`
	CategoryUUIDs []string        `json:"categoryUUIDs"`
	Features      []Feature       `json:"features"`
	Prices        []Price         `json:"prices"`
	Location      Location        `json:"location"`
	Boosts        []Boost         `json:"boosts"`
	People        People          `json:"people"`
	Type          Type            `json:"type"`
	Count         int             `json:"count"`
	Order         int             `json:"order"`
	IsActive      bool            `json:"isActive"`
	IsDeleted     bool            `json:"isDeleted"`
	IsValid       bool            `json:"isValid"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
}

type Owner struct {
	UUID     string `json:"uuid"`
	NickName string `json:"nickName"`
}

type Image struct {
	Url   string `json:"url"`
	Order int16  `json:"order"`
}

type People struct {
	MinAdult int `json:"minAdult"`
	MaxAdult int `json:"maxAdult"`
	MinKid   int `json:"minKid"`
	MaxKid   int `json:"maxKid"`
	MinBaby  int `json:"minBaby"`
	MaxBaby  int `json:"maxBaby"`
}

type Meta struct {
	Description string `json:"description"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
}

type Feature struct {
	CategoryInputUUID string      `json:"categoryInputUUID"`
	Value             interface{} `json:"value"`
	IsPayed           bool        `json:"isPayed"`
}

type Price struct {
	StartDate           time.Time `json:"startDate"`
	EndDate             time.Time `json:"endDate"`
	Price               float64   `json:"price"`
	IsCurrencyProtected bool      `json:"isCurrencyProtected"`
}

type Location struct {
	Country     string    `json:"country"`
	City        string    `json:"city"`
	Street      string    `json:"street"`
	Address     string    `json:"address"`
	IsStrict    bool      `json:"isStrict"`
	Coordinates []float64 `json:"coordinates"`
}

type Boost struct {
	UUID      string    `json:"uuid"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type Type string

type Locale string

const (
	LocaleEN Locale = "en"
	LocaleTR Locale = "tr"
)

const (
	TypeEstate     Type = "estate"
	TypeCar        Type = "car"
	TypeBoat       Type = "boat"
	TypeMotorcycle Type = "motorcycle"
	TypeOther      Type = "other"
)

func (t Type) String() string {
	return string(t)
}

func (l Locale) String() string {
	return string(l)
}
