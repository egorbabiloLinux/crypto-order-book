package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"go.yaml.in/yaml/v3"
)

type Config struct {
	Env 	  string   `mapstructure:"env"`
	DB 		  DBConfig `mapstructur: "db"`
	AppSecret string   `mapstructure:"app_secret"`
	HTTPServer		   `mapstructure:"http_server"`	
	SSOServer		   `mapstructure:"sso_server"`
}

type DBConfig struct {
	URL string `mapstructure:"url" validate:"required"`
}

type HTTPServer struct {
	Address 	string `mapstructure:"address" validate:"required"`
	Timeout 	string `mapstructre:"timeout"`
	IdleTimeout string `mapstructure:"idle_timeout"`
}

type SSOServer struct {
	GRPCAddr 	string `mapstructure:"grpc_addr" validate:"required"`
	GRPCTimeout string `mapstructure:"grpc_timeout"`
	Retries 	string `mapstructure:"retries"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		//log.Fatal()
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