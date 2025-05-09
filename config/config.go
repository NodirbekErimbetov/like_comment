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
	config.PostgresDatabase = "minimedium"
	config.PostgresPassword = "0021"
	config.PostgresUser = "nodirbek"
	config.PostgresHost = "localhost"
	config.ServiceHost = "localhost"
	config.ServiceHTTPPort = ":8081"

	return config
}
