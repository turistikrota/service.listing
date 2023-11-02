package post

import "time"

type Entity struct {
	UUID          string          `json:"uuid" bson:"_id,omitempty"`
	Owner         Owner           `json:"owner" bson:"owner" validate:"required,dive,required"`
	Images        []Image         `json:"images" bson:"images" validate:"required,min=1,max=30,dive,required"`
	Meta          map[Locale]Meta `json:"meta" bson:"meta" validate:"required,dive,required"`
	CategoryUUIDs []string        `json:"categoryUUIDs" bson:"categoryUUIDs" validate:"required,min=1,max=30,dive,required"`
	Features      []Feature       `json:"features" bson:"features" validate:"required,min=1,max=30,dive,required"`
	Prices        []Price         `json:"prices" bson:"prices" validate:"required,min=1,max=100,dive,required"`
	Location      Location        `json:"location" bson:"location" validate:"required,dive,required"`
	Boosts        []Boost         `json:"boosts" bson:"boosts" validate:"required,min=0,max=30,dive,required"`
	Validation    Validation      `json:"validation" bson:"validation" validate:"required,dive,required"`
	Order         *int            `json:"order" bson:"order" validate:"required,min=0,max=1000"`
	IsActive      bool            `json:"isActive" bson:"is_active"`
	IsDeleted     bool            `json:"isDeleted" bson:"is_deleted"`
	IsValid       bool            `json:"isValid" bson:"is_valid"`
	CreatedAt     time.Time       `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time       `json:"updatedAt" bson:"updated_at"`
}

type Owner struct {
	UUID     string `json:"uuid"`
	NickName string `json:"nickName"`
}

type Image struct {
	Url   string `json:"url" bson:"url" validate:"required,url"`
	Order *int16 `json:"order" bson:"order" validate:"required,min=0,max=20"`
}

type Meta struct {
	Description string `json:"description"  validate:"required,max=255,min=3"`
	Title       string `json:"title" validate:"required,max=255,min=3"`
	Slug        string `json:"slug"`
}

type Feature struct {
	CategoryInputUUID string      `json:"categoryInputUUID" bson:"category_input_uuid" validate:"required,uuid"`
	Value             interface{} `json:"value" bson:"value" validate:"required"`
	IsPayed           *bool       `json:"isPayed" bson:"is_payed" validate:"required"`
	Price             float64     `json:"price" bson:"price" validate:"omitempty,gt=0"`
}

type Price struct {
	StartDate time.Time `json:"startDate" bson:"start_date"`
	EndDate   time.Time `json:"endDate" bson:"end_date"`
	Price     float64   `json:"price" bson:"price"`
}

type Location struct {
	Country     string    `json:"country" validate:"required,max=255,min=3"`
	City        string    `json:"city" validate:"required,max=255,min=3"`
	Street      string    `json:"street" validate:"required,max=255,min=3"`
	Address     string    `json:"address" validate:"required,max=255,min=3"`
	IsStrict    *bool     `json:"isStrict" bson:"is_strict" validate:"required"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates" validate:"required,min=2,max=2,dive,required,min=-180,max=180"`
}

type Validation struct {
	MinAdult   *int  `json:"minAdult" bson:"min_adult" validate:"required,min=1,max=50,ltefield=MaxAdult"`
	MaxAdult   *int  `json:"maxAdult" bson:"max_adult" validate:"required,min=0,max=50,gtefield=MinAdult"`
	MinKid     *int  `json:"minKid" bson:"min_kid" validate:"required,min=0,max=50,ltefield=MaxKid"`
	MaxKid     *int  `json:"maxKid" bson:"max_kid" validate:"required,min=0,max=50,gtefield=MinKid"`
	MinBaby    *int  `json:"minBaby" bson:"min_baby" validate:"required,min=0,max=50,ltefield=MaxBaby"`
	MaxBaby    *int  `json:"maxBaby" bson:"max_baby" validate:"required,min=0,max=50,gtefield=MinBaby"`
	MinDate    *int  `json:"minDate" bson:"min_date" validate:"required,min=0,max=50,ltefield=MaxDate"`
	MaxDate    *int  `json:"maxDate" bson:"max_date" validate:"required,min=0,max=50,gtefield=MinDate"`
	OnlyFamily *bool `json:"onlyFamily" bson:"only_family" validate:"required"`
	NoPet      *bool `json:"noPet" bson:"no_pet" validate:"required"`
	NoSmoke    *bool `json:"noSmoke" bson:"no_smoke" validate:"required"`
	NoAlcohol  *bool `json:"noAlcohol" bson:"no_alcohol" validate:"required"`
}

type Boost struct {
	UUID      string    `json:"uuid"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type Locale string

const (
	LocaleEN Locale = "en"
	LocaleTR Locale = "tr"
)

func (l Locale) String() string {
	return string(l)
}
