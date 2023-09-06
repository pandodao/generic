# generic
This package implements some generic functions

## Usage

```go
import "github.com/pandodao/generic"
```

## Functions

### Must

**Must** ensures that the given error is nil. If not, it panics.

```go
id,err := uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
if err != nil {
    panic(err)
}
```

to 

```go
id := generic.Must(uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
```

### Try

**Try** executes the given function and ignore the error.

```go
d,_ := decimal.NewFromString("123.456")
```

to 

```go
d := generic.Try(decimal.NewFromString("123.456"))
```

### Ptr

**Ptr** returns a pointer to the given value.

```go
s := &"foo"
// compilation error:
// invalid operation: cannot take address of "foo" (untyped string constant)
```

to 

```go
s := generic.Ptr("foo")
```
