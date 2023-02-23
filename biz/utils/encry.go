package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strconv"
)

func demo() {
	orig := "hello world"
	key := "123456781234567812345678"
	fmt.Println("原文：", orig)

	encryptCode := AesEncrypt(orig, key)
	fmt.Println("密文：", encryptCode)

	decryptCode := AesDecrypt(encryptCode, key)
	fmt.Println("解密结果：", decryptCode)
}

func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	l := len(key)
	if l < 32 {
		for i := 0; i < 32-l; i++ {
			key += "0"
		}
	}
	k := []byte(key)[:32]

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		fmt.Println("NewCipher error", err)
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)

	str := base64.StdEncoding.EncodeToString(cryted)
	res := ""
	for _, char := range str {
		res += strconv.Itoa(int(char))
	}
	if len(res) > 128 {
		res = res[:128]
	}
	return res

}

func AesDecrypt(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	l := len(key)
	if l < 32 {
		for i := 0; i < 32-l; i++ {
			key += "0"
		}
	}
	fmt.Println(key)
	k := []byte(key)[:32]

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

// 补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
