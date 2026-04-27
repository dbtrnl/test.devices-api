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
	ErrDeviceInUse         ErrorCode = "err_device_in_use"
	ErrDeviceDeleted       ErrorCode = "err_device_deleted"
	ErrDeviceAlreadyExists ErrorCode = "err_device_already_exists"
)

type AppError struct {
	Code    ErrorCode
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewErrDeviceNotFound(id string) error {
	return &AppError{
		Code:    ErrDeviceNotFoundCode,
		Message: fmt.Sprintf("device with id %s not found", id),
	}
}

func NewErrDeviceActive(id string) error {
	return &AppError{
		Code:    ErrDeviceActiveCode,
		Message: fmt.Sprintf("device %s is active", id),
	}
}

func NewErrDeviceExistsDeleted(name, brand, uuid string) error {
	return &AppError{
		Code:    ErrDeviceExistsDeleted,
		Message: fmt.Sprintf("device %s brand %s uuid %s already exists and it's soft deleted", name, brand, uuid),
	}
}

func NewErrDeviceInUse(uuid string) error {
	return &AppError{
		Code:    ErrDeviceInUse,
		Message: fmt.Sprintf("device uuid %s is in-use", uuid),
	}
}

func NewErrDeviceDeleted(uuid string) error {
	return &AppError{
		Code:    ErrDeviceDeleted,
		Message: fmt.Sprintf("device uuid %s is soft-deleted", uuid),
	}
}

func NewErrDeviceAlreadyExists() error {
	return &AppError{
		Code:    ErrDeviceAlreadyExists,
		Message: "device with this name and brand already exists",
	}
}

func IsErrorCode(err error, code ErrorCode) bool {
	if appErr, ok := errors.AsType[*AppError](err); ok {
		return appErr.Code == code
	}
	return false
}
