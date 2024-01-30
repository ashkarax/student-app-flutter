package responsemodels

type StudentRes struct {
	Name        string `json:"name,omitempty" `
	RollNo      string `json:"roll_no,omitempty" `
	Age         string `json:"age,omitempty" `
	Department  string `json:"department,omitempty" `
	PhoneNumber string `json:"phone,omitempty" `
	ImageFile   string `json:"image,omitempty" `
}
