package requestmodels

type IdReciever struct {
	Id uint `json:"id" validate:"required"`
}
