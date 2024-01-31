package requestmodels

import "mime/multipart"

type NewStudent struct {
	Id          uint                  `form:"id"`
	Name        string                `form:"name" validate:"required,gte=2"`
	ROllNo      uint                  `form:"roll_no" validate:"required"`
	Age         uint                  `form:"age" validate:"required"`
	Department  string                `form:"department" validate:"required,gte=3"`
	PhoneNumber string                `form:"phone" validate:"required,e164"`
	ImageFile   *multipart.FileHeader `form:"image" binding:"required"`

	ImageUrl string
}
