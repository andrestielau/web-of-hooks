package webhooks

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"woh/package/utils"
	webhooksgrpcv1 "woh/webhooks/adapt/grpc/v1"
	webhookshttpv1 "woh/webhooks/adapt/http/v1"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/samber/lo"
)

func convertError(err error) *utils.Error {
	if err == nil {
		return nil
	}
	if IsPLSQLColumnNullableViolation(err) {
		return utils.NewError(http.StatusBadRequest, "", err.Error())
	} else if IsPLSQLForeignKeyConstraintViolation(err) {
		return utils.NewError(http.StatusBadRequest, "", err.Error())
	} else if IsPLSQLDuplicateKey(err) {
		return utils.NewError(http.StatusBadRequest, "", err.Error())
	} else if IsInvalidUUID(err) {
		return utils.NewError(http.StatusBadRequest, "", err.Error())
	} else if errors.Is(err, pgx.ErrNoRows) {
		return utils.NewError(http.StatusBadRequest, "", err.Error())
	}
	var plSQLError *pgconn.PgError
	if errors.As(err, &plSQLError) {
		log.Println(plSQLError.Code)
	}

	return utils.NewError(http.StatusInternalServerError, "", err.Error())
}

func HttpError(w http.ResponseWriter, err error) bool {
	var e utils.Error

	error := convertError(err)
	if error == nil {
		return false
	}

	if errors.As(err, &e) {
		http.Error(w, e.Reason, e.Code)
		return true
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
	return true
}

func HttpErrors(w http.ResponseWriter, err error) ([]webhookshttpv1.ErrorItem, bool) {
	var e utils.Errors
	if !errors.As(err, &e) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, true
	}
	return lo.Map(e, func(e utils.Error, _ int) webhookshttpv1.ErrorItem {
		return webhookshttpv1.ErrorItem{
			Code:   e.Code,
			Index:  e.Index,
			Reason: e.Reason,
		}
	}), false
}

func GrpcError(err error) *webhooksgrpcv1.Error {
	var e utils.Error

	error := convertError(err)
	if error == nil {
		return nil
	}

	if errors.As(err, &e) {
		return &webhooksgrpcv1.Error{
			Code:   int32(e.Code),
			Index:  e.Index,
			Reason: e.Reason,
		}
	}

	return &webhooksgrpcv1.Error{
		Code:   http.StatusInternalServerError,
		Index:  "",
		Reason: err.Error(),
	}
}

func GrpcErrors(err error) []*webhooksgrpcv1.Error {
	errors := make([]*webhooksgrpcv1.Error, 1)
	errors[0] = GrpcError(err)
	return errors
}

// IsPLSQLColumnNullableViolation determines if the provided error is a PLSQL non-null constraint violation
func IsPLSQLColumnNullableViolation(originalErr error) bool {
	var plSQLError *pgconn.PgError
	return errors.As(originalErr, &plSQLError) && plSQLError.Code == "23502"
}

// IsPLQLForeignKeyConstraintViolation determines if the provided error is a PLSQL foreign key constraint violation error
func IsPLSQLForeignKeyConstraintViolation(originalErr error) bool {
	var plSQLError *pgconn.PgError
	return errors.As(originalErr, &plSQLError) && plSQLError.Code == "23503"
}

// IsPLSQLDuplicateKey determines if the provided error is a PLSQL duplicate key error
func IsPLSQLDuplicateKey(originalErr error) bool {
	var plSQLError *pgconn.PgError
	return errors.As(originalErr, &plSQLError) && plSQLError.Code == pgerrcode.UniqueViolation
}

// IsInvalidUUID determines if the provided error cannot parse UUID
func IsInvalidUUID(originalErr error) bool {
	return strings.Contains(originalErr.Error(), "cannot parse UUID")
}
