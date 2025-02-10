# Типы данных

## Преподаватель  
**Рубаха Юрий**  

## Содержание занятия  
- Основные синтаксические конструкции языка
- Преобразование и присвоение типов
- Указатели
- Слайсы
- Словари
- Строки, руны и функции работы с ними
- Структуры, функции и методы

---

## Элементарные типы данных в Go

- Логические значения: `bool`
- Целые числа: `int`, `uint`, `int8`, `uint8`, `int16`, `uint16`, `int32`, `uint32`, `int64`, `uint64`
- Алиасы к целым числам: `byte` = `uint8`, `rune` = `int32`
- Числа с плавающей точкой: `float32`, `float64`
- Комплексные числа: `complex64`, `complex128`
- Строки: `string`
- Указатели: `*int`, `*string`, `*T` для любого `T`
>**\*Целочисленные значения**: `uintptr` (используется для хранения указателей как чисел)


### Константы

Константы — неизменяемые значения, доступные только во время компиляции

```go
const PI = 3            // принимает подходящий тип
const pi float32 = 3.14 // строгий тип

const (
    TheA = 1
    TheB = 2
)

const (
    X = iota            // 0
    Y                   // 1
    Z                   // 2
)
```

### Объявление переменных в Go

```go
var Storage map[string]string           // zero value
var storage = make(map[string]string)   // автовывод типа

func main() {
    var m int       // zero value
    k := m          // короткое объявление, только внутри функций
    var i int = 10
    n, j := m, i
}
```

### Публичные и приватные идентификаторы

Публичные идентификаторы — те, которые видны за пределами вашего пакета. Публичные идентификаторы начинаются с заглавной буквы `Storage`, `Printf`.

Приватные идентификаторы — начинаются со строчной буквы `i`, `j` и видны только в вашем пакете.

Структуры могут содержать как приватные так и публичные поля.

```go
type User struct {
    Name string         // Будет видно в json.Marshal.
    password string     // Не будет видно.
}
```

### Литералы числовых типов

```go
493         // десятичная система
0755        // восьмеричная система
0xDeadBeaf  // шестнадцатеричная, hex
3.14        // с плавающей точкой

.288
2.e+10

1+1i        // комплексные
```
### Особенности целых чисел в Go

- Есть значение "по умолчанию" — это 0
- Типы `int` и `uint` могут занимать 32 и 64 бита на разных платформах
- Нет автоматического преобразования типов
- `uintptr` — целое число, не указатель

### Преобразование типов

В Go всегда необходимо явное преобразование типов

```go
var i int32 = 42
var j uint32 = i            // ошибка!
var k uint32 = uint32(i)    // верно
var n int64 = i             // ошибка!
var m int64 = int64(i)      // верно
var r rune = i              // верно
```

## Массивы

```go
var arr [256]int                // фиксированная длина
var arr [10][10]string          // может быть многомерным
arr := [...]int{1, 2, 3}        // автоматически подсчет длины
arr := [10]int{1, 2, 3, 4, 5}
```

### Операции

```go
v := arr[1]     // чтение
arr[3] = 1      // запись
len(arr)        // длина массива
arr[2:4]        // получение слайса
```
## Слайсы

Слайсы — это те же "массивы", но переменной длины

Создание слайсов:

```go
var s []int             // неинициализированны слайс, nil
s := []int{}            // с помощью литерала слайса
s := make([]int, 3, 4)  // с помощью функции make, s == {0,0,0}
```

### Операции

```go
v := s[1]               // чтение
s[3] = 1                // запись
len(s)                  // длина слайса
cap(s)                  // емкость слайса
s[2:4]                  // получение подслайса
s = append(s, 1)        // добавляет 1 в конец слайса
s = append(s, 1, 2, 3)  // добавляет 1, 2, 3 в конец слайса
s = append(s, s2...)    // добавляет содержимое слайса s2 в конец s
var s []int             // s == nil
s = append(s, 1)        // s == {1} append умеет работать с nil-слайсами
```

## Строки в Go

Строки в Go - это неизменяемая последовательность байтов ( `byte` = `uint8` ).

В Go строка — это, по сути, срез байтов, доступный только для чтения.

```go
// src/runtime/string.go
type stringStruct struct {
 str unsafe.Pointer
 len int
}
```

### Строковые литералы

```go
s := "hello world"              // в двойных кавычках, на одной строке
s := "hello \n world \u9333"    // c непечатными символами
// если нужно включить в строку кавычки или переносы строки
// - используем обратные кавычки
s := `hello
"cruel"
'world'
`
```

### Операции

```go
s := "hello world"      // создавать
var c byte = s[0]       // получать доступ к байту(!) в строке
var s2 string = s[5:10] // получать подстроку (в байтах!)
s2 := s + " again"      // склеивать
l := len(s)             // узнавать длину в байтах
```

### Руны в Go

Руна в Go - это алиас к int32. Каждая руна представляет собой код символа стандарта Юникод.
Литералы рун выглядят так
```go
var r rune = 'Я'
var r rune = '\n'
var r rune = '本'
var r rune = '\xff'         // последовательность байт
var r rune = '\u12e4'       // unicode code-point
```

