# Тестирование в Go

## Преподаватель  
**Алексей Романовский**  

## Содержание занятия

- testing
- testify
- приемы тестирования

## Зачем нужны тесты?

- Упрощают рефакторинг (это процесс улучшения внутренней структуры кода без изменения его внешнего поведения).
- Документируют код
    Тесты = "живой" пример использования кода  
    - Их легче обновлять, чем комментарии в коде.
    - Они фиксируют ожидания и защищают от регрессий.
    - Они служат документацией, особенно в больших командах.
- Отделение интерфейса от реализации (mocks), менее связный код
    **Мок (mock)** — это поддельный объект, который имитирует поведение реальной зависимости в тестах. Он позволяет тестировать код изолированно, без необходимости вызывать настоящие внешние сервисы, базы данных или API.
- Помогают найти неактуальный код
- Помогают найти новые кейсы
    В процессе работы над созданием тестов могут быть найдены граничные, неочевидные кейсы, которые не были обработаны в коде
- Считают метрику для менеджмента (% покрытие)
- Определяют контракт
    **Контракт** — определённые требования к поведению системы должны быть удовлетворены, и эти требования явно формализуются через тесты
- Повышают качесиво кода
- Придают уверенности при деплое в продакшен

## Тестирование в Go

В Go для написания тестов используется пакет `testing`. Файлы с тестами должны иметь суффикс `_test.go`, а тестовые функции начинаться с `Test`.

### Простейший тест

Файл: `strings_test.go // <-- ..._test.go`


```go
package main

import (
    "strings"
    "testing"
)

func TestIndex(t *testing.T) { // <-- Test...(t *testing.T)
    const s, sub, want = "chicken", "ken", 4
    got := strings.Index(s, sub)
    if got != want {
        t.Errorf("Index(%q, %q) = %v; want %v", s, sub, got, want)
    }
}
```

### Разбор конструкций

1. Файл теста
    - Должен оканчиваться на `_test.go` (например, `strings_test.go`).

2. Импорт `testing`
    - Обязательно импортируется пакет `testing`.

3. Функция теста
    - Название начинается с `Test...` (например, `TestIndex`).
    - Принимает аргумент `t *testing.T`.

4. Проверка результата
    - Задаём входные данные (например, `const s, sub, want = "chicken", "ken", 4`)
    - Вызываем тестируемую функцию `strings.Index(s, sub)`
    - Сравниваем результат (`got`) с ожидаемым (`want`).
        - Если они не совпадают, используется `t.Errorf`, чтобы зафиксировать ошибку.

