package errormapper

import (
	"ExerciseManager/internal/domain"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"strings"
)

type PostgresErrUniqueViolationMapper struct{}

func (m *PostgresErrUniqueViolationMapper) MapError(err error) (error, bool) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		constraints := strings.Split(pgErr.ConstraintName, ",")

		var fields []string

		for _, constraint := range constraints {
			field := getMappedConstraintFieldName(constraint)
			fields = append(fields, field)
		}

		return &domain.ErrObjectUniqueConstraint{Fields: fields}, true
	}
	return nil, false
}

func getMappedConstraintFieldName(constraintName string) string {
	constraintFieldMap := map[string]string{
		"unique_username": "username",
	}
	if field, ok := constraintFieldMap[constraintName]; ok {
		return field
	}
	return "unknown"
}

func BuildPostgresErrorsMapperChain() *MapperChain {
	mc := NewChain()
	mc.registerMapper(&PostgresErrUniqueViolationMapper{})
	return mc
}