```go
s := "hey😉"
rs := []rune([]byte(s))     // cannot convert ([]byte)(s) (type []byte) to type []rune
bs := []byte([]rune(s)])    // cannot convert ([]rune)(s) (type []rune) to type []byte
```

### Строки: итерирование

По байтам

```go
for i := 0; i < len(s); i++ {
    b := s[i]
    // i строго последоваельно
    // b имеет тип byte, uint8
}
```

По рунам

```go
str := "Привет мир!"
for i, v := range str {
    fmt.Printf("%d:%s\n", i, string(v))
}
```

### Функции для работы со строками

Пакет из стандартной библиотеки `strconv`

#### Числовые преобразования

```go
i, err := strconv.Atoi("-42")
s := strconv.Itoa(-42)
```

```go
b, err := strconv.ParseBool("true")
f, err := strconv.ParseFloat("3.1415", 64)
i, err := strconv.ParseInt("-42", 10, 64)
u, err := strconv.ParseUint("42", 10, 64)
```

## Словари (map)

- Отображение `ключ` => `значение`.
- Реализованы как хэш-таблицы.
- Аналогичные типы в других языках: в Python — `dict`, в JavaScript — `Object`, в Java — `HashMap`, в C++ — `unordered_map`.

### Создание

```go
var cache map[string]string     // не-инициализированный словарь, nil
cache := map[string]string{}    // с помощью литерала, len(cache) == 0
cache := map[string]string{     // литерал с первоначальным значением
    "one": "один",
    "two": "два",
    "three": "три",
}
cache := make(map[string]string)        // тоже что и map[string]string{}
cache := make(map[string]string, 100)   // заранее выделить память на 100 ключей
```

### Операции

```go
value := cache[key]         // получение значения,
value, ok := cache[key]     // получить значение, и флаг того что ключ найден
_, ok := cache[key]         // проверить наличие ключа в словаре
cache[key] = value          // записать значение в инициализированный(!) словарь
delete(cache, key)          // удалить ключ из словаря, работает всегда
```

### Требования к ключам

Ключом может быть любо типа данных, для которого определена операция сравнения `==`:
- строки, числовые типы, `bool`, каналы (`chan`);
- интерфесы;
- указатели;
- структуры или массивы содержащие сравнимые типы.

```go
type User struct {
    Name string
    Host string
}
var cache map[User][]Permission
```

## Использование Zero Values

Для сласов и словаре, zero value — это `nil`.

В мапах с таким значением будут работать функции и операции читающие данные, например:

```go
var seq []string            // nil
var cache map[string]string // nil
l := len(seq)               // 0
c := cap(seq)               // 0
l := len(cache)             // 0
v, ok := cache[key]         // "", false
```

Для слайсов будет так же работать `append`

```go
var seq []string            // nil
seq = append(seq, "hello")  // []string{"hello"}
```

## Структуры

Структуры — фиксированный набор именованных переменных.
Переменные размещаются рядом в памяти и обычно используются совместно.

```go
struct{}            // Пустая структура, не занимает памяти
type User struct {  // Структура с именованными полями
    Id int64
    Name string
    Age int
    friends []int64 // Приватный элемент
}
```

### Литералы структур

```go
var u0 User                     // Zero Value для типа User
u1 := User{}                    // Zero Value для типа User
u2 := &User{}                   // То же, но указатель
u3 := User{1, "Vasya", 23, nil} // По номерам полей
u4 := User{                     // По именам полей
    Id: 1,
    Name: "Vasya",
    friends: []int64{1, 2, 3},
}
```

### Размер и выравнивание структур

```go
unsafe.Sizeof(1)    // 8 на моей машине
unsafe.Sizeof("A")  // 16 (длина + указатель)
var x struct {
    a bool          // 1 (offset 0)
    c bool          // 1 (offset 1)
    b string        // 16 (offset 8)
}
unsafe.Sizeof(x)    // 24!
```
### Анонимные типы и структуры

Анонимные типы задаются литералом, у такого типа нет имени.
Типичный сценарий использования: когда структура нужна только внутри одной функции.

```go
var wordCounts []struct{w string; n int}
```

```go
var resp struct {
    Ok          bool        `json:"ok"`
    Total       int         `json:"total"`
    Documents   []struct{
    Id          int         `json:"id"`
    Title       string      `json:"title"`
    } `json:"documents"`
}
json.Unmarshal(data, &resp)
fmt.Println(resp.Documents[0].Title)
```

### Встроенные структуры

В Go есть возможность "встраивать" типы внутрь структур.
При этом у элемента структуры НЕ задается имя.

```go
type LinkStorage struct {
    sync.Mutex                  // Только тип!
    storage map[string]string   // Тип и имя
}
```

Обращение к элементам встроенных типов:

```go
var storage LinkStorage
storage.Mutex.Lock()        // Имя типа используется
storage.Mutex.Unlock()      // как имя элемента структуры
```

### Тэги элементов структуры

К элементам структуры можно добавлять метаинформацию — тэги.
Тэг это просто литерал строки, но есть соглашение о структуре такой строки.

