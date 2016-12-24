iszero [![Build Status](https://secure.travis-ci.org/stephanos/iszero.png)](https://travis-ci.org/stephanos/iszero) [![Coverage Status](https://coveralls.io/repos/stephanos/iszero/badge.png)](https://coveralls.io/r/stephanos/iszero) [![GoDoc](https://camo.githubusercontent.com/6bae67c5189d085c05271a127da5a4bbb1e8eb2c/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f736d61727479737472656574732f676f636f6e7665793f7374617475732e706e67)](http://godoc.org/github.com/stephanos/iszero)
=========

This Go package checks if a value is equal to its type's zero value.


## Example
```go
iszero.Value("") // true
iszero.Value(42) // false

iszero.Value(time.Time{}) // true
iszero.Value(time.Now())  // false

iszero.Value(reflect.ValueOf("")) // true
iszero.Value(reflect.ValueOf(42)) // false
```

## Install
```bash
go get github.com/stephanos/iszero
```

## Documentation
[godoc.org](http://godoc.org/github.com/stephanos/iszero)

## License
MIT (see LICENSE).
