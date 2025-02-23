package main

import (
	"error/isas"
	"errors"
	"fmt"
	"io"
)

type MyError struct {
	s string
}

func (m *MyError) Error() string {
	return m.s
}

func New(ss string) error {
	return &MyError{s: ss}
}

func main() {
	fmt.Println(New("ds"))

	whoami := "error"
	err := fmt.Errorf("Im an %s \n", whoami)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(errors.New("EOF") == io.EOF)

	ErrorsPack()

	isas.IsAs()
}
