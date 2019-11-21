package sheru

import (
	"fmt"
	"strings"
)

type Constraint struct {
	Constraint   string // the database constraint name
	ErrorMessage string
	//Name         string
	ErrorType string
	Field     string
}

type Changeset struct {
	Constraints []Constraint
	Required    []string
	Validations []Validation
	metadata    schemaMetadata
}

func (changeset *Changeset) UniqueConstraint(field string) {
	constraint := Constraint{
		Constraint:   changeset.metadata.tableName + "_" + field + "_index",
		ErrorMessage: "has already been taken",
		ErrorType:    "unique",
		Field:        field,
		//Match: :exact,
		//Type: :unique
	}
	changeset.Constraints = append(changeset.Constraints, constraint)
}

func (changeset *Changeset) columns(adapter adapter) []string {
	var columns []string

	for _, field := range changeset.metadata.structFields {
		columns = append(columns, adapter.quote(field.dbName))
	}

	return columns
}

func (changeset *Changeset) toInsertSql(adapter adapter) (string, []interface{}) {
	var args []interface{}
	return fmt.Sprintf(
		//"INSERT INTO %v (%v)%v VALUES (%v)%v%v",
		"INSERT INTO %v (%v) VALUES",
		changeset.metadata.quotedTableName(),
		strings.Join(changeset.columns(adapter), ","),
	), args
}
