package usecases

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	interfaceRepository "github.com/ashkarax/student_data_managing/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/student_data_managing/internal/infrastructure/usecases/interfaces"
	requestmodels "github.com/ashkarax/student_data_managing/internal/models/request_models"
	responsemodels "github.com/ashkarax/student_data_managing/internal/models/response_models"
	"github.com/ashkarax/student_data_managing/pkg/aws"
	"github.com/go-playground/validator/v10"
)

type studentUsecase struct {
	StudentRepo interfaceRepository.IstudentRepo
}

func NewStudentUseCase(studentRepository interfaceRepository.IstudentRepo) interfaceUseCase.IstudentUseCase {
	return &studentUsecase{StudentRepo: studentRepository}
}

func (r *studentUsecase) AddStudent(studentData *requestmodels.NewStudent) (*responsemodels.StudentRes, error) {

	var studentResp responsemodels.StudentRes
	BucketFolder := "student-app-images/"

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(studentData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Name":
					studentResp.Name = "should be a valid Name. "
				case "RollNo":
					studentResp.RollNo = "should be a valid rollno. "
				case "Age":
					studentResp.Age = "should be a valid Age. "
				case "Department":
					studentResp.Department = "should be a valid department "
				case "PhoneNumber":
					studentResp.PhoneNumber = "should include the country code also."
				}
			}
		}
		return &studentResp, errors.New("did't fullfill the validation conditions ")
	}

	if studentData.ImageFile.Size > 1*1024*1024 { // 1 MB limit
		return &studentResp, errors.New("image size exceeds the limit (1MB)")
	}

	// Validate file types
	allowedTypes := map[string]struct{}{
		"image/jpeg": {},
		"image/png":  {},
		"image/gif":  {},
	}

	file, err := studentData.ImageFile.Open()
	if err != nil {
		return &studentResp, err
	}
	defer file.Close()

	// Read the first 512 bytes to determine the content type
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return &studentResp, err
	}

	// Reset the file position after reading
	_, err = file.Seek(0, 0)
	if err != nil {
		return &studentResp, err
	}

	// Get the content type based on the file content
	contentType := http.DetectContentType(buffer)

	// Check if the content type is allowed
	if _, ok := allowedTypes[contentType]; !ok {
		return &studentResp, errors.New("unsupported file type,should be a jpeg,png or gif")
	}

	sess, errInit := aws.AWSSessionInitializer()
	if errInit != nil {
		fmt.Println(errInit)
		return &studentResp, errInit
	}
	imageURL, err := aws.AWSS3ImageUploader(studentData.ImageFile, sess, &BucketFolder)
	if err != nil {
		fmt.Printf("Error uploading image  %v\n", err)
		return &studentResp, err
	}

	studentData.ImageUrl = *imageURL

	fmt.Println(studentData)

	errN := r.StudentRepo.AddStudent(studentData)
	if errN != nil {
		return &studentResp, errN
	}

	return &studentResp, nil
}

func (r *studentUsecase) GetStudentDetails() (*[]responsemodels.StudentRes, error) {
	data, err := r.StudentRepo.GetStudentDetails()
	if err != nil {
		return data, err
	}

	return data, nil

}

func (r *studentUsecase) GetStudentDetailsPagination(offset, limit string) (*[]responsemodels.StudentRes, error) {
	offsetInt, _ := strconv.Atoi(offset)
	limitInt, _ := strconv.Atoi(limit)

	offset = strconv.Itoa((offsetInt - 1) * limitInt)
	data, err := r.StudentRepo.GetStudentDetailsPagination(offset, limit)
	if err != nil {
		return data, err
	}

	return data, nil

}

func (r *studentUsecase) DeleteStudentById(id *requestmodels.IdReciever) error {

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(id)
	if err != nil {
		return err
	}

	errD := r.StudentRepo.DeleteStudentDetailById(id)
	if errD != nil {
		return errD
	}

	return nil

}

func (r *studentUsecase) EditStudentDetails(studentData *requestmodels.EditStudent) (*responsemodels.StudentRes, error) {

	var studentResp responsemodels.StudentRes
	if studentData.Id == 0 {
		return &studentResp, errors.New("enter student id")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(studentData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Name":
					studentResp.Name = "should be a valid Name. "
				case "RollNo":
					studentResp.RollNo = "should be a valid rollno. "
				case "Age":
					studentResp.Age = "should be a valid Age. "
				case "Department":
					studentResp.Department = "should be a valid department "
				case "PhoneNumber":
					studentResp.PhoneNumber = "should include the country code also."

				}
			}
		}
		return &studentResp, errors.New("did't fullfill the validation conditions ")
	}

	errN := r.StudentRepo.EditStudentDetailsById(studentData)
	if errN != nil {
		return &studentResp, errN
	}

	return &studentResp, nil
}

func (r *studentUsecase) SearchByNameRollNo(query *string) (*[]responsemodels.StudentRes, error) {
	data, err := r.StudentRepo.SearchByNameRollNo(query)
	if err != nil {
		return data, err
	}

	return data, nil

}
