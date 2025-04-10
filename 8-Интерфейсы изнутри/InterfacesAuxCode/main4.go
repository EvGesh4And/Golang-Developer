package main

type ImpossibleI interface {
	Method() string
	Method(a int) string
}
