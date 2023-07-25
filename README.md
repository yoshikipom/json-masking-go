# About
[json-masking-go](https://pkg.go.dev/github.com/yoshikipom/json-masking-go) is a tool to hide specific fields (for example personal information) in JSON for logging, communicaiton etc.

- CLI support
  - input JSON -> from stdin
  - config (filtering target etc) -> from command option or configuration file
  - output JSON -> to stdout
- Go library support
  - input JSON -> `[]byte`
  - config (filtering target etc) -> `MaskingInput` (struct)
  - output JSON -> `[]byte`
- Regex support for filtering target
- Filtering for nested fileds are supported

If you have any requests, please let me know in Issues!

## Usage (CLI)

### Input JSON for the examples
```
$  cat dev/data.json 
{
    "email": "test@mail",
    "id": 1,
    "name": "Tom",
    "friends": [
        {
            "email": "test@mail",
            "id": 2,
            "name": "Alice"
        },
        {
            "email": "test@mail",
            "id": 3,
            "name": "Bob"
        }
    ]
}
```

### Mask all email address with config file by using regex
```
$ cat dev/data.json | go run main.go mask --config ./dev/config-example-regex.yml 
{
  "email": "*",
  "friends": [
    {
      "email": "*",
      "id": 2,
      "name": "Alice"
    },
    {
      "email": "*",
      "id": 3,
      "name": "Bob"
    }
  ],
  "id": 1,
  "name": "Tom"
}
```

### Mask all email address without config file
```
$ cat dev/data.json | go run main.go mask -d "email$" --regex --format
{
  "email": "test@mail",
  "friends": [
    {
      "email": "test@mail",
      "id": 2,
      "name": "*"
    },
    {
      "email": "test@mail",
      "id": 3,
      "name": "*"
    }
  ],
  "id": 1,
  "name": "*"
}
```

### Mask all values in friends address without config file
```
$ cat dev/data.json | go run main.go mask --config ./dev/config-example-mask-object.yml
{
  "email": "test@mail",
  "friends": [
    {
      "email": "*",
      "id": "*",
      "name": "*"
    },
    {
      "email": "*",
      "id": "*",
      "name": "*"
    }
  ],
  "id": 1,
  "name": "Tom"
}
```

## Usage (Library)
Program
```go
package main

import (
	"fmt"

	"github.com/yoshikipom/json-masking-go/masking"
)

func main() {
	config := masking.MaskingConfig{
		DeniedKeyList: []string{"name"},
	}
	m := masking.New(&config)

	input := `{"key":"value","name":"John"}`
	output := m.Replace([]byte(input))
	fmt.Println(string(output))
}
```

Execution
```
$ go run main.go
{"key":"value","name":"*"}
```
