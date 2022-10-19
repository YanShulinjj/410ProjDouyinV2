/* ----------------------------------
*  @author suyame 2022-10-18 14:44:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package fmtx

import (
	"fmt"
	"math"
)

// 将数字转换成字符串
// 高于1000 -> %d.%dk  eg. 1245 -> 1.2k
// 高于100000 -> 10w+

func ItoA(num uint64) string {
	if num > 100000 {
		return "10w+"
	}
	if num > 1000 {
		return fmt.Sprintf("%.1fk", math.Trunc(float64(num)/100)/10)
	}
	return fmt.Sprintf("%d", num)
}
