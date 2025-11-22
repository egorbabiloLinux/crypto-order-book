package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"go.yaml.in/yaml/v3"
)

type Config struct {
	Env 	  string   `yaml:"env"`
	DB 		  DBConfig `yaml:"db"`
	AppSecret string   `yaml:"app_secret"`
	HTTPServer		   `yaml:"http_server"`	
	SSOServer		   `yaml:"sso_server"`
}

type DBConfig struct {
	URL string `yaml:"url" validate:"required"`
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

	var config Config
	if err = yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}
	//задать значения по дефолту 

	if err := validator.New().Struct(config); err != nil {
		log.Fatalf("error validating config: %v", err)
	}

	return &config
}