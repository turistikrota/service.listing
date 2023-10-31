package config

type MongoPost struct {
	Host       string `env:"MONGO_POST_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_POST_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_POST_USERNAME" envDefault:""`
	Password   string `env:"MONGO_POST_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_POST_DATABASE" envDefault:"account"`
	Collection string `env:"MONGO_POST_COLLECTION" envDefault:"accounts"`
	Query      string `env:"MONGO_POST_QUERY" envDefault:""`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Http struct {
	Host  string `env:"SERVER_HOST" envDefault:"localhost"`
	Port  int    `env:"SERVER_PORT" envDefault:"3000"`
	Group string `env:"SERVER_GROUP" envDefault:"account"`
}

type Redis struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
	Pw   string `env:"REDIS_PASSWORD"`
	Db   int    `env:"REDIS_DB"`
}

type CacheRedis struct {
	Host string `env:"REDIS_CACHE_HOST"`
	Port string `env:"REDIS_CACHE_PORT"`
	Pw   string `env:"REDIS_CACHE_PASSWORD"`
	Db   int    `env:"REDIS_CACHE_DB"`
}

type HttpHeaders struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	Domain           string `env:"HTTP_HEADER_DOMAIN" envDefault:"*"`
}

type TokenSrv struct {
	Expiration int    `env:"TOKEN_EXPIRATION" envDefault:"3600"`
	Project    string `env:"TOKEN_PROJECT" envDefault:"empty"`
}

type Session struct {
	Topic string `env:"SESSION_TOPIC"`
}

type RSA struct {
	PrivateKeyFile string `env:"RSA_PRIVATE_KEY"`
	PublicKeyFile  string `env:"RSA_PUBLIC_KEY"`
}

type Topics struct {
	Post     PostTopics
	Category CategoryTopics
}

type PostTopics struct {
	Created   string `env:"STREAMING_TOPIC_POST_CREATED"`
	Updated   string `env:"STREAMING_TOPIC_POST_UPDATED"`
	Deleted   string `env:"STREAMING_TOPIC_POST_DELETED"`
	Disabled  string `env:"STREAMING_TOPIC_POST_DISABLED"`
	Enabled   string `env:"STREAMING_TOPIC_POST_ENABLED"`
	ReOrdered string `env:"STREAMING_TOPIC_POST_REORDERED"`
	Restored  string `env:"STREAMING_TOPIC_POST_RESTORED"`
}

type CategoryTopics struct {
	PostValidationSuccess string `env:"STREAMING_TOPIC_CATEGORY_POST_VALIDATION_SUCCESS"`
}

type Nats struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type CDN struct {
	Url string `env:"CDN_URL" envDefault:"http://localhost:3000"`
}

type App struct {
	Protocol string `env:"PROTOCOL" envDefault:"http"`
	DB       struct {
		Post MongoPost
	}
	Http        Http
	HttpHeaders HttpHeaders
	I18n        I18n
	Topics      Topics
	Session     Session
	Nats        Nats
	Redis       Redis
	TokenSrv    TokenSrv
	CacheRedis  CacheRedis
	CDN         CDN
	RSA         RSA
}
