package main

import (
	"blog-api-golang/config"
	"blog-api-golang/db"
	"blog-api-golang/routers"
	"context"
	"time"
)

func main() {

	// Get config from yaml and process env file
	config.GetVariableConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// Connect to mongodb
	db.ConnectToDatabase(ctx)

	defer db.Disconnect(ctx)

	routers.InitRouter().Run(config.Config.Port)
}
