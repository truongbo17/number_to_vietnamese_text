# Number to Vietnamese Words Converter

## 📌 Introduction

This is a Go library that helps convert numbers to Vietnamese words accurately and quickly.

## 🚀 How to use

### Install

```sh
    go get github.com/truongbo17/number_to_vietnamese_text
```

### Example

```go
package main

import (
	"github.com/truongbo17/number_to_vietnamese_text"
)

func main() {
	// With struct options
	number := 123456789.003
	num := NumberToVietnameseWord{
		Number: number,
		Options: NumberToVietnameseWordOption{
			UcFirst:     true,
			FormatMoney: VNDFull,
		},
	}
	num.Convert()
	// => Một trăm hai mươi ba triệu bốn trăm năm mươi sáu nghìn bảy trăm tám mươi chín phẩy không không ba việt nam đồng

	// Or simple
	ToVietnameseWords(123456789)
	// => Một trăm hai mươi ba triệu bốn trăm năm mươi sáu nghìn bảy trăm tám mươi chín phẩy không không ba

}
```

## ✅ Run Unit Test

```sh
    go test ./...
```

## 📄 License

MIT
