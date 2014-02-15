# XXTEA for Golang

## Introduction

XXTEA is a fast and secure encryption algorithm. This is a XXTEA library for Golang.

It is different from the original XXTEA encryption algorithm. It encrypts and decrypts []byte instead of []uint32, and the key is also []byte.

## Installation

```bash
go get github.com/xxtea/xxtea-go/xxtea
```

## Usage

```go

package main

import (
    "fmt"
    "github.com/xxtea/xxtea-go/xxtea"
)

func main() {
    str := "Hello World! 你好，中国！"
    key := "1234567890"
    encrypt_data := xxtea.Encrypt([]byte(str), []byte(key))
    decrypt_data := string(xxtea.Decrypt(encrypt_data, []byte(key)))
    if str == decrypt_data {
        fmt.Println("success!")
    } else {
        fmt.Println("fail!")
    }
}
```