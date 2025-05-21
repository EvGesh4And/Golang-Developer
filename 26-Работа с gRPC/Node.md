# gRPC (теория)

## О чем будем говорить

- Посмотрим на Protobuf
- Что такое gRPC

## План занятия

- Знакомимся с Protocol buffers
- Прямая и обратная совместимость в Protocol buffers
- Описание API с помощью Protobuf
- Что такое gRPC и HTTP/2

## Protobuf

```proto
syntax = "proto3";              // Указывает, что используется версия proto3

message SearchRequest {
  string query = 1;             // Текст поискового запроса
  int32 page_number = 2;        // Номер страницы (для пагинации)
  int32 results_per_page = 3;   // Количество результатов на странице
}
```

**Краткие пояснения:**

- `message SearchRequest` — определяет структуру сообщения с именем `SearchRequest`.

- `string` и `int32` — типы полей (строка и 32-битное целое число).

- `= 1`, `= 2`, `= 3` — теги полей. Они используются при сериализации и не должны изменяться после публикации схемы, чтобы обеспечить совместимость.

Это сообщение может использоваться, например, для отправки запроса на поиск в API, где клиент указывает поисковую строку и параметры пагинации.


### Protobuf: tags

- Номера тегов уникальны в рамках сообщения
- Номера от 1 до 536.870.911
- Номера от 19.000 до 19.999 заразервированы компилятором
- Теги с 1 до 15 занимают 1 байт


В **protobuf** **теги** (порядковые номера полей) **всегда начинаются с 1, а не с 0**.

Вот почему:

**🔢 Почему теги начинаются с 1**

- **Тег 0 зарезервирован** внутри Protocol Buffers и **не может использоваться** для пользовательских полей.

- Это сделано по архитектурным причинам: тег 0 в wire format зарезервирован как "неопределённое поле" и используется внутренне.

- Поэтому **всегда начинаем с 1 и далее: 1, 2, 3, ...** — и при этом важно не менять их после публикации схемы.

**🛠 Практические рекомендации:**

- **Каждое поле должно иметь уникальный тег** внутри одного message.
- Неизменяемость тегов критична: **если ты меняешь номер тега у существующего поля**, получатель не сможет правильно расшифровать сообщение.
- Если ты удаляешь поле — лучше не переиспользовать его тег, чтобы избежать конфликтов.

### Protocol buffers: типы данных

https://developers.google.com/protocol-buffers/docs/encoding

**✅ Скалярные типы protobuf3**

| Тип          | Размер          | Описание         |
| ------------ | --------------- | ---------------- |
| `float`/`double`      | 32 бит/64 бит   | Соответствует `float32`/`float64`. Фиксированный размер. |
| `fixed32`/`fixed64`   | 32 бит/64 бит   | Соответствует `uint32`/`uint64`. Беззнаковое целое фиксированного размера. **Эффективен, если часто передаёшь большие числа.** |
| `sfixed32`/`sfixed64` | 32 бит/64 бит   | Соответствует `int32`/`int64`. Беззнаковое целое фиксированного размера. **Эффективен, если часто передаёшь большие числа.** |
| `int32`/`int64`       | varint/varint   | Целое со знаком. Эффективен для положительных чисел (0...127 = 1 байт). Отрицательные — до 10 байт. |
| `uint32`/`uint64`   | varint/varint          | Целое без знака. Эффективен для маленьких чисел.   |
| `sint32`/`sint64`   | varint (ZigZag)/varint (ZigZag) | Поддерживает отрицательные значения эффективнее, чем `int32`. Использует ZigZag-кодирование.      |
| `bool `    | 1 байт          | `true` или `false`.            |
| `string`   | длина + UTF-8   | Строка в кодировке UTF-8. Сериализуется как длина + байты. Поддержка Unicode.  |
| `bytes`    | длина + `[]byte` | Сырые байты. Подходит для бинарных данных.      |

**💡 Особенности сериализации**

**📏 Varint**

- Используется в `int32`, `int64`, `uint32`, `uint64`, `sint32`, `sint64`.

- **Маленькие числа занимают меньше места**

  - Пример: `0 ... 127` → 1 байт

- **Отрицательные значения в** `int32`/`int64` → всегда 5/10 байт!

  - Поэтому для них лучше использовать `sint32`/`sint64` + ZigZag

**⚡ ZigZag (в `sint32`/`sint64`)**

ZigZag-кодирование позволяет эффективно кодировать отрицательные числа:

| Значение | ZigZag | Varint |
| -------- | ------ | ------ |
| 0        | 0      | 0x00   |
| -1       | 1      | 0x01   |
| 1        | 2      | 0x02   |
| -2       | 3      | 0x03   |


