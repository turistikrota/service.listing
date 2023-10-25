package post

type fieldsType struct {
	UUID          string
	Owner         string
	Images        string
	Meta          string
	CategoryUUIDs string
	Features      string
	Prices        string
	Location      string
	Boosts        string
	People        string
	Type          string
	Order         string
	Count         string
	IsActive      string
	IsDeleted     string
	IsValid       string
	CreatedAt     string
	UpdatedAt     string
}

type ownerFieldsType struct {
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

type boostFieldsType struct {
	UUID      string
	StartDate string
	EndDate   string
}

var fields = fieldsType{
	UUID:          "_id",
	Owner:         "owner",
	Images:        "images",
	Meta:          "meta",
	CategoryUUIDs: "category_uuids",
	Features:      "features",
	Prices:        "prices",
	Location:      "location",
	Boosts:        "boosts",
	Type:          "type",
	Order:         "order",
	Count:         "count",
	People:        "people",
	IsActive:      "is_active",
	IsDeleted:     "is_deleted",
	IsValid:       "is_valid",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

var ownerFields = ownerFieldsType{
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

var peopleFields = peopleFieldsType{
	MinAdult: "min_adult",
	MaxAdult: "max_adult",
	MinKid:   "min_kid",
	MaxKid:   "max_kid",
	MinBaby:  "min_baby",
	MaxBaby:  "max_baby",
}

var boostFields = boostFieldsType{
	UUID:      "uuid",
	StartDate: "start_date",
	EndDate:   "end_date",
}

func ownerField(field string) string {
	return fields.Owner + "." + field
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

func peopleField(field string) string {
	return fields.People + "." + field
}