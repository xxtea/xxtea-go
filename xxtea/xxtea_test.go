package xxtea

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func Test_XXTEA(t *testing.T) {
	str := "Hello World! 你好，中国！"
	key := "1234567890"
	encrypt_data := Encrypt([]byte(str), []byte(key))
	fmt.Println(base64.StdEncoding.EncodeToString(encrypt_data))
	decrypt_data := string(Decrypt(encrypt_data, []byte(key)))
	if str != decrypt_data {
		t.Error(str)
		t.Error(decrypt_data)
		t.Error("fail!")
	}
}
