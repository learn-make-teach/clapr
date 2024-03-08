package clapr_test

import (
	"flag"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/learn-make-teach/clapr"
)

type testStruct struct {
	StringVal   string
	BoolVal     bool
	IntVal      int
	DurVal      time.Duration
	UintVal     uint
	StrSliceVal []string
}

func TestNotStruct(t *testing.T) {
	s := "test"
	if clapr.Parse(&s) == nil {
		t.Errorf("only struct should be parseable")
	}
}

func TestDefaultValues(t *testing.T) {
	defer clearFlags()
	conf := testStruct{
		StringVal: "default",
		DurVal:    99 * time.Millisecond,
	}

	err := clapr.Parse(&conf)
	if err != nil {
		t.Errorf("error while parsing: %v", err)
	}
	if conf.BoolVal != false {
		t.Errorf("bool value should still be false")
	}
	if conf.StringVal != "default" {
		t.Errorf("string value should still be default but is %s", conf.StringVal)
	}
	if conf.DurVal != 99*time.Millisecond {
		t.Errorf("duration value should still be default but is %s", conf.DurVal)
	}
}

func TestNewValues(t *testing.T) {
	defer clearFlags()
	conf := testStruct{
		StringVal: "default",
	}

	os.Args = []string{os.Args[0], "-bool-val", "-string-val", "newVal", "--dur-val=10ms", "-str-slice-val", "a,b,c"}
	err := clapr.Parse(&conf)
	if err != nil {
		t.Errorf("error while parsing: %v", err)
	}
	if conf.BoolVal != true {
		t.Errorf("bool value should still be true")
	}
	if conf.StringVal != "newVal" {
		t.Errorf("string value should be newVal but is %s", conf.StringVal)
	}
	if conf.DurVal != 10*time.Millisecond {
		t.Errorf("duration value should be 10ms but is %s", conf.DurVal)
	}
	if !reflect.DeepEqual([]string{"a", "b", "c"}, conf.StrSliceVal) {
		t.Errorf("string slice value should be [a,b,c] but is %v", conf.StrSliceVal)
	}
}

func clearFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}
