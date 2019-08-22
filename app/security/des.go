package security

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"errors"
	"log"
	"runtime"
)
const (
	ivdes = "lovevolt" //iv 长度必须为8
	key = "Whatdoyoudoifyouarenotin"//Whatdoyoudoifyouarenotinfear 加密key 字节长度必须为24
)

func DesEncrypt(encrStr string)(string,error){
	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil{
		return "",err
	}
	paddingText := pKCS5Padding([]byte(encrStr), block.BlockSize())

	var iv =[]byte(ivdes)
	blockMode := cipher.NewCBCEncrypter(block, iv)

	cipherText := make([]byte,len(paddingText))
	blockMode.CryptBlocks(cipherText,paddingText)
	return  base64.StdEncoding.EncodeToString(cipherText),nil
}

func DesDecrypt(decrStr string) (string,error){
	decrByte ,_ := base64.StdEncoding.DecodeString(decrStr)
	block, err := des.NewTripleDESCipher([]byte(key))
	if err!=nil{
		return "",err
	}

	defer func(){
		if err:=recover();err!=nil{
			switch err.(type){
			case runtime.Error:
				log.Println("runtime error:",err,"Check that the key is correct")
			default:
				log.Println("error:",err)
			}
		}
	}()

	var iv =[]byte(ivdes)
	blockMode := cipher.NewCBCDecrypter(block, iv)

	paddingText := make([]byte,len(decrByte)) //
	blockMode.CryptBlocks(paddingText,decrByte)


	plainText ,err:= pKCS5UnPadding(paddingText)
	if err!=nil{
		return "",err
	}
	return string(plainText),nil
}

func pKCS5Padding(plainText []byte, blockSize int) []byte{
	padding := blockSize - (len(plainText)%blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}

func pKCS5UnPadding(plainText []byte)([]byte,error){
	length := len(plainText)
	number:= int(plainText[length-1])
	if number>=length{
		return nil,errors.New("请检查密钥或iv")
	}
	return plainText[:length-number],nil
}