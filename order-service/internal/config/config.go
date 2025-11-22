package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.yaml.in/yaml/v3"
)

type Config struct {
	Env 	  string   `yaml:"env"`
	DB 		  DBConfig
	AppSecret string   `yaml:"app_secret"`
	HTTPServer		   `yaml:"http_server"`	
	SSOServer		   `yaml:"sso_server"`
}

type DBConfig struct {
	URL string `validate:"required"`
}

type HTTPServer struct {
	Address 	string `yaml:"addr" validate:"required"`
	Timeout 	string `yaml:"timeout"`
	IdleTimeout string `yaml:"idle_timeout"`
}

type SSOServer struct {
	GRPCAddr 	string `yaml:"grpc_addr" validate:"required"`
	GRPCTimeout string `yaml:"grpc_timeout"`
	Retries 	string `yaml:"retries"`
}

func MustLoad() *Config {
	err := godotenv.Load("internal/config/.env")
	if err != nil {
		log.Printf(".env file not found or failed to load: %v", err)
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("config path variable is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exists: " + configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	var cfg Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	dbUrl := os.Getenv("DB_URL") //TODO: through envconfig
	if dbUrl == "" {
		log.Fatal("database url variable is not set")
	}

	cfg.DB.URL = dbUrl

	if err := validator.New().Struct(cfg); err != nil {
		log.Fatalf("error validating config: %v", err)
	}

	return &cfg
}