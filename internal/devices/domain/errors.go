package domain

import (
	"errors"
	"fmt"
)

type ErrorCode string

const (
	ErrDeviceNotFoundCode  ErrorCode = "err_device_not_found"
	ErrDeviceActiveCode    ErrorCode = "err_device_active"
	ErrDeviceExistsDeleted ErrorCode = "err_device_exists_deleted"
)

type AppError struct {
	Code    ErrorCode
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewDeviceNotFoundError(id string) error {
	return &AppError{
		Code:    ErrDeviceNotFoundCode,
		Message: fmt.Sprintf("device with id %s not found", id),
	}
}

func NewDeviceActiveError(id string) error {
	return &AppError{
		Code:    ErrDeviceActiveCode,
		Message: fmt.Sprintf("device %s is active", id),
	}
}

func NewErrDeviceAExistsDeletedError(name, brand, uuid string) error {
	return &AppError{
		Code:    ErrDeviceExistsDeleted,
		Message: fmt.Sprintf("device %s brand %s uuid %s already exists and it's soft deleted", name, brand, uuid),
	}
}

func IsErrorCode(err error, code ErrorCode) bool {
	if appErr, ok := errors.AsType[*AppError](err); ok {
		return appErr.Code == code
	}
	return false
}
