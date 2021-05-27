package env

import (
	"os"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	p := &struct {
		First  string `env:"first"`
		Second string `env:"second"`
		Third  string `env:"third"`
	}{}
	os.Setenv("zero", "0")
	os.Setenv("first", "1")
	os.Setenv("second", "2")
	Unmarshal(p)
	if !(p.First == "1" && p.Second == "2" && p.Third == "") {
		t.FailNow()
	}
}
