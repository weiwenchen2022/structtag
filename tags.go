// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package structtag defines some helper types to easy parsing struct field tag.
package structtag

import (
	"reflect"
	"strings"
)

// A StructTag represents the tag string in a struct field.
type StructTag reflect.StructTag

// Lookup returns the tag structure associated with key in the struct field tag string.
// If there is no such key in the tag, Lookup returns the empty value and false.
func (st StructTag) Lookup(key string) (Tag, bool) {
	tag, ok := (reflect.StructTag)(st).Lookup(key)
	if !ok {
		return Tag{}, false
	}
	name, opts := parseTag(tag)
	return Tag{Key: key, Name: name, Options: opts}, true
}

// Tag represents a single struct's field tag associated with key in the tag string.
type Tag struct {
	// Key is the tag key, such as json, xml, etc...
	// i.e: `json:"field,omitempty,string"`. Here key is: "json"
	Key string

	// Name is a first part of the value.
	// i.e: `json:"field,omitempty,string"`. Here name is: "field"
	Name string

	// Options is a remaining part of the value.
	// i.e: `json:"field,omitempty,string"`. Here options is: "omitempty,string"
	Options TagOptions
}

func (tag Tag) String() string {
	return tag.Key + `:"` + tag.Name + "," + string(tag.Options) + `"`
}

// TagOptions is the string following a comma in a struct field's
// tag, or the empty string. It does not include the leading comma.
type TagOptions string

// parseTag splits a struct field's tag into its name and
// comma-separated options.
func parseTag(tag string) (string, TagOptions) {
	tag, opt, _ := strings.Cut(tag, ",")
	return tag, TagOptions(opt)
}

// Contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o TagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var name string
		name, s, _ = strings.Cut(s, ",")
		if optionName == name {
			return true
		}
	}
	return false
}
