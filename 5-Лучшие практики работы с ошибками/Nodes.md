# Лучшие практики работы с ошибками

## Преподаватель  
**Алексей Романовский**  

## Содержание занятия  

- Ошибки:
    - принципы обработки,
    - лучшие практики.
- panic, recover, defer

## Ошибки

- Ошибка - тип, реализующий интерфейс `error`
- Функции возвращают ошибки как обычные значения
- По конвенции, ошибка - последнее возвращаемое функцией значение
- Ошибки обрабатываются проверкой значения (и/или передаются выше через `return`)


```go
type error interface {
    Error() string
}
```

```go
func Marshal(v interface{}) ([]byte, error) {
    e := &encodeState{}
    err := e.marshal(v, encOpts{escapeHTML: true})
    if err != nil {
        return nil, err
    }
    return e.Bytes(), nil
}
```

### errors.go

Пакет `errors` предоставляет функцию `New`, которая создаёт ошибку с заданным текстовым сообщением.

Пример реализации `errors.New`:

```go
package errors

func New(text string) error {
    return &errorString{text}
}

type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```

Пример использования:

```go
err := errors.New("Im an error")
if err != nil {
    fmt.Print(err)
}
```

### fmt.Errorf

Функция `fmt.Errorf` позволяет форматировать сообщение об ошибке, используя параметры, как в `fmt.Sprintf`.

Пример использования:

```go
whoami := "error"
err := fmt.Errorf("Im an %s", whoami)
if err != nil {
    fmt.Print(err)
}
```

## Идиоматичная проверка ошибок

Идиоматичный способ обработки ошибок в Go предполагает немедленный возврат ошибки, если она возникла.

```go
func (router HttpRouter) parse(reader *bufio.Reader) (Request, error) {

    requestText, err := readCRLFLine(reader)
    if err != nil {
        return nil, err
    }

    requestLine, err := parseRequestLine(requestText)
    if err != nil {
        return nil, err
    }

    if request := router.routeRequest(requestLine); request != nil {
        return request, nil
    }

    return nil, requestLine.NotImplemented()
}
```

## Ошибка - это значение

Ошибки в Go являются обычными значениями, что позволяет сохранять их и обрабатывать в удобный момент.

Пример возврата ошибки из функции:

```go
func (s *Scanner) Scan() (token []byte, error) {
    scanner := bufio.NewScanner(input)
    for {
        token, err := scanner.Scan()
        if err != nil {
            return err // Возвращаем ошибку при её возникновении
        }
        // Обрабатываем токен
    }
}
```

Пример сохранения ошибки во внутренней структуре:

В этом примере ошибки не проверяются на каждой итерации цикла. Вместо этого они сохраняются во внутреннем состоянии `scanner`, а затем проверяются после завершения сканирования. Такой подход удобен, когда требуется сначала обработать все данные, а уже потом определить, возникли ли ошибки во время чтения.

```go
scanner := bufio.NewScanner(input)
for scanner.Scan() {
    token := scanner.Text()
    // Обрабатываем токен
}

// Проверяем, была ли ошибка после завершения сканирования
if err := scanner.Err(); err != nil {
    // Обрабатываем ошибку после завершения сканирования
}
```

## Обработка ошибок: sentinel values

В Go ошибки могут быть объявлены как фиксированные значения (sentinel values). Такие ошибки являются частью публичного API и используются для явной проверки.

```go
package io

// ErrShortWrite означает, что операция записи записала меньше байтов, чем ожидалось,
// но при этом не вернула явную ошибку.
var ErrShortWrite = errors.New("short write")

// ErrShortBuffer означает, что для чтения требовался буфер большего размера,
// чем был предоставлен.
var ErrShortBuffer = errors.New("short buffer")
```

Sentinel-ошибки сравниваются напрямую:

```go
if err == io.EOF {
    // Обрабатываем конец файла
}
```

## Сравнение ошибок в Go

В Go ошибки сравниваются по указателям и по значению. Даже если две ошибки содержат одинаковые строки, они могут быть разными объектами с разными указателями, что приведет к результату `false` при прямом сравнении с помощью `==`.

Пример:

`errors.New("EOF") == io.EOF` — это всегда `false`, потому что:

- `errors.New("EOF")` создает новый объект типа `*errorString`, каждый раз при вызове.
- `io.EOF` — это заранее определённая глобальная переменная типа `*errorString`, которая всегда указывает на один и тот же объект в памяти.
- Хотя оба объекта содержат строку `"EOF"`, это разные объекты с разными указателями в памяти, поэтому результат сравнения через `==` будет `false`.

