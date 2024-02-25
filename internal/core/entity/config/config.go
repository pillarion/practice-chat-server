package config

// Config holds the configuration for the application.
type Config struct {
	GRPC     GRPC     `env:"grpc"`
	Database Database `env:"database"`
}

// GRPC holds the configuration for the gRPC server.
type GRPC struct {
	Port int `env:"GRPC_PORT"`
}

// Database holds the configuration for the database.
type Database struct {
	Host string `env:"host"`
	Port string `env:"port"`
	Db   string `env:"db"`
	User string `env:"user"`
	Pass string `env:"pass"`
}
