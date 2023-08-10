package repositories

import (
	"errors"
	"strconv"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

func GeneratePlaceholders(n int) string {
	if n <= 0 {
		return ""
	}

	var builder strings.Builder
	sep := ", "
	for i := 1; i <= n; i++ {
		if i == n {
			sep = ""
		}
		builder.WriteString("$" + strconv.Itoa(i) + sep)
	}

	return builder.String()
}

func GenerateUpdatePlaceholders(fields []string) string {
	var builder strings.Builder
	sep := ", "

	totalField := len(fields)
	for i, field := range fields {
		if i == totalField-1 {
			sep = ""
		}

		builder.WriteString(field + " = $" + strconv.Itoa(i+1) + sep)
	}

	return builder.String()
}

func transformError(err error) error {
	pqErr := err.(*pgconn.PgError)
	switch pqErr.Code {
	case pgerrcode.UniqueViolation:
		return errors.New("unique violation")
	default:
		return err
	}
}
