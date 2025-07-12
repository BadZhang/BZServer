package flag_test

import (
	"github.com/badzhang/bzserver/v2/flag"
	"testing"
)

func TestString(t *testing.T) {
	t.Log(flag.Bool("test.v"))
	t.Log(flag.String("config", "./config"))
}
