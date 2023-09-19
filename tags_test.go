// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package structtag

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetStructTag(t *testing.T) {
	t.Parallel()

	type S struct {
		F string `species:"field,foobar,foo"`
	}

	st := (StructTag)(reflect.TypeOf(S{}).Field(0).Tag)
	tag, ok := st.Lookup("species")
	if !ok {
		t.Fatal(`not found tag "species"`)
	}

	if tag.Name != "field" {
		t.Fatalf(`Name = %q, want "field"`, tag.Name)
	}

	opts := []struct {
		opt  string
		want bool
	}{
		{"foobar", true},
		{"foo", true},
		{"bar", false},
	}
	for _, tt := range opts {
		if tt.want != tag.Options.Contains(tt.opt) {
			t.Errorf("Contains(%q) = %t", tt.opt, !tt.want)
		}
	}

	got := tag.String()
	want := string(reflect.TypeOf(S{}).Field(0).Tag)
	if want != got {
		t.Fatalf(`tag = %q, want %s`, got, want)
	}
}

func TestTag_string(t *testing.T) {
	t.Parallel()

	tests := []struct {
		i any
	}{
		{struct {
			F string `species:"-"`
		}{}},
		{struct {
			F string `species:"-,"`
		}{}},
	}

	for i, tt := range tests {
		f := reflect.TypeOf(tt.i).Field(0)
		tag, ok := (StructTag)(f.Tag).Lookup("species")
		if !ok {
			t.Errorf("#%d not found tag %s", i, string(f.Tag))
			continue
		}
		if !cmp.Equal(string(f.Tag), tag.String()) {
			t.Errorf("#%d mismatch: %s", i, cmp.Diff(string(f.Tag), tag.String()))
		}
	}
}
