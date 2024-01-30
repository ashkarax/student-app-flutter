package interfaceRepository

import requestmodels "github.com/ashkarax/student_data_managing/internal/models/request_models"

type IstudentRepo interface{
	AddStudent(*requestmodels.NewStudent) error
}