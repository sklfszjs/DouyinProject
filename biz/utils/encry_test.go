package utils

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	orig := "test1"
	key := "password1胡凡"
	fmt.Println("原文：", orig)

	encryptCode := AesEncrypt(orig, key)
	fmt.Println("密文：", encryptCode)

	decryptCode := AesDecrypt(encryptCode, key)
	fmt.Println("解密结果：", decryptCode)
}
