package variableswithinitializers

import (
	"regexp"
	"testing"
)

// TestVariableWithInitializers must return a struct.
func TestVariablesWithInitializers(t *testing.T) {
	i := 1
	j := 2
	c := true
	python := false
	java := "no!"

	wantJava := regexp.MustCompile(java)

	v := VariablesWithInitializers()

	if i != v.i || j != v.j || c != v.c || python != v.python || !wantJava.MatchString(v.java) {
		t.Fatalf(`VariablesWithInitializers() = %d, %d, %t, %t, %q, want match for %d, %d, %t, %t, %q`, v.i, v.j, v.c, v.python, wantJava, i, j, c, python, java)
	}

}
