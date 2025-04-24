package config

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

type CValidator struct {
	validator *validator.Validate
	trans     ut.Translator
}

func NewValidator() *CValidator {
	v := validator.New()
	v.RegisterValidation("dateTimeFormat", ValidateDateTime)
	v.RegisterValidation("dateFormat", ValidateDate)
	v.RegisterValidation("penaltyStatus", ValidatePenaltyStatus)
	trans := NewTranslation(v)
	return &CValidator{
		validator: v,
		trans:     trans,
	}
}

func NewTranslation(v *validator.Validate) ut.Translator {
	id := id.New()
	uni := ut.New(id, id)
	trans, _ := uni.GetTranslator("id")
	id_translations.RegisterDefaultTranslations(v, trans)
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return fieldNameMap[name]
	})
	return trans
}

func ValidateDateTime(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	_, err := time.Parse(time.DateTime, dateStr)
	return err == nil
}
func ValidateDate(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	_, err := time.Parse(time.DateOnly, dateStr)
	return err == nil
}
func ValidatePenaltyStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	if status == "" {
		return true
	}
	values := strings.Split(status, ",")
	allowedStatuses := map[string]bool{
		"1": true,
		"2": true,
		"3": true,
	}
	for _, val := range values {
		val = strings.TrimSpace(val)
		if !allowedStatuses[val] {
			return false
		}
	}
	return true
}

func (cv *CValidator) Validate(payload interface{}) error {
	var res []string

	errs := cv.validator.Struct(payload)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			res = append(res, err.Translate(cv.trans))
		}
	}
	if len(res) > 0 {
		return errors.New(strings.Join(res, ", "))
	}
	return nil
}

var fieldNameMap = map[string]string{
	"user_name": "Nama Pengguna",
	"name":      "nama",
	"badge":     "No. Badge",
	"role_type": "Jenis Pengguna",
}
