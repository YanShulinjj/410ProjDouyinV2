/* ----------------------------------
*  @author suyame 2022-10-18 14:48:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package fmtx

import "testing"

func TestItoA(t *testing.T) {
	examples := []uint64{
		123,
		1234,
		12345,
		65432,
		99999,
		99999999,
	}
	expected := []string{
		"123",
		"1.2k",
		"12.3k",
		"65.4k",
		"99.9k",
		"10w+",
	}

	for i, e := range examples {
		if get := ItoA(e); get != expected[i] {
			t.Errorf("example: %d get: %s != expect: %s, Value Match Err!",
				e, get, expected[i])
		}
	}
}
