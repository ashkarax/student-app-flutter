package requestmodels

type EditStudent struct {
	Id          uint   `json:"id"`
	Name        string `json:"name" validate:"required,gte=2"`
	RollNo      uint   `json:"roll_no" validate:"required"`
	Age         uint   `json:"age" validate:"required"`
	Department  string `json:"department" validate:"required,gte=3"`
	PhoneNumber string `json:"phone" validate:"required,e164"`
}