[TestIndex](https://goplay.tools/snippet/yybc8Np1JjK)

### Запуск тестов

```go
go test
```

или

```go
go test -v  // Для подробного вывода
```

Такой подход используется для написания юнит-тестов в Go.


### Вывод

- **Go требует, чтобы тесты запускались внутри модуля**.
- Названия тестов должны начинаться с `Test`.
- Ошибки фиксируются через `t.Errorf`.
- Запуск тестов — `go test -v`.
- Команда: `go test -v` **запускает все тесты во всех пакетах текущего модуля** (включая все файлы `*_test.go` в этих пакетах).
- `go test -v .` — тесты в текущем пакете.
- `go test -v ./mypackage` — тесты только в mypackage.
- `go test -v ./...` — тесты во всех подпакетах.

## testing: Error vs Fatal


В пакете **testing** используются методы `t.Error` и `t.Fatal` для логирования ошибок в тестах, но они имеют разное поведение.

### **`t.Error`**

- Логирует ошибку, но тест продолжает выполняться.
- **Позволяет фиксировать несколько ошибок в одном тесте**.
- Эквивалентен `t.Log` + `t.Fail`.

**Пример:**
```go
func TestExample(t *testing.T) {
    t.Error("This is an error, but test continues")
    fmt.Println("This will be printed")
}
```

### **`t.Fatal`**

- Логирует ошибку и немедленно завершает выполнение текущего теста.
- Код после `t.Fatal` не выполняется.
- Эквивалентен `t.Log` + `t.FailNow`.

**Пример:**
```go
func TestExample(t *testing.T) {
    t.Fatal("This is a fatal error, test stops here")
    fmt.Println("This will NOT be printed")
}
```

### Различия

| Функция   | Логирует сообщение | Продолжает выполнение теста | Вызывает `FailNow` |
|-----------|------------------|---------------------------|--------------------|
| `t.Error` | ✅ | ✅ | ❌ |
| `t.Fatal` | ✅ | ❌ | ✅ |

Вывод:

- Использовать `t.Error`, если хотиv зафиксировать несколько ошибок в одном тесте.
- Использовать `t.Fatal`, если дальнейший смысл теста теряется при первой ошибке.
- 

```go
func TestAtoi(t *testing.T) {

	const str, want = "43a", 42
	got, err := strconv.Atoi(str)
	if err != nil {
		t.Fatalf("strconv.Atoi(%q) returns unexpeted error: %v", str, err)
	}
	if got != want {
		t.Errorf("strconv.Atoi(%q) = %v; want %v", str, got, want)
	}

	fmt.Printf("The end of the test")
}
```

[TestAtoi](https://goplay.tools/snippet/vjAsrBrQrxu)

## testing: практика

### Задание
- Дописать существующие тесты.
- Придумать один новый тест.

```go
package main

import (
	"testing"

	"github.com/kulti/titlecase"
)

// TitleCase(str, minor) returns a str string with all words capitalized except minor words.
// The first word is always capitalized.
//
// E.g.
// TitleCase("the quick fox in the bag", "") = "The Quick Fox In The Bag"
// TitleCase("the quick fox in the bag", "in the") = "The Quick Fox in the Bag"

func TestEmpty(t *testing.T) {
	const str, minor, want = "", "", ""
	got := titlecase.TitleCase(str, minor)
	if got != want {
		t.Errorf("TitleCase(%v, %v) = %v; want %v", str, minor, got, want)
	}
}

func TestWithoutMinor(t *testing.T) {
	t.Error("not implemented")
}

func TestWithMinorInFirst(t *testing.T) {
	t.Error("not implemented")
}
```

## testify

[testify](https://github.com/stretchr/testify)

### Assert и Require в `testify`  

Пакеты `assert` и `require` в `testify` облегчают проверку условий в тестах Go.  
Они работают схоже, но `require` немедленно завершает тест (`t.FailNow()`),  
а `assert` просто фиксирует ошибку и продолжает выполнение (`t.Fail()`).  

#### Разница между `assert` и `require`  
| Функция  | Аналог в `testing`  | `assert` (ошибка, но тест продолжается) | `require` (фатальная ошибка, тест завершается) |
|----------|---------------------|------------------------------------------|-----------------------------------------------|
| Равенство значений | `if a != b { t.Errorf(...) }` | `assert.Equal(t, expected, actual)` | `require.Equal(t, expected, actual)` |
| Не равно | `if a == b { t.Errorf(...) }` | `assert.NotEqual(t, expected, actual)` | `require.NotEqual(t, expected, actual)` |
| Проверка `nil` | `if a != nil { t.Errorf(...) }` | `assert.Nil(t, obj)` | `require.Nil(t, obj)` |
| Проверка не `nil` | `if a == nil { t.Errorf(...) }` | `assert.NotNil(t, obj)` | `require.NotNil(t, obj)` |
| Логическое `true` | `if !a { t.Errorf(...) }` | `assert.True(t, condition)` | `require.True(t, condition)` |
| Логическое `false` | `if a { t.Errorf(...) }` | `assert.False(t, condition)` | `require.False(t, condition)` |
| Ошибка `nil` | `if err != nil { t.Errorf(...) }` | `assert.NoError(t, err)` | `require.NoError(t, err)` |
| Ошибка не `nil` | `if err == nil { t.Errorf(...) }` | `assert.Error(t, err)` | `require.Error(t, err)` |
| Подстрока в строке | `if !strings.Contains(a, b) { t.Errorf(...) }` | `assert.Contains(t, str, substr)` | `require.Contains(t, str, substr)` |

#### Дополнительные функции  
| Функция  | Описание  |
|----------|-----------|
| `assert.Len(t, obj, expectedLen)` | Проверяет длину объекта (слайс, строка, map). |
| `assert.Greater(t, a, b)` | `a > b` |
| `assert.Less(t, a, b)` | `a < b` |
| `assert.ElementsMatch(t, slice1, slice2)` | Проверяет, что слайсы содержат одни и те же элементы (в любом порядке). |
| `assert.JSONEq(t, json1, json2)` | Сравнивает JSON-объекты по содержимому. |
| `assert.Panics(t, func())` | Проверяет, что вызов паникует. |
| `assert.NotPanics(t, func())` | Проверяет, что вызов не паникует. |

#### Пример кода  
```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
    result := 42

    assert.Equal(t, 42, result) // Успех
    assert.NotEqual(t, 0, result) // Успех
    assert.Greater(t, result, 10) // Успех
    assert.Less(t, result, 100) // Успех
}
```

⚡ Используйте `assert`, если хотите проверить много условий в одном тесте.  
⚡ Используйте `require`, если при ошибке тест не имеет смысла и должен сразу завершиться.


```go
func TestAtoi(t *testing.T) {
    const str, want = "42", 42
    got, err := strconv.Atoi(str)
    require.NoError(t, err)
    require.Equal(t, want, got)
}
```
## Табличные тесты


```go
func TestParseInt(t *testing.T) {
    tests := []struct {
        str string
        expected int64
    }{
        {"-128", -128},
        {"0", 0},
        {"127", 127},
    }
    for _, tc := range tests {
        got, err := strconv.ParseInt(tc.str, 10, 8)
        require.NoError(t, err)
        require.Equal(t, tc.expected, got)
    }
}
func TestParseIntErrors(t *testing.T) {
    for _, str := range []string{"-129", "128", "byaka"} {
        _, err := strconv.ParseInt(str, 10, 8)
        require.Error(t, err)
    }
}
```

### t.Run

```go
func TestParseInt(t *testing.T) {
	tests := []struct {
		str      string
		expected int64
	}{
		{"-128", -128},
		{"0", 0},
		{"127", 127},
	}

	for _, tc := range tests {
		t.Run(tc.str, func(t *testing.T) {
			got, err := strconv.ParseInt(tc.str, 10, 8)
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}
```



## Coverage в Go


1. **Просмотр покрытия тестами**

   Для того чтобы увидеть покрытие тестами, выполните команду:
   ```bash
   go test -cover
   ```

   Эта команда выполнит тесты и отобразит процент покрытия кода тестами.

2. **Запись отчета о покрытии в файл**

   Чтобы записать отчет о покрытии в файл, используйте команду:
   ```bash
   go test -coverprofile=c.out
   ```

   Это создаст файл `c.out`, содержащий данные о покрытии. В нем будут указаны строки кода, которые были покрыты тестами, и те, которые не были.

3. **Просмотр отчета о покрытии в HTML-формате**

   Чтобы просмотреть отчет о покрытии в виде HTML-страницы, выполните команду:
   ```bash
   go tool cover -html=c.out
   ```

   Команда откроет отчет в браузере, который наглядно покажет, какие строки кода покрыты тестами, а какие нет.

### Примечания

- Покрытие тестами позволяет оценить, насколько полно код тестируется, и помогает выявить его неохваченные участки.
- Использование опции `-coverprofile` позволяет сохранить подробный отчет для дальнейшего анализа.
- HTML-отчет помогает визуально анализировать результаты тестирования.


## Golden Files

### Определение

**Golden files** (или **"золотые файлы"**) — это файлы, которые содержат заранее подготовленные, правильные данные, используемые для сравнения с результатами работы программы или теста. Золотые файлы обычно служат для того, чтобы убедиться, что изменения в коде не сломали ожидаемое поведение, особенно при тестировании функций, которые генерируют или обрабатывают данные (например, генерация отчетов, вывод в файлы или взаимодействие с внешними сервисами).

### Применение

Golden files часто используются в следующих сценариях:
- **Тестирование вывода программы**: Когда программа должна генерировать текст или другие данные, golden file сохраняет правильный вывод, с которым потом можно сравнивать актуальный результат.
- **Тестирование сериализации данных**: При тестировании сериализации/десериализации данных (например, в JSON или XML), golden file может содержать заранее сериализованные данные, которые будут использованы для проверки.
- **Тестирование API**: При тестировании API золотой файл может содержать заранее подготовленные ответы от сервера для проверки корректности обработки запросов.

### Пример использования

1. **Создание golden file**
   Обычно в первом тесте или вручную генерируется правильный результат (golden file). Например:
   ```txt
   {
       "name": "Test User",
       "age": 30
   }
   ```

2. **Тест с использованием golden file**
   В тесте используется золотой файл для сравнения. Например, при тестировании функции, которая сериализует объект в JSON:
   ```go
   func TestSerializeUser(t *testing.T) {
       user := User{Name: "Test User", Age: 30}
       serializedData, err := json.Marshal(user)
       if err != nil {
           t.Fatal("Failed to serialize:", err)
       }

       // Ожидаемый результат (golden file)
       goldenFile, err := os.ReadFile("golden_user.json")
       if err != nil {
           t.Fatal("Failed to read golden file:", err)
       }

       // Сравнение с golden file
       if !bytes.Equal(serializedData, goldenFile) {
           t.Errorf("Expected %s, but got %s", string(goldenFile), string(serializedData))
       }
   }
   ```

3. **Обновление golden file**
   Когда поведение программы меняется (например, формат вывода изменяется), golden file должен быть обновлён. Это можно сделать вручную или автоматически, добавив флаг в тесты для записи нового результата в golden file:
   ```go
   func updateGoldenFile(t *testing.T, filename string, data []byte) {
       err := os.WriteFile(filename, data, 0644)
       if err != nil {
           t.Fatal("Failed to update golden file:", err)
       }
   }
   ```

### Преимущества

1. **Контроль над ожидаемыми результатами**: Golden files предоставляют четкое ожидание для результатов, что помогает легко отслеживать изменения.
2. **Устранение ложных срабатываний**: Эти файлы помогают избежать ложных срабатываний, когда код изменяется, но результат все равно остается корректным.
3. **Простота внедрения**: Золотые файлы легко интегрировать в существующие тесты и легко использовать для сравнения с любыми типами данных.

### Недостатки

1. **Трудности при обновлении**: При изменении структуры или формата данных golden files необходимо обновлять вручную, что может быть трудоемким процессом.
2. **Зависимость от версий**: Золотые файлы могут стать неподходящими, если изменяется версия программы или данные, с которыми работает тест.
3. **Поддержка больших файлов**: Когда golden files становятся слишком большими (например, для большого объема данных), это может вызвать проблемы с хранением и использованием.

### Рекомендации

- **Малые данные**: Используйте golden files для проверки вывода небольших данных, например, результатов сериализации.
- **Избегайте частых изменений**: Постарайтесь минимизировать изменения золотых файлов, чтобы не нарушать стабильность тестов.
- **Автоматизация обновлений**: Внедрите механизм автоматического обновления golden files, если это возможно.

### Заключение

Golden files — это мощный инструмент для тестирования программ, который позволяет легко сравнивать актуальные результаты с заранее ожидаемыми. Они обеспечивают точность и стабильность тестов, однако требуют внимательности при обновлениях и поддержке.
