package repository

import (

	interfaceRepository "github.com/ashkarax/student_data_managing/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/student_data_managing/internal/models/request_models"
	"gorm.io/gorm"
)

type studentRepository struct {
	DB *gorm.DB
}

func NewStudentRepository(DB *gorm.DB) interfaceRepository.IstudentRepo {
	return &studentRepository{DB: DB}
}

func (d *studentRepository) AddStudent(studentData *requestmodels.NewStudent) error {

	query:="INSERT INTO students ( name,roll_no,age,department,phone_number,image_url) VALUES(?,?,?,?,?,?)"
	err :=d.DB.Exec(query,studentData.Name,studentData.ROllNo,studentData.Age,studentData.Department,studentData.PhoneNumber,studentData.ImageUrl).Error
	if err != nil {
		return err
	}
	return nil
}
