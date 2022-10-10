### 一、Swift Demangler in Go

#### 1、What is it?

​	这是一个Swift 函数的Demangle工具（:neutral_face: em...是的，现在只能Demangle函数）。你可以通过引入本Package来快速的解析Mangle后的Swift符号🚀


### 二、Detail

#### 1、条件

- Go >= 1.5

#### 2、用法

- 下载此Package

```shell
go get -u github.com/Swift-Demangler-in-Go/swift-demangle
```

- 在代码中import

```go
import swiftDemangle "github.com/Swift-Demangler-in-Go/swift-demangle"
```

- 使用示例

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

输出结果：

```shell
Pack.class.Func(Swift.String) -> Swift.Bool
```

- 请尽可能确保输入的参数合法，否则可能导致死循环
