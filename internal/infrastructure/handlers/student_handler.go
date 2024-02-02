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

	validationRes, err := u.StudentUsecase.AddStudent(&studentData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't add Student", validationRes, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "student added succesfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (u *StudentHandler) GetStudentDetails(c *gin.Context) {

	studentData, err := u.StudentUsecase.GetStudentDetails()
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't fetch students details", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := responsemodels.Responses(http.StatusOK, "students details fetched succesfully", studentData, nil)
	c.JSON(http.StatusOK, response)
}

func (u *StudentHandler) GetStudentDetailsPagination(c *gin.Context) {

	limit:=c.DefaultQuery("limit", "10")
	offset:=c.DefaultQuery("offset","1")

	studentData, err := u.StudentUsecase.GetStudentDetailsPagination(offset, limit)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't fetch students details", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := responsemodels.Responses(http.StatusOK, "students details fetched succesfully", studentData, nil)
	c.JSON(http.StatusOK, response)
}

func (u *StudentHandler) DeleteStudentDetails(c *gin.Context) {
	var id requestmodels.IdReciever
	if err := c.ShouldBind(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := u.StudentUsecase.DeleteStudentById(&id)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't delete Student", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "student deleted succesfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (u *StudentHandler) EditStudentDetails(c *gin.Context) {
	var studentData requestmodels.EditStudent
	if err := c.ShouldBind(&studentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationRes, err := u.StudentUsecase.EditStudentDetails(&studentData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't edit Student Details", validationRes, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "student details edited succesfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (u *StudentHandler) SearchByNameRollNo(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		response := responsemodels.Responses(http.StatusBadRequest, "can't find Student Details", nil, "no query parameters found")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	studData, err := u.StudentUsecase.SearchByNameRollNo(&query)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't find Student Details", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "student details found succesfully", studData, nil)
	c.JSON(http.StatusOK, response)
}
