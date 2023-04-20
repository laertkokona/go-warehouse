package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestCopyNonEmptyFields tests the CopyNonEmptyFields function
func TestCopyNonEmptyFields(t *testing.T) {
	type A struct {
		A string
		B string
		C []string
	}
	type B struct {
		A string
		B string
		C []string
	}
	a := A{"a", "b", []string{"c", "d"}}
	b := B{"", "", []string{}}
	CopyNonEmptyFields(&b, &a)
	assert.Equal(t, b.A, "a", "Expected b.A to be a, got", b.A)
	assert.Equal(t, b.B, "b", "Expected b.B to be b, got", b.B)
	assert.Equal(t, b.C, []string{"c", "d"}, "Expected b.C to be c, got", b.C)
}
