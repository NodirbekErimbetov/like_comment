package config

type Config struct {
	PostgresPort     string
	PostgresPassword string
	PostgresDatabase string
	PostgresUser     string
	PostgresHost     string
	ServiceHost      string
	ServiceHTTPPort  string
}

func Load() Config {
	var config Config
	config.PostgresPort = "5432"
	config.PostgresDatabase = "project"
	config.PostgresPassword = "0021"
	config.PostgresUser = "nodirbek"
	config.PostgresHost = "localhost"
	config.ServiceHost = "localhost"
	config.ServiceHTTPPort = ":8080"

	return config
}
