English | [CN中文](README_ZH.md)

### 一、Swift Demangler in Go

#### 1、What is it?

This is a Swift Fuction (:neutral_face: em, yes, just Swift Function right now) Demangler. You can simply import this package and demangle your mangled swift symbols:rocket:

### 二、Detail

#### 1、Installation

- Go >= 1.5 

#### 2、Usage
- install this package

```shell
go get -u github.com/Swift-Demangler-in-Go/swift-demangle
```

- import this package

```go
import swiftDemangle "github.com/Swift-Demangler-in-Go/swift-demangle"
```
- Usage Example

```go
package main

import (
	"fmt"

	swiftDemangle "github.com/Swift-Demangler-in-Go/swift-demangle"
)

func main() {
	symbol, err := swiftDemangle.ToString("_TFC4Pack5class4FuncFSSSb")
	if err != nil {
		panic(err)
	}
	fmt.Println(symbol)
}
```

output:

```
Pack.class.Func(Swift.String) -> Swift.Bool
```

- Make sure that your input is legal, or it would possibly fall into a dead circle.


