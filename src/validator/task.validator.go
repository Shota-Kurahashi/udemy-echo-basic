package validator

import (
	"ge-rest-api/src/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task, validation.Field(
		&task.Title,
		validation.Required.Error("Title is required"),
		validation.Length(3, 50).Error("Title must be between 3 and 50 characters"),
	))
}
