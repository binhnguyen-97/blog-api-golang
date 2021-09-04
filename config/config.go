package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type configStruct struct {
	PORT                    string
	MONGODB_URL             string
	DATABASE_NAME           string
	ARTICLE_COLLECTION_NAME string
	WRITER_COLLECTION_NAME  string
	USER_COLLECTION_NAME    string
	JWT_SECRET_KEY          string
}

var Config configStruct

// GetVariableConfig: func that get variable from env and parse it to your project
func GetVariableConfig() {
	v := viper.New()
	v.SetConfigName("blog-api-go-env")
	v.AddConfigPath(".")
	v.AddConfigPath("../")
	v.SetConfigType("yaml")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.BindEnv("PORT")
	v.BindEnv("MONGODB_URL")
	v.BindEnv("DATABASE_NAME")
	v.BindEnv("ARTICLE_COLLECTION_NAME")
	v.BindEnv("WRITER_COLLECTION_NAME")
	v.BindEnv("USER_COLLECTION_NAME")
	v.BindEnv("JWT_SECRET_KEY")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := v.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

}
