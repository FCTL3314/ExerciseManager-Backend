package errormapper

import (
	"ExerciseManager/internal/domain"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type PostgresErrUniqueViolationMapper struct{}

func (m *PostgresErrUniqueViolationMapper) MapError(err error) (error, bool) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		field := getMappedConstraintFieldName(pgErr.ConstraintName)
		return &domain.ErrObjectUniqueConstraint{Fields: []string{field}}, true
	}
	return nil, false
}

func getMappedConstraintFieldName(constraintName string) string {
	constraintToFieldMap := map[string]string{
		"users_username_key": "username",
	}
	if field, ok := constraintToFieldMap[constraintName]; ok {
		return field
	}
	return "unknown"
}

func BuildPostgresErrorsMapperChain() *MapperChain {
	mc := NewChain()
	mc.registerMapper(&PostgresErrUniqueViolationMapper{})
	return mc
}
