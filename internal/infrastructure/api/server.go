package server

import (
	"fmt"
	"log"

	"github.com/ashkarax/student_data_managing/internal/infrastructure/handlers"
	"github.com/gin-gonic/gin"
)

type ServerHttp struct {
	engin *gin.Engine
}

func NewServerHttp(studentHandler *handlers.StudentHandler) *ServerHttp {
	engin := gin.Default()

	engin.POST("/", studentHandler.AddStudent)
	//get request to output data
	//search with id/name to find a student
	//delete student records
	//edit student records

	return &ServerHttp{engin: engin}
}

func (server *ServerHttp) Start() {
	err := server.engin.Run(":8082")
	if err != nil {
		log.Fatal("gin engin couldn't start")
	}
	fmt.Println("gin engin start")
}
