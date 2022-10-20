# Go Learning

Why not.

## Standard Library

Go has a large standard library ready to use that we can use for anything from network connectivity to math, crypto, image processing, filesystem access, and more.

<https://pkg.go.dev/std>

## Building

```sh
# windows
go build

# linux
npx cross-env GOOS=linux go build main.go

# mac
npx cross-env GOOS=darwin go build main.go
```

## Variable Types

Integers (int, int8, int16, int32, rune, int64, uint, uintptr, uint8, uint16, uint64)

Floats (float32, float64), useful to represent decimals

Complex types (complex64, complex128), useful in math

Byte (byte), represents a single ASCII character

Strings (string), a set of bytes

Booleans (bool), either true or false

