package repository

import (
	"errors"
	"fmt"

	interfaceRepository "github.com/ashkarax/student_data_managing/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/student_data_managing/internal/models/request_models"
	responsemodels "github.com/ashkarax/student_data_managing/internal/models/response_models"
	"gorm.io/gorm"
)

type studentRepository struct {
	DB *gorm.DB
}

func NewStudentRepository(DB *gorm.DB) interfaceRepository.IstudentRepo {
	return &studentRepository{DB: DB}
}

func (d *studentRepository) AddStudent(studentData *requestmodels.NewStudent) error {

	query := "INSERT INTO students ( name,roll_no,age,department,phone_number,image_url) VALUES(?,?,?,?,?,?)"
	err := d.DB.Exec(query, studentData.Name, studentData.ROllNo, studentData.Age, studentData.Department, studentData.PhoneNumber, studentData.ImageUrl).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *studentRepository) GetStudentDetails() (*[]responsemodels.StudentRes, error) {
	var studentDetails []responsemodels.StudentRes

	query := "SELECT * FROM students"
	err := d.DB.Raw(query).Scan(&studentDetails)
	if err.Error != nil {
		return &studentDetails, err.Error
	}

	return &studentDetails, nil
}

func (d *studentRepository) GetStudentDetailsPagination(offset , limit string) (*[]responsemodels.StudentRes, error) {
	var studentDetails []responsemodels.StudentRes

	query := "SELECT * FROM students order by names limit $1 offset $2"
	err := d.DB.Raw(query, limit, offset).Scan(&studentDetails)
	if err.Error != nil {
		return &studentDetails, err.Error
	}

	return &studentDetails, nil
}

func (d *studentRepository) DeleteStudentDetailById(id *requestmodels.IdReciever) error {
	query := "DELETE FROM students WHERE id=?"
	err := d.DB.Exec(query, id.Id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("no student with id=%d found,enter a valid student id", id.Id))
	}
	return nil
}

func (d *studentRepository) EditStudentDetailsById(studentData *requestmodels.NewStudent) error {

	query := "UPDATE students SET name=?,roll_no=?,age=?,department=?,phone_number=?,image_url=?  WHERE id=?"
	err := d.DB.Exec(query, studentData.Name, studentData.ROllNo, studentData.Age, studentData.Department, studentData.PhoneNumber, studentData.ImageUrl, studentData.Id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("no student with id=%d found,enter a valid student id", studentData.Id))
	}
	return nil
}

func (d *studentRepository) SearchByNameRollNo(q *string) (*[]responsemodels.StudentRes, error) {
	var studentDetails []responsemodels.StudentRes

	//Error code
	// query := `SELECT * FROM students WHERE name ILIKE $1 OR rollno = $2`
	// err := d.DB.Raw(query, "%"+*q+"%", *q).Scan(&studentDetails)
	//outputted_error:invalid input syntax for type bigint: "ayoob" (SQLSTATE 22P02)
	// /////////////////////////////////////////////////////////////////////////////////

	query := `SELECT * FROM students WHERE name ILIKE $1 OR roll_no::text ILIKE $2`
	err := d.DB.Raw(query, "%"+*q+"%", "%"+*q+"%").Scan(&studentDetails)
	if err.Error != nil {
		return &studentDetails, err.Error
	}

	return &studentDetails, nil
}
