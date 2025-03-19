package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		// тело - начало
		defer func() { fmt.Println(i) }()
		i = 10
		// тело - конец
	}
}

// аналог
// важно!
// 1. в теле цикла своя переменная
// 2. значения ее передаются в глобальную i
// 3. defer также с изюминкой, так как i можем менять в теле цикла

func main() {
	{
		i := 0
	suda:
		if i < 3 {
			i = func() int {
				i := i
				// тело - начало
				defer func() { fmt.Println(i) }()
				i = 10
				// тело - конец
				return i
			}()
		} else {
			goto tuda
		}
		i++
		goto suda
	tuda:
	}
}
