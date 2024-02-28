package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ashkarax/student_data_managing/internal/infrastructure/handlers"
	"github.com/gin-gonic/gin"
)

var validAPIKeys = map[string]bool{
	"apikey@studentapp": true,
}

type ServerHttp struct {
	engin *gin.Engine
}

func NewServerHttp(studentHandler *handlers.StudentHandler) *ServerHttp {
	engin := gin.Default()

	// Middleware to validate API key
	engin.Use(func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if !validAPIKeys[apiKey] {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}
		c.Next()
	})

	engin.POST("/", studentHandler.AddStudent)
	engin.GET("/", studentHandler.GetStudentDetails)
	engin.GET("/pagination", studentHandler.GetStudentDetailsPagination)
	engin.DELETE("/", studentHandler.DeleteStudentDetails)
	engin.PATCH("/", studentHandler.EditStudentDetails)
	engin.GET("/search", studentHandler.SearchByNameRollNo)


	//search with id/name to find a student
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
