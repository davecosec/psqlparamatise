package psqlparamatise

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Query struct {
	Columns    string
	Parameters string
	Values     []any
}

// RetrieveFields add tags 'column_name' to your struct and it will build the params required for saving
func RetrieveFields(s interface{}) (*Query, error) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	query := Query{
		Columns:    "",
		Parameters: "",
		Values:     []any{},
	}

	columns := []string{}
	parameters := []string{}

	var count int = 0
	for i := 0; i < t.NumField(); i++ {

		value := v.Field(i).Interface()

		switch value.(type) {
		case string:
			if len(value.(string)) == 0 {
				continue
			}
			break
		case bool:
			break
		case int:
			break
		case float32:
			break
		default:
			return nil, errors.New("type is unknown when retrieving fields")
		}

		tag := t.Field(i).Tag.Get("column_name")
		if len(tag) == 0 {
			continue
		}

		count++

		columns = append(columns, t.Field(i).Tag.Get("column_name"))
		parameters = append(parameters, fmt.Sprintf("$%d", count))
		query.Values = append(query.Values, value)
	}

	query.Columns = strings.Join(columns[:], ",")
	query.Parameters = strings.Join(parameters[:], ",")

	return &query, nil
}
