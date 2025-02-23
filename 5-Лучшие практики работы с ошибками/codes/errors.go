package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func ErrorsPack() {
	err := errors.New("error")
	err1 := errors.Wrap(err, "open failed")
	err2 := errors.Wrap(err1, "read config failed")
	fmt.Println(err2) // read config failed: open failed: error
	// fmt.Printf("%+v\n", err2)        // Напечатает stacktrace.
	print(err == errors.Cause(err2), "\n") // true
}
