package responsemodels

type StudentRes struct {
	Id          string `json:"id,omitempty" `
	Name        string `json:"name,omitempty" `
	RollNo      string `json:"roll_no,omitempty" `
	Age         string `json:"age,omitempty" `
	Department  string `json:"department,omitempty" `
	PhoneNumber string `json:"phone,omitempty" `
	ImageUrl   string `json:"image,omitempty" `
}
