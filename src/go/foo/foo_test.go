package foo

import (
	"testing"
)

func TestFoo(t *testing.T) {
	if Foo() != nil {
		t.Fail()
	}
}
