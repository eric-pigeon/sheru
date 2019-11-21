package sheru

type validator interface {
	valid() bool
}

type Validation struct {
	Field        string
	ErrorMessage string
	validator    validator
	// some interface
}

func (changeset *Changeset) ValidateRequired(fields ...string) {
	changeset.Required = append(changeset.Required, fields...)
}

type IntegerValidations struct {
	LessThan             int
	GreaterThan          int
	LessThanOrEqualTo    int
	GreaterThanOrEqualTo int
	EqualTo              int
	NotEqualTo           int
}

func (validations IntegerValidations) valid() bool {
	// TODO
	return false
}

func (changeset *Changeset) ValidateInteger(field string, integerValidations IntegerValidations) {
	validation := Validation{
		Field:     field,
		validator: integerValidations,
	}
	changeset.Validations = append(changeset.Validations, validation)
}