Например,

```go
type User struct {
    Id int64        `json:"-"` // Игнорировать в encode/json
    Name string     `json:"name"`
    Age int         `json:"user_age" db:"how_old"`
    friends []int64
}
```

Получить информацию о тэгах можно через `reflect`

```go
var u User
ageField := reflect.TypeOf(u).FieldByName("Age")
jsonFieldName := ageField.Get("json")               // "user_age"
```

### Пустые структуры

```go
type Set map[int]struct{}
```

```go
ch := make(chan struct{})
ch <- struct{}{}
```

### Экспортируемые и приватные элементы

Поля структур, начинающиеся со строчной буквы — **приватные**, они будут видны только в том же пакете, где и структура.

Поля, начинающиеся с заглавной — **публичные**, они будут видны везде.

```go
type User struct {
    Id int64
    Name string         // Экспортируемое поле
    Age int
    friends []int64     // Приватное поле
}
```

Не совсем очевидное следствие: пакеты стандартной библиотеки, например, encoding/json тоже не могут работать с приватными полями :)

Доступ к приватным элементам (на чтение!) все же можно получить с помощью пакета `reflect`.


## Объявление функции

```go
//       Имя функции         возвращаемые значения
//          |                     |      |
func TrySayHello(name string) (string, error)
//                 |    |
//            параметр тип параметра
 ```

Интересное:
- нет дефолтных значений для параметров
- функция может возвращать несколько значений
- функция — `first class value`, можем работать как с обычным значением
- параметры в функцию передаются по **значению**

### Примеры функций

```go
func Hello() {
    fmt.Println("Hello World!")
}
func add(x int, y int) int {
    return x + y
}
func add(x, y int) int {
    return x + y
}
func addMult(a, b int) (int, int) {
    return a + b, a * b
}

```

### Пример variadic функции

```go
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
 fmt.Println(total)
}
func main() {
    sum(5, 7)
    sum(3, 2, 1)
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
```

### Анонимные функции

```go
func main() {
    func() {
        fmt.Println("Hello ")
    }()

    sayWorld := func() {
        fmt.Println("World!")
    }

    sayWorld()
}

```

### Определение методов

В Go можно определять методы у именованых типов (кроме интерфейсов)

```go
type User struct {
    Id int64
    Name string
    Age int
    friends []int64
}
func (u User) IsOk() bool {
    for _, fid := range u.friends {
        if u.Id == fid {
            return true
        }
    }
    return false
}
var u User
fmt.Println(u.IsOk()) // (User).IsOk(u)
```

## Замыкания

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
 fmt.Println(nextInt()) // 1
 fmt.Println(nextInt()) // 2
 fmt.Println(nextInt()) // 3
 newInts := intSeq()
 fmt.Println(newInts()) // 1
}
```

## Методы типа и указателя на тип

Методы объявленные над типом получают копию объекта, поэтому не могут его изменять!

```go
func (u User) HappyBirthday() {
    u.Age++        // Это изменение будет потеряно
}
```

Методы объявленные над указателем на тип — могут.

```go
func (u *User) HappyBirthday() {
    u.Age++ // OK
}
```

Метод типа можно вызывать у значения и у указателя.
Метод указателя можно вызывать у указателя и у значения, если оно адресуемо.

### Функции-конструкторы

В Go принят подход Zero Value: постарайтесь сделать так, что бы ваш тип работал без инициализации, как реализованы, например

```go
var b strings.Builder
var wg sync.WaitGroup
```

Если ваш тип содержит словари, каналы или инициализация обязательна — скройте ее от пользователя, создав функции-конструкторы:

```go
func NewYourType() (*YourType) {
 // ...
}
func NewYourTypeWithOption(option int) (*YourType) {
 // ...
}
```

## Указатели

Указатель — это адрес некоторого значения в памяти.
Указатели строго типизированы.
Zero Value для указателя — `nil`.

```go
x := 1      // Тип int
xPtr := &x  // Тип *int
var p *int  // Тип *int, значение nil
```

### Получение адреса

Можно получать адрес не только переменной, но и поля структуры или элемента массива или слайса. Получение адреса осуществляется с помощью оператора `&`.

```go
var x struct {
    a int
    b string
    c [10]rune
}
bPtr := &x.b
c3Ptr := &x.c[2]
```

Но не значения в словаре!

```go
dict := map[string]string{"a": "b"}
valPtr := &dict["a"]                // Не скомпилируется
```

### Разыменование указателей

Разыменование осуществляется с помощью оператора `*`:

```go
a := "qwe"  // Тип string
aPtr := &a  // Тип *string
b := *aPtr  // Тип string, значение "qwe"
var n *int  // nil
nv := *n    // panic
```

В случае указателей на структуры вы можете обращаться к полям структуры без разыменования:

```go
p := struct{x, y int }{1, 3}
pPtr := &p
fmt.Println(pPtr.x)         // (*pPtr).x
fmt.Println((*pPtr).y)
pPtr = nil
fmt.Println(pPtr.x)         // panic
```

### Копирование указателей

В Go нет понятия передачи по ссылке — все всегда передается только по значению!
