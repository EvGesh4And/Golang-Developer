package main

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(fname string) error {
	var err *os.PathError // nil

	if err == nil {
		log.Println("err is nil")
		return err
	}
	// Do some work...
	return err
}

func main() {
	var EmptyPointer *os.PathError
	var MyNill interface{} = EmptyPointer
	var MyNill2 interface{} = (*os.PathError)(nil)

	if err := ReadFile(""); err != nil {
		log.Printf("ERR: (%T, %v)", err, err)
		if err == MyNill {
			fmt.Println("А тут - равно")
		}

	} else {
		log.Println("OK")
	}

}
