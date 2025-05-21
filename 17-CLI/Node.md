# CLI (Command-Line Interface)

## Цели занятия

Научиться работать с ОС из программы на Go

## Краткое содержание

- обработка аргументов командной строки: `flag`, `pflag`, `cobra`
- работа с переменными окружения
- запуск внешних программ
- создание временных файлов
- обработка сигналов

## Соглашения и стандарты на CLI

- [POSIX](https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap12.html)

```
utility_name[-a][-b][-c option_argument]
    [-d|-e][-f[option_argument]][operand...]
```

- [GNU](https://www.gnu.org/prep/standards/standards.html#Command_002dLine-Interfaces)

```
utility_name -vutility_name --verbose
Соглашения и стандартны на CLI7 / 23
```

## `flag`/`pflag`

- [flag](https://pkg.go.dev/flag)
- [pflag](https://github.com/spf13/pflag)

📦 flag (стандартная библиотека)

- Простая, встроена в Go.
- Поддерживает только длинные флаги: `-flag=value`.
- Не поддерживает короткие (`-f`) и двойные (`--flag`) формы.
- Методы: `flag.String()`, `flag.Int()`, `flag.Bool()`, и т.д.
- Для пользовательских типов — `flag.Var()` (тип должен реализовать интерфейс `flag.Value`)

**Пример:**

```go
import "flag"

var name = flag.String("name", "world", "name to greet")

func main() {
    flag.Parse()
    fmt.Println("Hello", *name)
}
```

**Пример с `Var`:**

```go
type myBool bool

func (b *myBool) String() string   { return fmt.Sprint(*b) }
func (b *myBool) Set(s string) error {
    v, err := strconv.ParseBool(s)
    if err != nil {
        return err
    }
    *b = myBool(v)
    return nil
}

var b myBool
flag.Var(&b, "mybool", "custom bool flag")
```

🚀 pflag (из spf13/pflag)

- Расширение `flag`, совместим с ним.
- Поддерживает длинные `--flag` и короткие `-f`.
- Методы: `String()`, `Int()`, `Bool()` (аналогично `flag`).
- Добавлены методы с постфиксом `P`, например, `StringP()`, `BoolP()` — позволяют задать короткий и длинный флаг.
- Для пользовательских типов — `Var()` и `VarP()`.

**Пример:**

```go
import "github.com/spf13/pflag"

var name = pflag.StringP("name", "n", "world", "name to greet")

func main() {
    pflag.Parse()
    fmt.Println("Hello,", *name)
}
```

Вызов:
```
go run main.go -n Alice
go run main.go --name=Alice
```

📝 Итог:
| Фича           |	flag    |	   pflag |
|---------------|-----------|------------|
| Длинные флаги (`--flag`) |	    ❌ | ✅ |
| Короткие флаги (`-f`) |	        ✅ | ✅ |
| Совместимость с `flag` |	    ✅ | ✅ |
| Методы `Var()` |	            ✅ |	✅ |
| Методы `VarP()` |	            ❌ | ✅ |
| Методы `StringP()`, `BoolP()` |	❌ |	✅ |

### pflag: флаги без значений

```go
pflag.StringVar(&flagvar, "port", "80", "message to print")
pflag.Lookup("port").NoOptDefVal = "8080"
```

| Флаг          | Значение      |
|---------------|---------------|
| --port=9999   | flagvar=9999  |
| --port        | flagvar=8080  |
| [nothing]     | flagvar=80    |

### Сложные CLI приложения

```
git commit -m 123
docker pull
aws s3 ls s3://bucket-name
```

- https://clig.dev/
- https://github.com/spf13/cobra/
- https://github.com/urfave/cli



