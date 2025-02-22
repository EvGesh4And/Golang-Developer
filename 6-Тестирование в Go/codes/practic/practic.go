package practic

import "fmt"

func FuncGetSq(ch int) (zn int, err error) {
	if ch >= 0 {
		zn = ch * ch
	} else {
		err = fmt.Errorf("mrack: %v", ch)
	}
	return
}
