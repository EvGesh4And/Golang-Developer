# Concurrency patterns

## О чем будем говорить
- паттерны синхронизации данных;
- функции-генераторы и пайплайн;
- работа с многими каналами: or, fanin, fanout, etc.


## Конкурентный код

Глобально мы обеспечиваем безопасность за счет:
- примитивов синхронизации (e.g. sync.Mutex, etc)
- каналы
- confinement-техники

## Confinement-техники

Confinement-техники в программировании (и в частности в Go) используются для ограничения доступа к данным, чтобы избежать гонок данных и обеспечить безопасность параллельных вычислений.

Варианты confinement-техник:

1. Неизменяемые данные (Immutable data)

   - Идеальный вариант, так как неизменяемые структуры могут безопасно использоваться в разных потоках без необходимости синхронизации.
   - В Go структуры можно имитировать неизменяемыми, не экспортируя поля или передавая их копии:
      ```go
      type Config struct {
         value int
      }

      func NewConfig(v int) Config {
         return Config{value: v}
      }
      ```
   - Однако полная неизменяемость не всегда возможна, особенно при работе с большими структурами.

2. Ad hoc confinement
   - Временные или специфичные решения для ограничения доступа к данным.
   - Например, передача данных только в один поток (goroutine), где они используются:
   ```go
   ch := make(chan int, 1)

   go func() {
      ch <- 42 // Данные используются только в этой горутине
   }()
   ```
3. Lexical confinement (лексическое ограничение)

- Переменная доступна только в определённой области видимости, что предотвращает её конкурентное изменение.
- Например, использование переменной внутри одной горутины:
   ```go
   func processData() {
    data := 42 // Доступно только в этой функции (лексическая область видимости)
    fmt.Println(data)
   }
   ```
### Ad hoc

По сути, неявная договоренность, что "я - читаю, а ты пишешь", поэтому мы не используем никакие средства синхронизации.

```go
data := make([]int, 4)

loopData := func(handleData chan<- int) {
   defer close(handleData)
   for i := range data {
      handleData <- data[i]
   }
}

handleData := make(chan int)
go loopData(handleData)

for num := range handleData {
   fmt.Println(num)
}
```

### Lexical

Никакой договоренности нет, но она, по сути, неявно создана кодом.

```go
chanOwner := func() <-chan int {
   results := make(chan int, 5)
   go func() {
      defer close(results)
      for i := 0; i <= 5; i++ {
         results <- i
      }
   }()
   return results
}

consumer := func(results <-chan int) {
   for result := range results {
      fmt.Printf("Received: %d\n", result)
   }
   fmt.Println("Done receiving!")
}

results := chanOwner()
consumer(results)
```

### For-select цикл

**Пример 1**

```go
for _, i := range []int{1, 2, 3, 4, 5} {
   select {
   case <-done:
      return
   case intStream <- i:
   }
}
```

**Пример 2 (активное ожидание)**

```go
for {
   select {
   case <- done:
      return
   default:
   }
}
```

### Как предотвратить утечку горутин

Проблема:

```go
doWork := func(strings <-chan string) <-chan struct{} {
   completed := make(chan struct{})
   go func() {
      defer fmt.Println("doWork exited.")
      defer close(completed)
      for s := range strings {
         fmt.Println(s)
      }
   }()
   return completed
}

doWork(nil)
time.Sleep(time.Second * 5)
fmt.Println("Done.")
```

[Невидимые ошибки Go-разработчика. Артём Картасов](https://youtu.be/TVe8pIFn2mY)


Решение - **явный индиктор** того, что пора завершаться:

```go
doWork := func(done <-chan struct{}, strings <-chan string) <-chan struct{} {
   terminated := make(chan struct{})
   go func() {
      defer fmt.Println("doWork exited.")
      defer close(terminated)
      for {
         select {
         case s := <-strings:
            fmt.Println(s)
         case <-done:
            return
         }
      }
   }()
   return terminated
}
...
```








