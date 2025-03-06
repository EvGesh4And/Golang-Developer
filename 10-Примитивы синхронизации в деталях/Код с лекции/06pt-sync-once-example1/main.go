package main

import (
	"sync"
)

func main() {
	l := &List{}
	go l.PushFront(1)
	go l.PushFront(2)
	go l.PushFront(3)
	go l.PushFront(4)
	go l.PushFront(5)

}

type List struct {
	once sync.Once
}

func (l *List) PushFront(v interface{}) {
	l.once.Do(l.init)
	//<--- тут и дальше какая-то полезная работа

}
func (l *List) init() {
	func() {
		//что то делаем
		//какая-то одноразовая инициализация
	}()
}
