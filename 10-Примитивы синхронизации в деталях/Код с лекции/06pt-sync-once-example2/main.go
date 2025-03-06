package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	l := GetInstance()
	l.Print()
	l2 := GetInstance()
	l2.Print()
}

type Singleton struct {
	starTime time.Time
}

func (m *Singleton) Print() {
	fmt.Printf("I am alive from %s\n", m.starTime)
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{time.Now()}
	})
	
	return instance
}
