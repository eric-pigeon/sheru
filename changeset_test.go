package sheru

import (
	"github.com/matryer/is"
	"testing"
)

func TestUniqueConstraint(t *testing.T) {
	is := is.New(t)

	post := Post{}
	changeset := Change(post)

	changeset.UniqueConstraint("title")

	expected := []Constraint{Constraint{
		Constraint:   "posts_title_index",
		ErrorMessage: "has already been taken",
		ErrorType:    "unique",
		Field:        "title",
	}}
	is.Equal(expected, changeset.Constraints)
}

func TestToInsertSql(t *testing.T) {
	is := is.New(t)

	post := Post{}
	changeset := Change(post)

	expected := `INSERT INTO "posts" ("id","title") VALUES ()`
	adapter := postgres{}
	sql, _ := changeset.toInsertSql(adapter)
	is.Equal(expected, sql)
}
