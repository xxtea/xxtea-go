# XXTEA 加密算法的 Golang 实现

<a href="https://github.com/xxtea/">
    <img src="https://avatars1.githubusercontent.com/u/6683159?v=3&s=86" alt="XXTEA logo" title="XXTEA" align="right" />
</a>

## 简介

XXTEA 是一个快速安全的加密算法。本项目是 XXTEA 加密算法的 Golang 实现。

它不同于原始的 XXTEA 加密算法。它是针对 []byte 类型数据进行加密的，而不是针对 uint32 数组。同样，密钥也是 []byte 类型。

## 安装

```sh
go get github.com/xxtea/xxtea-go/xxtea
```

## 使用

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
