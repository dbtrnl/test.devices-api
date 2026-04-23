package domain

import "fmt"

type ErrorCode string

const (
    ErrDeviceNotFoundCode ErrorCode = "err_device_not_found"
    ErrDeviceActiveCode   ErrorCode = "err_device_active"
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