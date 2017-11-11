# jconf

The golang-package for working with configuration json files

## Example

```go
// ...

config, err := NewConfig("./config.json")
if err != nil {
	panic(err)
}

name, err := config.Get("name")
if err != nil {
	panic(err)
}

// ...
```