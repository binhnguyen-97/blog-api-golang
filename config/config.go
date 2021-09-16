package config

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type configStruct struct {
	Port    string `mapstructure:"port"`
	MongoDb struct {
		URL      string `mapstructure:"url"`
		Database string `mapstructure:"database"`
	} `mapstructure:"mongodb"`
	Collections struct {
		Article string `mapstructure:"article"`
		Writer  string `mapstructure:"writer"`
		Uer     string `mapstructure:"user"`
	} `mapstructure:"collections"`
	JWT struct {
		SecretKey  string        `mapstructure:"secret_key"`
		ExpireTime time.Duration `mapstructure:"expire_time"`
	}
	MailService struct {
		Email    string `mapstructure:"email"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
	} `mapstructure:"mail_service"`
	Redis struct {
		Host     string `mapstructure:"host"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`
}

var Config configStruct

// GetVariableConfig: func that get variable from env and parse it to your project
func GetVariableConfig() {
	viper.SetConfigName("blog-api-go-env")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("/Users/binhnguyen/WorkSpace/Persional/")
	viper.SetConfigType("yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
