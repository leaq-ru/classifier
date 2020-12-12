package config

type Config struct {
	ServiceName string
	Grpc        grpc
	LogLevel    string `envconfig:"LOGLEVEL"`
}

type grpc struct {
	Port string `envconfig:"GRPC_PORT"`
}
