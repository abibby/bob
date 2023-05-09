package pkg_test

import (
	"fmt"
	"testing"

	"github.com/abibby/bob/bob-cli/pkg"
)

func TestPkg(t *testing.T) {
	p := pkg.New()
	p.Add("fmt.Print", "test - %s\n", "foo")
	p.Add("github.com/abibby/test.Foo", 1, true, 5.5, "a").WithError().ReturnCount(2)
	fmt.Print(p.ToGo())
	t.Fail()
}
