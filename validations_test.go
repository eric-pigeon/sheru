package sheru

import (
	"github.com/matryer/is"
	"testing"
)

func TestValidateRequired(t *testing.T) {
	is := is.New(t)

	post := Post{}
	changeset := Change(post)

	changeset.ValidateRequired("title")
	expected := []string{"title"}

	is.Equal(expected, changeset.Required)
}

func TestValidateNumber(t *testing.T) {
	is := is.New(t)

	post := Post{}
	changeset := Change(post)

	changeset.ValidateInteger("number", IntegerValidations{EqualTo: 3})

	expected := []Validation{Validation{
		Field:     "number",
		validator: IntegerValidations{EqualTo: 3},
	}}

	is.Equal(expected, changeset.Validations)
}
