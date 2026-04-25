package dberrors

import (
	"errors"

	"github.com/dbtrnl/test.devices-api/internal/devices/domain"
	"github.com/jackc/pgx/v5/pgconn"
)

func Translate(err error, externalID string) error {
	if dbErr, ok := errors.AsType[*pgconn.PgError](err); ok {
		switch dbErr.Code {
		case "P1001":
			return domain.NewErrDeviceInUse(externalID)
		}
	}

	return err
}
