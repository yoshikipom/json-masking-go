# json-masking-go

## Usage Examples

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

### Remove all email address with config file
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

### Remove email address in root level only with config file
```
$ cat dev/data.json | go run main.go mask --config ./dev/config-example.yml      
{
  "email": "*",
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
  ],
  "id": 1,
  "name": "Tom"
}
```

### Remove all email address without config file
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
