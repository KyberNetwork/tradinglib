package httpclient

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	True  = "true"
	False = "false"
)

type Query url.Values

func NewQuery(keyValues ...interface{}) Query {
	q := make(Query)
	q.Pair(keyValues...)
	return q
}

func (q Query) Pair(keyValues ...interface{}) Query {
	for i := 0; i+1 < len(keyValues); {
		key := keyValues[i]
		value := keyValues[i+1]
		i += 2
		q.SetString(fmt.Sprint(key), fmt.Sprint(value))
	}
	return q
}

func (q Query) SetString(key, value string) Query {
	q.setString(key, value, nil)
	return q
}

func stringsContains(ss []string, token string) bool {
	for _, s := range ss {
		if s == token {
			return true
		}
	}
	return false
}

func isOmitEmpty(options []string) bool {
	return stringsContains(options, "omitempty")
}

func (q Query) url() url.Values {
	return (url.Values)(q)
}

func (q Query) setString(key, value string, options []string) Query {
	if value == "" && isOmitEmpty(options) {
		return q
	}
	q.url().Set(key, value)
	return q
}

func (q Query) setInt64(key string, value int64, options []string) Query {
	if value == 0 && isOmitEmpty(options) {
		return q
	}
	q.SetString(key, strconv.FormatInt(value, 10)) // nolint: gomnd
	return q
}

func (q Query) Int64(key string, value int64) Query {
	q.setInt64(key, value, nil)
	return q
}

func (q Query) setUint64(key string, value uint64, options []string) Query {
	if value == 0 && isOmitEmpty(options) {
		return q
	}
	q.SetString(key, strconv.FormatUint(value, 10)) // nolint: gomnd
	return q
}

func (q Query) Uint64(key string, value uint64) Query {
	q.setUint64(key, value, nil)
	return q
}

func (q Query) setFloat(key string, value float64, options []string) Query {
	if value == 0 && isOmitEmpty(options) {
		return q
	}
	q.SetString(key, strconv.FormatFloat(value, 'f', -1, 64)) // nolint: gomnd
	return q
}

func (q Query) Float(key string, value float64) Query {
	q.setFloat(key, value, nil)
	return q
}

func (q Query) Unix(key string, value time.Time) Query {
	q.Int64(key, value.Unix())
	return q
}

func (q Query) UnixMillis(key string, value time.Time) Query {
	q.Int64(key, value.UnixMilli())
	return q
}

func (q Query) Struct(object interface{}) Query {
	if object == nil {
		return q
	}
	q.setWithStruct(object)
	return q
}

func (q Query) setTime(key string, value time.Time, options []string) Query {
	if len(options) == 0 {
		q.setString(key, value.String(), options)
	}
	switch options[0] {
	case "unix":
		q.Int64(key, value.Unix())
	case "unixMilli":
		q.Int64(key, value.UnixMilli())
	case "unixMicro":
		q.Int64(key, value.UnixMicro())
	case "unixNano":
		q.Int64(key, value.UnixNano())
	default:
		q.SetString(key, value.Format(options[0]))
	}
	return q
}

func (q Query) Bool(key string, value bool) Query {
	if value {
		q.SetString(key, True)
	} else {
		q.SetString(key, False)
	}
	return q
}

func (q Query) String() string {
	return url.Values(q).Encode()
}

var (
	stringer = reflect.TypeOf((*fmt.Stringer)(nil)).Elem() // nolint: gochecknoglobals
	timeType = reflect.TypeOf(time.Time{})                 // nolint: gochecknoglobals
)

func unwrapValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Pointer {
		return v.Elem()
	}
	return v
}

func (q Query) setWithStruct(object interface{}) { // nolint: funlen,cyclop
	oValue := unwrapValue(reflect.ValueOf(object))
	oType := oValue.Type()
	for i := 0; i < oValue.NumField(); i++ {
		structField := oType.Field(i)
		fieldValue := oValue.Field(i)
		fieldValueType := fieldValue.Type()

		if !structField.IsExported() {
			continue
		}
		rawValue := fieldValue.Interface()
		keyName, option := q.tagNameAndOption(structField)

		if fieldValue.Kind() == reflect.Pointer {
			if fieldValue.IsNil() {
				q.setString(keyName, "", option)
				continue
			}
			rawValue = fieldValue.Elem().Interface()
			fieldValueType = fieldValue.Elem().Type()
		}
		if fieldValueType == timeType {
			q.setTime(keyName, rawValue.(time.Time), option) // nolint: forcetypeassert
			continue
		}
		if fieldValueType.Implements(stringer) {
			q.setString(keyName, rawValue.(fmt.Stringer).String(), option) // nolint: forcetypeassert
			continue
		}
		if isEmptyValue(fieldValue) && isOmitEmpty(option) {
			continue
		}
		q.setString(keyName, fmt.Sprint(rawValue), option)
	}
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() { // nolint: exhaustive
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	default:
		return false
	}
}

func (q Query) tagNameAndOption(tt reflect.StructField) (string, []string) {
	var (
		keyName string
		option  string
	)
	keyName, option, _ = strings.Cut(tt.Tag.Get("url"), ",")
	if keyName == "" {
		keyName, option, _ = strings.Cut(tt.Tag.Get("json"), ",")
	}
	if keyName == "" {
		keyName = tt.Name
	}
	return keyName, strings.Split(option, ",")
}
