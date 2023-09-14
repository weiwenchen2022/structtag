# structtag [![PkgGoDev](https://pkg.go.dev/badge/github.com/weiwenchen2022/structtag)](https://pkg.go.dev/github.com/weiwenchen2022/structtag)

Package structtag defines some helper types to easy parsing struct field tag.

# Install

```bash
go get github.com/weiwenchen2022/structtag
```

# Example

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/weiwenchen2022/structtag"
)

func main() {
	type T struct {
		F string `json:"field,omitempty,string" xml:"foo"`
	}

	// get struct field tag
	tag := structtag.StructTag(reflect.TypeOf(T{}).Field(0).Tag)

	// get a single tag
	jsonTag, ok := tag.Lookup("json")
	if err != nil {
		panic("not found json tag")
	}

	fmt.Println(jsonTag.Key)
	fmt.Println(jsonTag.Name)
	fmt.Println(jsonTag.Options.Contains("omitempty"))

	fmt.Printf("%#v\n", jsonTag)
	fmt.Println(jsonTag)

	// Output:
	// json
	// field
	// true
	// structtag.Tag{Key:"json", Name:"field", Options:"omitempty,string"}
	// json:"field,omitempty,string"
}
```