**📦 fixed vs varint**
- `fixed32` / `fixed64` **всегда 4 / 8 байт**, независимо от значения.
- **Используй их, если значения часто большие**, и ты не хочешь тратить до 10 байт на varint.

### Protocol buffers: repeated fields

Слайс реализуется через **repeated**:

```proto
message SearchResponse {
    repeated Result results = 1;
}
message Result {
    string url = 1;
    string title = 2;
    repeated string snippets = 3;
}
```

``` proto
...
Snippets    []string `protobuf:"bytes,3,rep,name=snippets,proto3" json:"snippets,omitempty"`
...
Results     []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
```

### Protocol buffers: Enums

```proto
enum EyeColor {
    EYE_COLOR_UNSPECIFIED = 0;
    EYE_COLOR_GREEN = 1;
    EYE_COLOR_BLUE = 2;
}
message Person {
    string name = 1;
    EyeColor eye_color = 2;
}
```

```go
type EyeColor int32
const (
    EyeColor_UNSPECIFIED    EyeColor = 0  // обязательное "нулевое" значение
    EyeColor_EYE_GREEN      EyeColor = 1
    EyeColor_EYE_BLUE       EyeColor = 2
)
```

### Protocol buffers: дефолтные значения

| Тип поля   | Значение по умолчанию     | Комментарий                                               |
| ---------- | ------------------------- | --------------------------------------------------------- |
| `string`   | `""`                      | Пустая строка                                             |
| `bool`     | `false`                   |                                                           |
| `int32`    | `0`                       | То же для `int64`, `uint32`, `uint64`, `float`, `double`  |
| `bytes`    | `[]byte{}` (пустой слайс) |                                                           |
| `enum`     | Первое объявленное        | Всегда начинается с нуля                                  |
| `repeated` | Пустой слайс              | Не `nil`, а именно пустой список (в сериализованном виде) |
| `message`  | `nil` в Go                | То есть поле отсутствует, если явно не задано             |

[Дефолтные значения](https://developers.google.com/protocol-buffers/docs/reference/go-generated#singular-message)


### Protocol buffers: oneof, map


`oneof` - только одно поле из списка может иметь значение и не может быть **repeated**.


```proto
message Message {
    int32 id = 1;
    oneof auth {
        string mobile = 2;
        string email = 3;
        int32 userid = 4;
    }
}
```

`map` - ассоциативный массив;
ключи - скаляры (кроме `float`/`double`);
значения - любые типы, не может быть **repeated**.

```proto
message Result {
    string result = 1;
}
message SearchResponse {
    map<string, Result> results = 1;
}
```


### Protocol buffers: wire types


В Protocol Buffers при сериализации каждое поле состоит из двух частей:

- key — кодирует номер поля (tag) и wire type,
- value — собственно данные.

**Всего 6 wire types:**

| Wire Type | Номер (3 бита) | Описание                              | Какие типы данных используется                           |
| --------- | -------------- | ------------------------------------- | -------------------------------------------------------- |
| 0         | 0              | **Varint** (переменная длина int)     | int32, int64, uint32, uint64, sint32, sint64, bool, enum |
| 1         | 1              | **64-битное фиксированное**           | fixed64, sfixed64, double                                |
| 2         | 2              | **Length-delimited** (длина + данные) | string, bytes, embedded messages, packed repeated fields |
| 3         | 3              | **Start group** (устаревшее)          | deprecated, не используется в proto3                     |
| 4         | 4              | **End group** (устаревшее)            | deprecated, не используется в proto3                     |
| 5         | 5              | **32-битное фиксированное**           | fixed32, sfixed32, float                                 |


**Детали:**

- **Wire type 0** — **varint**:

    Числа кодируются переменной длиной — по 7 бит на байт + 1 бит продолжения.
    Используется для целых чисел и enum.

- **Wire type 1** — **64-битное фиксированное**:

    Всегда 8 байт подряд. Для `double`, `fixed64` и `sfixed64`.

- **Wire type 2** — **length-delimited**:
    
    Для строк, байтовых массивов и вложенных сообщений.

- **Wire types 3 и 4** — группы:
    
    Устаревшие, не поддерживаются в proto3.

- **Wire type 5** — 32-битное фиксированное:
    
    Всегда 4 байта подряд. Для `float`, `fixed64` и `sfixed64`.

![alt text](image-1.png)

### Protocol buffers: encoding format

```proto
message Person {
    required string user_name = 1;
    optional int64 favorite_number = 2;
    repeated string interests = 3;
}
```

![alt text](image-3.png)

