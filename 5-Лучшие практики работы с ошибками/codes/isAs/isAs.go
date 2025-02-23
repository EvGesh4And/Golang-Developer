package isas

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func IsAs() {
	baseErr := &MyError{0x08006, "db connection error"}
	err := fmt.Errorf("read user: %w", baseErr)

	// Проверьте, имеет ли ошибка тип MyError
	if errors.Is(err, baseErr) {
		fmt.Println("Error is of type DbError")
	}

	// Попробуйте извлечь базовое значение MyError
	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Printf("Extracted DbError: %v\n", myErr.Code)
	}
}
