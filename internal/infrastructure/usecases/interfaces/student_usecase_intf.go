package interfaceUseCase

import (
	requestmodels "github.com/ashkarax/student_data_managing/internal/models/request_models"
	responsemodels "github.com/ashkarax/student_data_managing/internal/models/response_models"
)

type IstudentUseCase interface {
	AddStudent(*requestmodels.NewStudent) (*responsemodels.StudentRes, error)
	GetStudentDetails() (*[]responsemodels.StudentRes, error)
	DeleteStudentById(*requestmodels.IdReciever) error
	EditStudentDetails(*requestmodels.NewStudent) (*responsemodels.StudentRes, error)
}
