package clapr

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type intVar struct {
	innerVal reflect.Value
}

func (v *intVar) String() string {
	if !v.innerVal.IsValid() || v.innerVal.IsZero() {
		return ""
	}
	return fmt.Sprintf("%d", v.innerVal.Int())
}

func (v *intVar) Set(s string) error {
	i, err := strconv.ParseInt(s, 10, 0)
	v.innerVal.SetInt(i)
	return err
}

type uintVar struct {
	innerVal reflect.Value
}

func (v *uintVar) String() string {
	if !v.innerVal.IsValid() || v.innerVal.IsZero() {
		return ""
	}
	return fmt.Sprintf("%d", v.innerVal.Uint())
}

func (v *uintVar) Set(s string) error {
	i, err := strconv.ParseUint(s, 10, 0)
	v.innerVal.SetUint(i)
	return err
}

type stringVar struct {
	innerVal reflect.Value
}

func (v *stringVar) String() string {
	return fmt.Sprintf(`"%s"`, v.innerVal.String())
}

func (v *stringVar) Set(s string) error {
	v.innerVal.SetString(s)
	return nil
}

type boolVar struct {
	innerVal reflect.Value
}

func (v *boolVar) String() string {
	if !v.innerVal.IsValid() || v.innerVal.IsZero() {
		return "false"
	}
	return fmt.Sprintf(`%v`, v.innerVal.Bool())
}

func (v *boolVar) Set(s string) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	v.innerVal.SetBool(b)
	return nil
}

func (b *boolVar) IsBoolFlag() bool { return true }

type durationVar struct {
	innerVal reflect.Value
}

func (v *durationVar) String() string {
	if !v.innerVal.IsValid() || v.innerVal.IsZero() {
		return ""
	}
	return fmt.Sprintf("%v", v.innerVal.Interface())
}

func (v *durationVar) Set(s string) error {
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	v.innerVal.Set(reflect.ValueOf(d))
	return nil
}

type stringSliceVar struct {
	innerVal reflect.Value
}

func (v *stringSliceVar) String() string {
	if !v.innerVal.IsValid() || v.innerVal.IsZero() {
		return "[]"
	}
	return fmt.Sprintf("%v", v.innerVal.Interface())
}

func (v *stringSliceVar) Set(s string) error {
	s = strings.TrimSpace(s)
	elems := strings.Split(s, ",")
	v.innerVal.Set(reflect.ValueOf(elems))
	return nil
}

type intSliceVar struct {
	innerVal reflect.Value
}

func (v *intSliceVar) String() string {
	if !v.innerVal.IsValid() || v.innerVal.IsZero() {
		return "[]"
	}
	return fmt.Sprintf("%v", v.innerVal.Interface())
}

func (v *intSliceVar) Set(s string) error {
	s = strings.TrimSpace(s)
	elems := strings.Split(s, ",")
	var intElems []int
	var errs error
	for _, e := range elems {
		i, err := strconv.ParseInt(e, 10, 0)
		if err != nil {
			errors.Join(errs, err)
		}
		intElems = append(intElems, int(i))
	}
	v.innerVal.Set(reflect.ValueOf(intElems))
	return errs
}
