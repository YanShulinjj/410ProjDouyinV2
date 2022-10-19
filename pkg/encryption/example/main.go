/* ----------------------------------
*  @author suyame 2022-10-13 16:23:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package main

import (
	"410proj/pkg/encryption"
	"fmt"
)

func main() {
	s := "1234"
	fmt.Println(encryption.Md5ByString(s))
}
