package interfaceRepository

import (
	requestmodels "github.com/ashkarax/student_data_managing/internal/models/request_models"
	responsemodels "github.com/ashkarax/student_data_managing/internal/models/response_models"
)

type IstudentRepo interface{
	AddStudent(*requestmodels.NewStudent) error
	GetStudentDetails() (*[]responsemodels.StudentRes,error)
	DeleteStudentDetailById(*requestmodels.IdReciever) error
	EditStudentDetailsById(*requestmodels.EditStudent) error
	SearchByNameRollNo(*string)(*[]responsemodels.StudentRes,error)
	GetStudentDetailsPagination(string,string) (*[]responsemodels.StudentRes,error)
}