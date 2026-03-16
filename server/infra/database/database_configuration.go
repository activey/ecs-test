package database

type DatabaseConfiguration struct {
	ConnectionURL string `yaml:"connection_url"`
}

func NewDatabaseConfiguration() *DatabaseConfiguration {
	return &DatabaseConfiguration{
		ConnectionURL: "host=localhost user=test password=test dbname=test port=5432 sslmode=disable TimeZone=Europe/Warsaw",
	}
}