### Как правильно сравнивать ошибки

Чтобы сравнивать ошибки по содержимому, используйте метод `.Error()` для получения строкового значения ошибки. Этот метод возвращает строку, которая хранится внутри ошибки.

Пример:
```go
errors.New("EOF").Error() == io.EOF.Error()
Это сравнение вернёт true, если обе ошибки содержат одинаковую строку (например, "EOF").
```

Функции могут вернуть `io.EOF`, чтобы сигнализировать о том, что поток данных завершён, и это нормально. Это специфическая ошибка, которая широко используется для управления потоками ввода/вывода в Go.

Поскольку `io.EOF` — это глобальная переменная, сравнение с ней будет корректным и точным, если ошибка, возвращаемая функцией, является именно этой ошибкой.

### Итог

Типы одинаковые, но указатели на объекты разные. Это приводит к тому, что прямое сравнение через `==` даёт `false`.
Для корректного сравнения ошибок по содержимому используйте метод `.Error()`, чтобы сравнить строки.


## Проверка ошибок

### Типы

Структура `PathError` записывает ошибку, а также операцию и путь к файлу, которые её вызвали.

```go
// PathError записывает ошибку, операцию и путь к файлу, которые её вызвали.
type PathError struct {
    Op   string // Операция, например, "open", "unlink" и т. д.
    Path string // Путь к файлу, связанный с ошибкой.
    Err  error  // Ошибка, возвращённая системным вызовом.
}
```
Метод `Error()` для структуры `PathError` возвращает строковое представление ошибки, включая операцию, путь и саму ошибку.

```go
func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```
Этот метод создаёт строку, которая содержит информацию о том, какая операция (например, `open` или `unlink`) не удалась на каком пути, а также саму ошибку, которую вернул системный вызов.

Пример:
```go
open /etc/passwx: no such file or directory
```

```go
err := readConfig()
switch err := err.(type) {
    case nil:
        // Вызов функции прошел успешно, ошибки нет
    case *PathError:
        fmt.Println("invalid config path:", err.Path)
    default:
        // Неизвестная ошибка
}
```


### Интерфейсы

#### Интерфейс `Error` в пакете `net`

В пакете `net` расширяется стандартный интерфейс `error`. Интерфейс `Error` добавляет два метода:

```go
package net

type Error interface {
    error          // Встраиваем стандартный интерфейс error.
    Timeout() bool // Возвращает true, если ошибка связана с тайм-аутом.
    Temporary() bool // Возвращает true, если ошибка временная.
}
```

- **`Timeout()`**: Возвращает `true`, если ошибка связана с тайм-аутом (например, запрос в сети не был завершён в положенный срок).
- **`Temporary()`**: Возвращает `true`, если ошибка временная (например, сервер временно недоступен, и можно попробовать повторить операцию позже).

#### Проверка ошибок

Для обработки ошибок и проверки их типа используется утверждение типа. Пример проверки ошибки типа `net.Error`:

```go
if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
    time.Sleep(1e9)  // Ждем 1 секунду (1e9 наносекунд) и продолжаем выполнение.
    continue
}
```

- **`err.(net.Error)`**: Преобразует ошибку `err` в тип `net.Error`. Если преобразование удаётся, переменная `nerr` получает значение типа `net.Error`, а переменная `ok` будет равна `true`.
- **`nerr.Temporary()`**: Проверяет, является ли ошибка временной. Если ошибка временная, выполняется пауза, и программа пытается выполнить операцию снова.

### Завершение программы в случае ошибки

Если ошибка не временная, программа завершает выполнение:

```go
if err != nil {
    log.Fatal(err) // Если ошибка не nil, завершаем программу с выводом ошибки.
}
```

#### Использование пакета `net`

Пакет `net` предоставляет множество полезных функций для работы с сетевыми операциями. Дополнительную информацию о типах и методах пакета можно найти в официальной документации:

