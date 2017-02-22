package ioext_test

import (
	"testing"
	"strings"
	"github.com/lummie/tmf/ioext"
	"github.com/lummie/assert"
)

func TestCount_singleLine(t *testing.T) {
	s := `aaa`
	r := strings.NewReader(s)
	i, err := ioext.CountLines(r, []byte{'\n'})
	if err != nil {
		t.Error(err)
	}
	assert.Assert(t, i, assert.EqualInt,1, "Expected 1 row")
}


func TestCount_multiLine(t *testing.T) {
	s := `aaa
	bbb`
	r := strings.NewReader(s)
	i, err := ioext.CountLines(r, []byte{'\n'})
	if err != nil {
		t.Error(err)
	}
	assert.Assert(t, i, assert.EqualInt,2, "Expected 1 row")
}


func TestCount_multiLineEmptyLast(t *testing.T) {
	s := `aaa
	bbb
	`
	r := strings.NewReader(s)
	i, err := ioext.CountLines(r, []byte{'\n'})
	if err != nil {
		t.Error(err)
	}
	assert.Assert(t, i, assert.EqualInt,3, "Expected 1 row")
}


func TestCount_Empty(t *testing.T) {
	s := ""
	r := strings.NewReader(s)
	i, err := ioext.CountLines(r, []byte{'\n'})
	if err != nil {
		t.Error(err)
	}
	assert.Assert(t, i, assert.EqualInt,0, "Expected no rows")
}


