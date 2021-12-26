# Printable
[![Go Reference](https://pkg.go.dev/badge/github.com/sohomdatta1/printable.svg)](https://pkg.go.dev/github.com/sohomdatta1/printable)
[![Go Version](https://img.shields.io/badge/go-1.17.5-blue.svg)](https://golang.org/doc/install)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

Go library to check if a byte is a printable character.

## Install

```
go get github.com/sohomdatta1/printable
```

## Usage

### Check if a byte is a printable character

```go
printable.IsPrintable(0x41) // true (0x41 -> 'A')

printable.IsPrintable(0x00) // false
```

### Check if a byte is printable and does not mess up the terminal

```go
printable.IsPrintableAndLayoutSafe('\n') // false

printable.IsPrintableAndLayoutSafe('\t') // false
```

### Iterate over all printable characters

```go
for i, c := range printable.AllPrintable {
    // do something with all printable characters
    if strings.IndexByte(printable.LayoutChanging, c) != -1 {
        // do something with layout changing characters
        continue
    }
    // do something with all printable non layout changing characters
}
```
