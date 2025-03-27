---
marp: true
theme: default
class: invert
paginate: true
---

# Конфигурация и логирование
Лектор: Алексей Романовский


---


## Маршрут вебинара
* Знакомство
* О конфигурациях
* О логировании
* Рефлексия

---

## Цели вебинара
1. Узнать о существующих подходах к хранению конфигурации
2. Узнать об инструментах для работы с конфигурацией
3. Изучить подходы, использующиеся в логировании
4. Изучить существующие инструменты для вывода логов

---

## Смысл 
<sub>(зачем вам это уметь)<sub>
1. Улучшить безопасность и удобство работы с конфигурацией
2. Сделать доступнее анализ работы приложения через его логи 

---

# Конфигурация

---

# Вопрос
Что такое переменные окружения?

Напишите своё мнение в чат или просто “–”

---

## Что хранить в конфигурации?
Все данные, которые могут отличаться в разных развертываниях (у разных пользователей). 
Знайте свои развертывания!

---

* дефолтные значения
* секреты
* идентификаторы внешних ресурсов
  * строки подключения к бд
  *  хосты внешних апи
* feature flags
* подгоночные параметры (например, интервалы опросов)
* develop\staging\prod

---

## Чеклист 
для проработки конфигурации:
* Кто будет иметь доступ к конфигурации?
   * разраб, админ, менеджер, админ клиента, пользователь
* только чтение? надо менять на ходу? 
* кто создаст?
* где хранить (гит, вне гита - файлы или переменные окружения, бд, облако)


---

Отвечая на эти вопросы, 
вы скорее всего обнаружите,
что вам надо несколько разных систем

---

https://12factor.net/ru/config

---

## Вопрос
Какой ваш любимый подход к хранению конфигурации? 

Что бы вы использовали в новом проекте вида:
1. консольная утилита для коллег в команде
2. серверное приложение с API


Напишите своё мнение в чат или просто “–”

---
## Совет по архитектуре

* Организуйте отдельный пакет для конфига
* Другие части приложения не должны 
знать\зависеть от выбранного способа хранения конфигов.
* Типизируйте параметры конфига (вместо того чтобы всё было строками)

---
### Частая ошибка: 
Отдельный пакет есть, структура задана. Но все значения - строки. 

И форматы дат, булевых значений и т.п. из файла логов “растеклись” по всему приложению.

---

## Конфигурация в файле

---
```go 
import (
   "io/ioutil"
   "gopkg.in/yaml.v2"
)
type Config struct {
   Domain    string   `yaml:"domain"`
   Blacklist []string `yaml:"blacklist"`
}
func main() {
   var c Config
   yamlFile, err := ioutil.ReadFile("conf.yaml") 
   err = yaml.Unmarshal(yamlFile, &c)
}
```

---


```yaml
domain: abs.com
blacklist:
 - evil.com
 - bad.com
```

---

### Переменные окружения
(aka “environment variables”)

