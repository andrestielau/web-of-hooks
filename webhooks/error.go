package webhooks

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"woh/package/utils"
	webhooksv1 "woh/webhooks/adapt/http/v1"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/samber/lo"
)

func Error(w http.ResponseWriter, err error) bool {
	var e utils.Error

	if err == nil {
		return false
	}

	if IsPLSQLColumnNullableViolation(err) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	} else if IsPLSQLForeignKeyConstraintViolation(err) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	} else if IsPLSQLDuplicateKey(err) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	} else if IsInvalidUUID(err) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	} else if errors.Is(err, pgx.ErrNoRows) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	} else if errors.As(err, &e) {
		http.Error(w, e.Reason, e.Code)
		return true
	}
	var plSQLError *pgconn.PgError
	if errors.As(err, &plSQLError) {
		log.Println(plSQLError.Code)
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
	return true
}

func Errors(w http.ResponseWriter, err error) ([]webhooksv1.ErrorItem, bool) {
	var e utils.Errors
	if !errors.As(err, &e) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, true
	}
	return lo.Map(e, func(e utils.Error, _ int) webhooksv1.ErrorItem {
		return webhooksv1.ErrorItem{
			Code:   e.Code,
			Index:  e.Index,
			Reason: e.Reason,
		}
	}), false
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
