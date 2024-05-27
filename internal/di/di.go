package di

import (
	"fmt"

	"github.com/ashkarax/student_data_managing/internal/config"
	server "github.com/ashkarax/student_data_managing/internal/infrastructure/api"
	"github.com/ashkarax/student_data_managing/internal/infrastructure/db"
	"github.com/ashkarax/student_data_managing/internal/infrastructure/handlers"
	"github.com/ashkarax/student_data_managing/internal/infrastructure/repository"
	"github.com/ashkarax/student_data_managing/internal/infrastructure/usecases"
	"github.com/ashkarax/student_data_managing/pkg/aws"
)

func InitializeApi(config *config.Config) (*server.ServerHttp, error) {

	DB, err := db.ConnectDatabase(config.DB)
	if err != nil {
		fmt.Println("ERROR CONNECTING DB FROM DI.GO")
		return nil, err
	}

	aws.AWSS3FileUploaderSetup(config.AwsS3)

	studentRepository := repository.NewStudentRepository(DB)
	studentUsecase := usecases.NewStudentUseCase(studentRepository)
	studentHandler := handlers.NewStudentHandler(studentUsecase)

	serverHttp := server.NewServerHttp(studentHandler, &config.ApiKey)

	return serverHttp, nil
}