* Универсальный механизм для разных языков, платформ, ОС
* С ними знакомы многие и в т.ч. не разработчики
* Широко поддерживаются системами оркестрации
* Сложно случайно закоммитить
* Нельзя менять на ходу (извне)
* Особо удобно для секретов
* Рекомендованы [12factor.net](https://12factor.net/ru/config)


---
```shell
MYAPP_HOST=localhost MYAPP_PORT=7777 go run main.go

```


---

## Переменные окружения в Go
(без библиотек)

```go
type Config struct {
   Port int
   Host string
}

// ...
   httpPort, err := strconv.Atoi(os.Getenv("MYAPP_PORT"))
   if err != nil {
       panic(fmt.Sprint("MYAPP_PORT not defined"))
   }
   shortenerHost := os.Getenv("MYAPP_HOST")
   if shortenerHost == "" {
       panic(fmt.Sprint("MYAPP_HOST not defined"))
   }
   config := Config{httpPort, shortenerHost}

```


---

## Библиотеки для работы с конфигурацией	

Выбор библиотек:
* https://go.libhunt.com/categories/463-configuration
* https://github.com/avelino/awesome-go#configuration

---

Библиотеки для переменных окружения:
* envcfg - Un-marshaling environment variables to Go structs.
* envconf - Configuration from environment.
* **envconfig** - Read your configuration from environment variables.
* envh - Helpers to manage environment variables.
* godotenv - Go port of Ruby's dotenv library (Loads environment variables from .env).

https://github.com/avelino/awesome-go#configuration

---
### envconfig
"github.com/kelseyhightower/envconfig"

Анмаршаллинг данных из переменных окружения в структуру + генератор документации.

---

```go
type Config struct {
   ApiUrl      string        `required:"true"`
   WorkerCount int           `default:"1"`
   Interval    time.Duration `default:"1m"`
   LogLevel    zapcore.Level `default:"info" split_words:"true"`
}
const EnvVarPrefix = "myapp"
func main() {
   if len(os.Args) > 1 && (os.Args[1] == "--help") {
       err := envconfig.Usage(EnvVarPrefix, &Config{})
       if err != nil {
           panic(err)
       }
       return
   }

   config := Config{}
   envconfig.MustProcess(EnvVarPrefix, &config)
}
```

---

### Универсальные библиотеки\фреймворки
"github.com/spf13/viper"

Viper - комбинирует данные из нескольких источников, позволяет писать и отслеживать изменения.

---

### Приоритеты источников:
* explicit call to Set
* flag
* env
* config
* key/value store
* default

---

Рассмотрим пример использования. Если шрифт мелкий - говорите сразу!

https://github.com/alexus1024/go23_config_log/tree/main/config

---

### Универсальные библиотеки\фреймворки
"github.com/heetch/confita"

* поддерживает примитивы Go
* поддерживает несколько "бэкэндов"

---

```go
   loader := confita.NewLoader(
       env.NewBackend(),
       file.NewBackend("/path/to/config.json"),
       file.NewBackend("/path/to/config.yaml"),
       flags.NewBackend(),
       etcd.NewBackend(etcdClientv3),
       consul.NewBackend(consulClient),
       vault.NewBackend(vaultClient),
   )
```


---

# Логирование	

---

## Вопрос
Что такое лог?

Напишите своё мнение в чат или просто “–”

---

## Подходы в логировании
* Куда выводить (консоль, файлы, апи)
* Формат: структурированные или нет
* Если структурированные - то какой формат - json?
* Кто и как их будет читать? Глазами или парсить и смотреть в тулзе.
* Уровни логирования (debug/info/warning/error/fatal)

---

### Неструктурированное логирование (пакет log)
```go
import (
   "log"
)
func init() {
   log.SetPrefix("LOG: ")
   log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
   log.Println("init started")
}
func main() {
   // Println writes to the standard logger.
   log.Println("main started")
   // Fatalln is Println() followed by a call to os.Exit(1)
   log.Fatalln("fatal message")
   // Panicln is Println() followed by a call to panic()
   log.Panicln("panic message")
}
```

---

### Неструктурированное логирование (в файл)

```go
import (
   "log"
   "os"
)
func main() {
   file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
   if err != nil {
       log.Fatal(err)
   }
   defer file.Close()
   log.SetOutput(file)
   log.Print("Logging to a file in Go!")
}
```

---

# Структурированные 
# логи

---

Одна запись содержит несколько полей. 
* Часть из них есть всегда (время, уровень, сообщение)
* часть - уникальна для источника/контекста  (статус http или RequestID в апи)
* или даже для каждой записи (ид клиента в запросе к бд)

---

Могут формироваться для:
* чтения людьми 
* или для машинного разбора (json stream)


**Сообщение - константа**

---

### Структурированные логи

Выбор библиотеки
* [libhunt chart](https://go.libhunt.com/categories/504-logging)
* [Slog in go 1.21](https://go.dev/blog/slog) - похоже, это станет стандартом

---

### См также:
* [pkg.go.dev/log/slog](https://pkg.go.dev/log/slog)
* [просто статья с деталями](https://betterstack.com/community/guides/logging/logging-in-go/)


---

## Рассмотрим пример проекта
## "log/slog"

---

https://github.com/go-slog/awesome-slog

---

## Big picture
* Как сделать конфигурацию в проекте
* Конфигурация в файлах
* В переменных окужения
* Библиотеки работы с для конфигурацией
* envconfig
* viper, confita
* Логи, какие бывают
* Библиотеки для логов: slog, zap, logrus
