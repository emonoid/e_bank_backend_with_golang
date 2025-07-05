package api

import (
	"github.com/emonoid/islami_bank_go_backend/utils"
	"github.com/go-playground/validator/v10"
)

var validateCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok:= fieldLevel.Field().Interface().(string); ok {
		return utils.IsCurrencySupported(currency)
	}
	return false
}