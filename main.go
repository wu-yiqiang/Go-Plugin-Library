// main
package main

import (
	"Go-Plugin-Library/encrypt"
	"fmt"
)

func main() {
	var PwdKey = []byte("KDJDKJJFJ*LKJSD)")
	var originData = []byte("lalala")
	fmt.Println("--------------")
	fmt.Println(encrypt.AesEcrypt(originData, PwdKey))
	fmt.Println("--------------")
}