[net package documentation](https://golang.org/pkg/net/#pkg-index)

#### Общая идея

- Использование интерфейса `Error` позволяет обрабатывать сетевые ошибки более гибко, проверяя, связана ли ошибка с тайм-аутом или является временной.
- В случае временных ошибок программа может повторить попытку через некоторое время, что полезно для работы с нестабильными сетями или серверами.
- В случае других типов ошибок программа завершает работу с выводом сообщения об ошибке.

### Больше о пользовательских типах ошибок

[CustomErrors](https://gobyexample.com/custom-errors)

[CustomErrors2](https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go)



### Антипаттерны проверки ошибок

```go
if err.Error() == "smth" { // Строковое представление - для людей.
    // ...
}
```

**Проблема**: Сравнение строкового представления ошибки с фиксированным значением не является хорошей практикой, потому что строковое представление может изменяться в разных версиях, а также не позволяет легко обрабатывать различные типы ошибок.

### 2. Пропуск проверки ошибки

```go
func Write(w io.Writer, buf []byte) {
    w.Write(buf) // Забыли проверить ошибку
}
```

**Проблема**: Пропуск проверки ошибок может привести к невидимым багам, когда ошибки, такие как проблемы с записью в файл или сетевое подключение, не обрабатываются должным образом.

### 3. Множественное логирование ошибки

```go
func Write(w io.Writer, buf []byte) error {
    _, err := w.Write(buf)
    if err != nil {
        // Логируем ошибку вероятно несколько раз
        // на разных уровнях абстракции.
        log.Println("unable to write:", err)
        return err
    }
    return nil
}
```

**Проблема**: Логирование одной и той же ошибки на разных уровнях абстракции может привести к избыточным и неинформативным логам. Лучше делать логирование ошибки в одном месте, чтобы избежать дублирования.


## Оборачивание ошибок

Если ошибка не может быть обработана на текущем уровне, и мы хотим сообщить её вызывающему с дополнительной информацией

```go
func ReadAndCalcLen() error {
    from, to := readFromTo()
    a, err := Len(from, to)
    if err != nil {
        return fmt.Errorf("calc len for %i and %i: %w", from, to, err)
    }
}
//Результат: calc len for 2 and 1: from should be less than to
```

### Соглашения об оборачивании ошибок

#### Когда

- Необходимо обернуть, если в функции есть 2 или более мест, возвращающих ошибку.
- Можно вернуть исходную ошибку, если есть только 1 `return`.
- Перед добавлением второго `return`, рекомендуется рефакторинг первого `return`.

#### Как

- Текст при оборачивании описывает место в текущей функции. Например, для функции `openCong(...)`:
    - Да: `fmt.Errorf("open le: %w", err)`  // ✅ правильно 
    - Нет: `fmt.Errorf("startup: %w", err)` // ❌ плохо
- не начинается с заглавной буквы
- не содержит знаков препинания в конце
- разделитель `"слоёв"` - `": "`
- избегайте префиксов `"fail to"` / `"error at"` / `"can not"` в сообщениях обертки.
  - Можно для корневых ошибок и логирования:
    - `log.Warn("fail to read: %v" , err)`


#### Рекомендация

- Сделайте сообщение как можно более уникальным
    (для всего приложения)
- Параметры - в конец

Так будет проще находить код по тексту ошибки.

#### Пример

```go
fail to read form's data: get user: open db connection: network error 0x123
```

### github.com/pkg/errors


Использовалась до того как go научился оборачивать ошибки.
Сейчас - legacy проекты и для стектрейса в ошибке.

```go
_, err := ioutil.ReadAll(r)
if err != nil {
    return errors.Wrap(err, "read failed")
}
```

```go
package main

import "fmt"
import "github.com/pkg/errors"

func main() {
    err := errors.New("error")
    err1 := errors.Wrap(err, "open failed")
    err2 := errors.Wrap(err1, "read config failed")
    fmt.Println(err2)                   // read config failed: open failed: error
    fmt.Printf("%+v\n", err2)           // Напечатает stacktrace.
    print(err == errors.Cause(err2))    // true
}
```

### Обработка обёрнутых ошибок

Ччто, если обёрнуто?

```go
err := doSomethingWithNetwork()

if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
    time.Sleep(1e9)
    continue
}

if err != nil {
    log.Fatal(err)
}
```

`err.(net.Error)` - антипаттерн


### `errors.Is` & `errors.As`

#### `errors.Is`
[Is](https://pkg.go.dev/errors#Is)

Позволяет проверить, является ли ошибка конкретной ожидаемой ошибкой.
```go
var ErrNotFound = errors.New("not found")
func main() {
    err := fmt.Errorf("wrap: %w", ErrNotFound)
    if errors.Is(err, ErrNotFound) {
        fmt.Println("Ошибка - not found")
    }
}
```

#### `errors.As`
[As](https://pkg.go.dev/errors#As)

Позволяет проверить, соответствует ли ошибка заданному типу, и привести её к этому типу.

```go
import "errors"

type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return e.Message
}

func main() {
    baseErr := &MyError{0x08006, "db connection error"}
    err := fmt.Errorf("read user: %w", baseErr)

    var myErr *MyError
    if errors.As(err, &myErr) {
        fmt.Printf("Extracted MyError: %v\n", myErr.Code)
    }
}
```

#### Пример

```go
import "errors"
type MyError struct {
    Code int
    Message string
}

func (e *MyError) Error() string {
    return e.Message
}

func main() {
    baseErr := &MyError{0x08006, "db connection error"}
    err := fmt.Errorf("read user: %w", baseErr)

    // Проверьте, имеет ли ошибка тип MyError
    if errors.Is(err, &MyError{}) {
        fmt.Println("Error is of type DbError")
    }

    // Попробуйте извлечь базовое значение MyError
    var myErr *MyError
    if errors.As(err, &myErr) {
        fmt.Printf("Extracted DbError: %v\n", myErr.Code)
    }
}
```

### Интерфейс Is и As 

[wrap](https://github.com/golang/go/blob/master/src/errors/wrap.go#L58)

Учитываем, когда создаём свои типы ошибок

- **Is** - если надо проверить соответствие ошибки шаблону (тип, значение)
- **As** - если надо ещё и привести ошибку к искомому типу

### Интерфейс `Is` для собственных типов ошибок

Если создаёте кастомные ошибки, можно переопределить метод `Is`, чтобы errors.`Is` работал правильно

```go
type MyError struct {
    Code int
}

func (e *MyError) Error() string {
    return fmt.Sprintf("error with code %d", e.Code)
}

func (e *MyError) Is(target error) bool {
    t, ok := target.(*MyError)
    return ok && t.Code == e.Code
}
```

### Итого

- Проверяйте ошибки.
- Лишний раз не логируйте.
- Проверяйте поведение, а не тип.
- Ошибки - это значения.
- Оборачивайте правильно



## Defer, Panic и Recover

### Defer

`defer` позволяет назначить выполнение вызова функции непосредственно перед выходом из вызывающей функции

```go
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close() // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...)
        if err != nil {
            return "", err // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```

Аргументы отложенного вызова функции вычисляются тогда, когда вычисляется команда `defer`.

```go
func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}
```

Вывод:
```go
0
```

Отложенные вызовы функций выполняются в порядке LIFO: последний отложенный вызов будет вызван первым — после того, как объемлющая функция завершит выполнение.

```go
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}
```

Вывод:
```go
3210
```

Отложенные функции могут читать и устанавливать именованные возвращаемые значения объемлющей функции.

```go
func c() (i int) {
    defer func() { i++ }()
    return 1
}
```

Вывод:
```go
2
```

### Panic и Recover

`Panic` — это встроенная функция, которая останавливает обычный поток управления и начинает паниковать. Когда
функция `F` вызывает `panic`, выполнение `F` останавливается, все отложенные вызовы в `F` выполняются нормально,
затем `F` возвращает управление вызывающей функции. Для вызывающей функции вызов `F` ведёт себя как вызов
`panic`. Процесс продолжается вверх по стеку, пока все функции в текущей го-процедуре не завершат выполнение,
после чего аварийно останавливается программа. Паника может быть вызвана прямым вызовом `panic`, а также
вследствие ошибок времени выполнения, таких как доступ вне границ массива.

`Recover` — это встроенная функция, которая восстанавливает контроль над паникующей го-процедурой. `Recover`
полезна только внутри отложенного вызова функции. Во время нормального выполнения, `recover` возвращает `nil` и не
имеет других эффектов. Если же текущая го-процедура паникует, то вызов `recover` возвращает значение, которое было
передано `panic` и восстанавливает нормальное выполнение.

Паниковать стоит только в случае, если ошибку обработать нельзя, например

```go
var user = os.Getenv("USER")
func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
```

"поймать" панику можно с помощью `recover`: вызов `recover` останавливает выполнение отложенных функций и
возвращает аргумент, переданный `panic`

```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}
func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```


Пример из `encoding/json`:


```go
// jsonError is an error wrapper type for internal use only.
// Panics with errors are wrapped in jsonError so that
// the top-level recover can distinguish intentional panics
// from this package.
type jsonError struct{ error }

func (e *encodeState) marshal(v interface{}, opts encOpts) (err error) {
    defer func() {
        if r := recover(); r != nil {
            if je, ok := r.(jsonError); ok {
                err = je.error
            } else {
                panic(r)
            }
        }
    }()
    e.reflectValue(reflect.ValueOf(v), opts)
    return nil
}
```