package main

import (
	"fmt"
	"log"

	"github.com/ashkarax/student_data_managing/internal/config"
	"github.com/ashkarax/student_data_managing/internal/di"
)

func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("error at loading env file using viper")
	}
	server, diErr := di.InitializeApi(config)
	if diErr != nil {
		log.Fatal("error for server creation")
	}

	server.Start()
	fmt.Println("server started on port:8085")

}
