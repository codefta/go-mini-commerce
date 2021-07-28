package validator

import (
	"fmt"

	"github.com/fathisiddiqi/go-mini-commerce/models"
	"github.com/go-playground/validator"
)

type Validate struct {
	validate *validator.Validate
}

func New() *Validate {
	validate := validator.New()
	validate.SetTagName("form")

	return &Validate{validate: validate}
}

func (v *Validate) ProductValidator(productForm models.ProductForm) error {
	err := v.validate.Struct(productForm)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}