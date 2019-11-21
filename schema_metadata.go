package sheru

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"go/ast"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

type schemaMetadata struct {
	tableName    string
	structFields []*structField
	modelType    reflect.Type
}

// StructField model field's struct definition
type structField struct {
	dbName string
	name   string
	//Names           []string
	//IsPrimaryKey    bool
	//IsNormal        bool
	//IsIgnored       bool
	//IsScanner       bool
	//HasDefaultValue bool
	tag reflect.StructTag
	//TagSettings     map[string]string
	reflect_struct reflect.StructField
	//IsForeignKey    bool
	//Relationship    *Relationship

	//tagSettingsLock sync.RWMutex
}

func (metadata *schemaMetadata) quotedTableName() string {
	newStrs := []string{}

	for _, str := range strings.Split(metadata.tableName, ".") {
		newStrs = append(newStrs, fmt.Sprintf(`"%s"`, str))
	}
	return strings.Join(newStrs, ".")
}

type tabler interface {
	TableName() string
}

func schemaDataFor(model interface{}) schemaMetadata {
	modelType := reflect.TypeOf(model)
	fields := structFields(modelType)
	metadata := schemaMetadata{
		modelType:    modelType,
		structFields: fields,
	}

	if tabler, ok := model.(tabler); ok {
		metadata.tableName = tabler.TableName()
	} else {
		metadata.tableName = strings.ToLower(inflection.Plural(metadata.modelType.Name()))
	}

	return metadata
}

func structFields(modelType reflect.Type) []*structField {
	fields := []*structField{}

	for i := 0; i < modelType.NumField(); i++ {
		if fieldStruct := modelType.Field(i); ast.IsExported(fieldStruct.Name) {

			field := &structField{
				dbName:         strcase.ToSnake(fieldStruct.Name),
				reflect_struct: fieldStruct,
				name:           fieldStruct.Name,
				tag:            fieldStruct.Tag,
			}

			indirectType := fieldStruct.Type
			for indirectType.Kind() == reflect.Ptr {
				indirectType = indirectType.Elem()
			}

			fieldValue := reflect.New(indirectType)
			if _, isScanner := fieldValue.Interface().(sql.Scanner); isScanner {
				// TODO handler sql.Scanner
			} else if fieldStruct.Anonymous {
				for _, subField := range structFields(fieldValue.Elem().Type()) {
					fields = append(fields, subField)
				}
				continue
			}

			fields = append(fields, field)
		}
	}

	return fields
}
