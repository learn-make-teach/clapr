package clapr

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"
)

var (
	ErrNotStruct = errors.New("only structs are supported")
)

func Parse(s any) error {
	var err error
	st := reflect.TypeOf(s).Elem()
	if st.Kind() != reflect.Struct {
		return ErrNotStruct
	}
	v := reflect.ValueOf(s).Elem()
	for i := range st.NumField() {
		short, long, desc := parseTags(st.Field(i))
		var flagVar flag.Value
		switch v.Field(i).Interface().(type) {
		case time.Duration:
			flagVar = &durationVar{v.Field(i)}
		default:
			switch st.Field(i).Type.Kind() {
			case reflect.Slice, reflect.Array:
				switch st.Field(i).Type.Elem().Kind() {
				case reflect.String:
					flagVar = &stringSliceVar{v.Field(i)}
				case reflect.Int:
					flagVar = &intSliceVar{v.Field(i)}
				default:
					err = errors.Join(err, fmt.Errorf("type %v not supported", st.Field(i).Type))
				}
			case reflect.String:
				flagVar = &stringVar{v.Field(i)}
			case reflect.Bool:
				flagVar = &boolVar{v.Field(i)}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				flagVar = &intVar{v.Field(i)}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				flagVar = &uintVar{v.Field(i)}
			default:
				err = errors.Join(err, fmt.Errorf("type %v not supported", st.Field(i).Type))
			}
		}
		if flagVar != nil {
			if short != "" {
				flag.Var(flagVar, short, desc)
			}
			if long != "" {
				flag.Var(flagVar, long, desc)
			}

		}
	}
	flag.Parse()
	return err
}

func parseTags(field reflect.StructField) (short, long, desc string) {
	tag := field.Tag.Get("clapr")
	tagElements := strings.Split(tag, ",")
	for _, s := range tagElements {
		kv := strings.Split(s, "=")
		if len(kv) < 2 {
			continue
		}
		switch kv[0] {
		case "short":
			short = kv[1]
		case "long":
			long = kv[1]
		case "desc":
			desc = kv[1]
		}
	}
	if short == "" && long == "" {
		long = toSnakeCase(field.Name)
	}
	return
}

func toSnakeCase(s string) string {
	buf := bytes.Buffer{}
	buf.WriteByte(s[0])
	for _, c := range s[1:] {
		if unicode.IsUpper(c) {
			buf.WriteByte('-')
		}
		buf.WriteRune(c)
	}
	return strings.ToLower(buf.String())
}
