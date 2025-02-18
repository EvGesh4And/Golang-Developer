# Особенности языка и типовые ошибки

## Преподаватель  
**Рубаха Юрий**  

## Содержание занятия  
- Затенения и ошибки связанные с областью видимости
- Замыкания и ошибки связанные с именованными значениями
- Устройство слайсов
- Мапы
- Ошибки при использовании слайсов и мап

---

## Области видимости и блоки

```go
var a = 1                       // <- уровень пакета

func main() {
    fmt.Println("1: ", a)
    a := 2                      // <-- уровень блока функции
    fmt.Println("2: ", a)
    {
        a := 3                  // <-- уровень пустого блока
        fmt.Println("3: ", a)
    }
    fmt.Println("4: ", a)       // <-- a = 2
    f()
}
func f() {
    fmt.Println("5: ", a)       // <-- a = 1
}

```

### Неявные блоки: if, for, switch, case, select

#### **if**
```go
func classicIf() {
	if x := 10; x > 5 { // x создаётся в неявном блоке
		fmt.Println("x больше 5")
	} // Здесь x уничтожается
}

func withNeyavnyBlockIf() {
	{ // Неявный блок if
		x := 10    // x создаётся
		if x > 5 { // Используется x
			fmt.Println("x больше 5")
		}
	} // Здесь x уничтожается
}
```

#### **for**

```go
func classicFor() {
	for i := 0; i < 3; i++ { // i создаётся в неявном блоке
		x := i * 2
		fmt.Println(x)
	} // Здесь i уничтожается
}

func withNeyavnyBlockFor() {
	{ // Неявный блок for
		i := 0 // i создаётся в неявном блоке
	tuta: // метка для goto
		if i < 3 {
			// начало тела функции
			x := i * 2
			fmt.Println(x)
			// конец тела функции
			i++
			goto tuta // возвращаемся на метку
		}
	} // Здесь i уничтожается
}
```

#### **switch**

```go
func classicSwitch() {
	switch x := 2; x { // x создаётся в неявном блоке
	case 1:
		fmt.Println("Один")
	case 2:
		fmt.Println("Два")
	} // Здесь x уничтожается
}

func withNeyavnyBlockSwitch() {
	{ // Неявный блок for
		x := 2 // x создаётся в неявном блоке
		if x == 1 {
			fmt.Println("Один")
		}
		if x == 2 {
			fmt.Println("Два")
		}
	} // Здесь x уничтожается
}
```

#### **select**

```go
select {
case msg := <-ch1:
    fmt.Println(msg) // msg доступен только в этом блоке
case msg := <-ch2:
    fmt.Println(msg) // msg объявляется заново
}
```

#### **Примеры**

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

```go
if v, err := doSmth(); err != nil {
    fmt.Println(err)
} else {
    process(v)
}
```

```go
switch i := 2; i * 4 {
case 8:
    j := 0
    fmt.Println(i, j)
default:
    // "j" is undefined here
    fmt.Println(i)
}
```

#### Вопрос: сколько раз объявлен x?

```go
package main

import "fmt"

func f(x int) {                 // <-- 1-ое объявление 
    for x := 0; x < 10; x++ {   // <-- 2-ое объявление
        fmt.Println(x)
    }
}

var x int                       // <-- 3-ое объявление

func main() {
    var x = 200                 // <-- 4-ое объявление
    f(x)
}
```

> **4-е раза**

### Опасное затенение

```go
func main() {
    data, err := callServer()
    if err != nil {
        fmt.Println(err)
        return
    }
    defer func() {
        if err != nil {
        fmt.Println(err)
        }
    }()
    if err := saveToDB(data); err != nil {  // <-- тут затенение
        fmt.Println(err)
        return
    }
    return
}

func callServer() (int, error) {return 0, nil}
func saveToDB(a int) error {return fmt.Errorf("save error")}
```

## Функции: именованные возвращаемые значения

```go
func sum(a, b int) (s int) {
    s = a + b
    return
}
```

### Опасный `defer`

```go
func main() {
    if err := DoDBRequest(); err != nil {
        fmt.Println(err)
    }
}

func DoDBRequest() (err error) {
    defer func() {
        if err = close(); err != nil {
            return
        }
    }()
    
    err = request()
    return
}

func request() error {return fmt.Errorf("request error")}
func close() error {return nil}
```

### Замыкания


```go
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
func main() {
    nextInt := intSeq()
    fmt.Println(nextInt()) // <-- 1
    fmt.Println(nextInt()) // <-- 2
    fmt.Println(nextInt()) // <-- 3
    newInts := intSeq()
    fmt.Println(newInts()) // <-- 1
}
```

### Замыкания: middleware

```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/hello", timed(hello))
    http.ListenAndServe(":3000", nil)
}

func timed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    f(w, r)
    end := time.Now()
    fmt.Println("The request took", end.Sub(start))
    }
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "<h1>Hello!</h1>")
}
```

### Опасное замыкание

```go
func main() {
    s := "hello"
    defer fmt.Println(s)    // <-- вызывается вторым (замыкания нет)
    defer func() {          // <-- вызывается первым (есть замыкание)
        fmt.Println(s)
    }()
    s = "world"
}
```
## Слайсы

### Как они устроены?

```go
// runtime/slice.go
type slice struct {
    array unsafe.Pointer
    len int
    cap int
}
```

```go
l := len(s) // len — вернуть длину слайса
c := cap(s) // cap — вернуть емкость слайса
```

```go
s := make([]int, 3, 10) // s == {0, 0, 0}
```


[Хороший источник про slice](https://go.dev/blog/slices-intro)

