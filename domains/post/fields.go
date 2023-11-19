package post

type fieldsType struct {
	UUID          string
	Business      string
	Images        string
	Meta          string
	CategoryUUIDs string
	Features      string
	Prices        string
	Location      string
	Boosts        string
	Validation    string
	Order         string
	IsActive      string
	IsDeleted     string
	IsValid       string
	CreatedAt     string
	UpdatedAt     string
}

type businessFieldsType struct {
	UUID     string
	NickName string
}

type imageFieldsType struct {
	Url   string
	Order string
}

type peopleFieldsType struct {
	MinAdult string
	MaxAdult string
	MinKid   string
	MaxKid   string
	MinBaby  string
	MaxBaby  string
}

type metaFieldsType struct {
	Locale      string
	Description string
	Title       string
	Slug        string
	MarkdownURL string
}

type featureFieldsType struct {
	CategoryInputUUID string
	Value             string
	IsPayed           string
}

type priceFieldsType struct {
	StartDate string
	EndDate   string
	Price     string
	Currency  string
}

type locationFieldsType struct {
	Country     string
	City        string
	Street      string
	Address     string
	IsStrict    string
	Coordinates string
}

type coordinateFieldsType struct {
	Latitude  string
	Longitude string
}

type validationFieldsType struct {
	MinAdult    string
	MaxAdult    string
	MinKid      string
	MaxKid      string
	MinBaby     string
	MaxBaby     string
	MinDate     string
	MaxDate     string
	OnlyFamily  string
	NoPet       string
	NoSmoke     string
	NoAlcohol   string
	NoParty     string
	NoUnmarried string
	NoGuest     string
}

type boostFieldsType struct {
	UUID      string
	StartDate string
	EndDate   string
}

var fields = fieldsType{
	UUID:          "_id",
	Business:      "business",
	Images:        "images",
	Meta:          "meta",
	CategoryUUIDs: "category_uuids",
	Features:      "features",
	Prices:        "prices",
	Location:      "location",
	Boosts:        "boosts",
	Order:         "order",
	Validation:    "validation",
	IsActive:      "is_active",
	IsDeleted:     "is_deleted",
	IsValid:       "is_valid",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

var businessFields = businessFieldsType{
	UUID:     "uuid",
	NickName: "nick_name",
}

var imageFields = imageFieldsType{
	Url:   "url",
	Order: "order",
}

var metaFields = metaFieldsType{
	Locale:      "locale",
	Description: "description",
	Title:       "title",
	Slug:        "slug",
	MarkdownURL: "markdown_url",
}

var featureFields = featureFieldsType{
	CategoryInputUUID: "category_input_uuid",
	Value:             "value",
	IsPayed:           "is_payed",
}

var priceFields = priceFieldsType{
	StartDate: "start_date",
	EndDate:   "end_date",
	Price:     "price",
	Currency:  "currency",
}

var locationFields = locationFieldsType{
	Country:     "country",
	City:        "city",
	Street:      "street",
	Address:     "address",
	IsStrict:    "is_strict",
	Coordinates: "coordinates",
}

var boostFields = boostFieldsType{
	UUID:      "uuid",
	StartDate: "start_date",
	EndDate:   "end_date",
}

var validationFields = validationFieldsType{
	MinAdult:    "min_adult",
	MaxAdult:    "max_adult",
	MinKid:      "min_kid",
	MaxKid:      "max_kid",
	MinBaby:     "min_baby",
	MaxBaby:     "max_baby",
	MinDate:     "min_date",
	MaxDate:     "max_date",
	OnlyFamily:  "only_family",
	NoPet:       "no_pet",
	NoSmoke:     "no_smoke",
	NoAlcohol:   "no_alcohol",
	NoParty:     "no_party",
	NoUnmarried: "no_unmarried",
	NoGuest:     "no_guest",
}

func businessField(field string) string {
	return fields.Business + "." + field
}

func metaField(locale string, field string) string {
	return fields.Meta + "." + locale + "." + field
}

func locationField(field string) string {
	return fields.Location + "." + field
}

func priceField(field string) string {
	return fields.Prices + "." + field
}

func featureField(field string) string {
	return fields.Features + "." + field
}

func validationField(field string) string {
	return fields.Validation + "." + field
}
