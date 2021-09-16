package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type configStruct struct {
	PORT    string `mapstructure:"port"`
	MONGODB struct {
		URL      string `mapstructure:"url"`
		DATABASE string `mapstructure:"database"`
	} `mapstructure:"mongodb"`
	COLLECTIONS struct {
		ARTICLE string `mapstructure:"article"`
		WRITER  string `mapstructure:"writer"`
		USER    string `mapstructure:"user"`
	} `mapstructure:"collections"`
	JWT_SECRET_KEY string `mapstructure:"jwt_token"`
}

var Config configStruct

// GetVariableConfig: func that get variable from env and parse it to your project
func GetVariableConfig() {
	viper.SetConfigName("blog-api-go-env")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
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

	log.Printf("%v", Config)
}
