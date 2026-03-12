package resource

type resourceDTO struct {
	ID           uint    `json:"id"`
	Version      int     `json:"version"`
	TextField    *string `json:"textField" validate:"required,notblank"`
	NumberField  *int    `json:"numberField" validate:"required"`
	BooleanField *bool   `json:"booleanField" validate:"required"`
}
