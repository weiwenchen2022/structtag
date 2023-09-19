package structtag_test

import (
	"fmt"
	"reflect"

	"github.com/weiwenchen2022/structtag"
)

func Example() {
	type S struct {
		F string `json:"field,omitempty,string" xml:"foo"`
	}

	// get struct field tag
	tag := structtag.StructTag(reflect.TypeOf(S{}).Field(0).Tag)

	// get a single tag
	jsonTag, ok := tag.Lookup("json")
	if !ok {
		panic("not found json tag")
	}

	fmt.Println(jsonTag.Key)
	fmt.Println(jsonTag.Name)
	fmt.Println(jsonTag.Options.Contains("omitempty"))
	fmt.Println(jsonTag)

	// Output:
	// json
	// field
	// true
	// json:"field,omitempty,string"
}
