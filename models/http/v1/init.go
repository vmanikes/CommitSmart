package v1

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

var (
	validate             *validator.Validate
	validatorInitializer sync.Once
)

func init() {
	validatorInitializer.Do(func() {
		validate = validator.New()
	})
}
