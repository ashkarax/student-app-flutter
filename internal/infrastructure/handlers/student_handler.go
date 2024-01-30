package handlers

import (
	"net/http"

	interfaceUseCase "github.com/ashkarax/student_data_managing/internal/infrastructure/usecases/interfaces"
	requestmodels "github.com/ashkarax/student_data_managing/internal/models/request_models"
	responsemodels "github.com/ashkarax/student_data_managing/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	StudentUsecase interfaceUseCase.IstudentUseCase
}

func NewStudentHandler(studentUsecase interfaceUseCase.IstudentUseCase) *StudentHandler {
	return &StudentHandler{StudentUsecase: studentUsecase}
}

func (u *StudentHandler) AddStudent(c *gin.Context) {
	var studentData requestmodels.NewStudent
	if err := c.ShouldBind(&studentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationRes,err := u.StudentUsecase.AddStudent(&studentData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't add Student", validationRes, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "student added succesfully", nil, nil)
	c.JSON(http.StatusOK, response)
}
