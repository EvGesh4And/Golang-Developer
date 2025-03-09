# Примеры использования каналов

[Статья](https://go101.org/article/channel-use-cases.html)


Перед прочтением этой статьи рекомендуется ознакомиться с материалом «[Каналы в Go](https://go101.org/article/channel.html)», где подробно объясняются типы и значения каналов. Новичкам в Go может потребоваться несколько раз перечитать обе статьи, чтобы лучше понять программирование с использованием каналов.

В оставшейся части статьи будут рассмотрены различные сценарии использования каналов. Надеюсь, что этот материал убедит вас в том, что:

- асинхронное и конкурентное программирование с каналами в Go — это просто и удобно;
- техника синхронизации с помощью каналов более универсальна и гибка, чем некоторые другие решения, используемые в других языках, такие как [модель акторов](https://en.wikipedia.org/wiki/Actor_model) или [паттерн async/await](https://en.wikipedia.org/wiki/Async/await).

Важно помнить, что цель данной статьи — продемонстрировать как можно больше вариантов использования каналов. Однако каналы — не единственный метод синхронизации в Go, и в некоторых случаях их использование может быть не самым эффективным решением. Чтобы изучить альтернативные подходы, ознакомьтесь со статьями о[ атомарных операциях](https://go101.org/article/concurrent-atomic-operation.html) и [других техниках синхронизации в Go](https://go101.org/article/concurrent-synchronization-more.html).

## Использование каналов как Future/Promise
В языках программирования, таких как JavaScript, Python и Java, широко применяются Future и Promise. Они используются для обработки асинхронных запросов и получения результатов после их выполнения.

### Возвращение каналов только для получения (`<-chan`) в качестве результата

Рассмотрим следующий пример. Функция `sumSquares` вычисляет сумму квадратов двух чисел, получаемых асинхронно. Операции получения значений из каналов блокируются до тех пор, пока в соответствующий канал не будет отправлено значение. В результате вычисление занимает всего три секунды, а не шесть, что позволяет значительно сократить время выполнения.

```go
package main

import (
	"time"
	"math/rand"
	"fmt"
)

func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		// Simulate a workload.
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20

	a, b := longTimeRequest(), longTimeRequest()
	fmt.Println(sumSquares(<-a, <-b))
}
```

### Передача каналов только для отправки (`chan<-`) в качестве аргументов

Как и в предыдущем примере, в следующем коде два аргумента функции `sumSquares` запрашиваются конкурентно. Однако, в отличие от прошлого примера, функция `longTimeRequest` принимает канал только для отправки (`chan<- int`) в качестве параметра, вместо того чтобы возвращать канал только для чтения (`<-chan int`).

```go
package main

import (
	"time"
	"math/rand"
	"fmt"
)

func longTimeRequest(r chan<- int32)  {
	// Simulate a workload.
	time.Sleep(time.Second * 3)
	r <- rand.Int31n(100)
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20

	ra, rb := make(chan int32), make(chan int32)
	go longTimeRequest(ra)
	go longTimeRequest(rb)

	fmt.Println(sumSquares(<-ra, <-rb))
}
```

На самом деле, в указанном выше примере нам не нужно два канала для передачи результатов. Достаточно использовать один канал.

```go
...

	// The channel can be buffered or not.
	results := make(chan int32, 2)
	go longTimeRequest(results)
	go longTimeRequest(results)

	fmt.Println(sumSquares(<-results, <-results))
}
```

Это своего рода агрегация данных, о которой будет сказано дальше.



### Первый ответ побеждает

Это улучшенная версия варианта с использованием одного канала из предыдущего примера.

Иногда данные можно получить из нескольких источников, чтобы снизить задержки. Однако из-за различных факторов время отклика разных источников может значительно отличаться. Даже у одного источника время ответа не всегда одинаково.

Чтобы минимизировать задержку, можно отправить запрос ко всем источникам одновременно в отдельных горутинах. При этом будет использован только первый полученный ответ, а остальные (более медленные) будут проигнорированы.

Важное замечание
Если источников **N**, то ёмкость канала должна быть не менее **N-1**, чтобы избежать блокировки горутин, чьи ответы были отброшены.

```go
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3) + 1
	// Sleep 1s/2s/3s.
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
}

func main() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20

	startTime := time.Now()
	// c must be a buffered channel.
	c := make(chan int32, 5)
	for i := 0; i < cap(c); i++ {
		go source(c)
	}
	// Only the first response will be used.
	rnd := <- c
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}
```

Существует несколько других способов реализовать сценарий «первый ответ побеждает», используя механизм `select` и буферизированный канал с ёмкостью 1. Другие методы будут рассмотрены ниже.

### Другие варианты запросов-ответов

Каналы параметров и результатов могут быть буферизированными, чтобы стороне, отправляющей ответ, не приходилось ждать, пока сторона, отправляющая запрос, заберёт переданные данные.

Иногда запрос может не вернуть корректный ответ. По разным причинам может возникнуть ошибка. В таких случаях можно использовать структуру вида `struct{ v T; err error }` или пустой интерфейс `interface{}` в качестве типа данных канала.

В некоторых случаях ответ может приходить гораздо дольше, чем ожидалось, или вовсе не приходить. В таких ситуациях можно использовать механизм тайм-аута, который будет рассмотрен далее.

Иногда ответ может представлять последовательность значений. Это своего рода механизм потока данных, который будет описан далее.

### Использование каналов для уведомлений

Уведомления можно рассматривать как особый вид запросов-ответов, в которых значение ответа не имеет значения.

Обычно в качестве типа данных используется пустая структура `struct{}`, так как её размер равен нулю, а значит, её значения не занимают память.

#### Уведомления 1-к-1 через отправку значения в канал

Если в канале нет доступных значений для чтения, следующая операция получения заблокируется, пока другая горутина не отправит туда данные.

Таким образом, можно отправлять значение в канал, чтобы уведомить другую горутину, которая ожидает данные из этого же канала.

В следующем примере канал `done` используется как сигнальный канал для отправки уведомлений.

```go
package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)

func main() {
	values := make([]byte, 32 * 1024 * 1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // can be buffered or not

	// The sorting goroutine
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		// Notify sorting is done.
		done <- struct{}{}
	}()

	// do some other things ...

	<- done // waiting here for notification
	fmt.Println(values[0], values[len(values)-1])
}
```

#### Уведомление 1-к-1 через получение значения из канала

Если буфер значений канала заполнен (**VBQ**) (у небуферизированного канала буфер значений всегда считается заполненным), то операция отправки в канал заблокируется, пока другая горутина не получит значение из этого канала.

Таким образом, можно получать значение из канала, чтобы уведомить другую горутину, которая ожидает возможности отправить данные в этот же канал.

Как правило, для такого способа уведомления используется небуферизированный канал.

Этот метод реже используется, чем способ, рассмотренный в предыдущем примере.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
		// The capacity of the signal channel can
		// also be one. If this is true, then a
		// value must be sent to the channel before
		// creating the following goroutine.

	go func() {
		fmt.Print("Hello")
		// Simulate a workload.
		time.Sleep(time.Second * 2)

		// Receive a value from the done
		// channel, to unblock the second
		// send in main goroutine.
		<- done
	}()

	// Blocked here, wait for a notification.
	done <- struct{}{}
	fmt.Println(" world!")
}
```

На самом деле, **фундаментальных различий** между уведомлением через отправку и получение значений нет. Оба подхода можно обобщить так: быстрые горутины уведомляются медленными.

#### Уведомления N-к-1 и 1-к-N

Небольшое расширение рассмотренных выше случаев позволяет легко реализовать механизмы уведомлений **N-к-1** и **1-к-N**.

```go
package main

import "log"
import "time"

type T = struct{}

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // block here and wait a notification
	log.Print("Worker#", id, " starts.")
	// Simulate a workload.
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("Worker#", id, " job done.")
	// Notify the main goroutine (N-to-1),
	done <- T{}
}

func main() {
	log.SetFlags(0)

	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	// Simulate an initialization phase.
	time.Sleep(time.Second * 3 / 2)
	// 1-to-N notifications.
	ready <- T{}; ready <- T{}; ready <- T{}
	// Being N-to-1 notified.
	<-done; <-done; <-done
}
``` 

На самом деле, методы 1-к-N и N-к-1 уведомлений, рассмотренные в этом разделе, редко используются на практике.

Как делают в реальных задачах:
Для N-к-1 уведомлений чаще применяют sync.WaitGroup.
Для 1-к-N уведомлений обычно закрывают канал.
Подробнее об этом читайте в следующем разделе.

#### Широковещательные (1-к-N) уведомления через закрытие канала

Метод 1-к-N уведомлений из предыдущего раздела практически не используется, так как существует более удобный способ.

Благодаря тому, что из закрытого канала можно бесконечно получать значения, можно просто закрыть канал, чтобы расслать уведомления всем горутинам.

Например, в предыдущем примере три операции отправки `ready <- struct{}{}` можно заменить одной операцией `close(ready)`, чтобы реализовать 1-к-N уведомление.

```go
...
	close(ready) // broadcast notifications
...
```


Конечно, закрытие канала можно использовать и для 1-к-1 уведомлений.
На практике это самый распространённый способ уведомлений в Go.

Особенность, что из закрытого канала можно бесконечно получать значения, используется во многих других сценариях, которые будут рассмотрены ниже.
Более того, эта особенность активно применяется в стандартной библиотеке Go.

Пример:
Пакет `context` использует этот механизм для обработки отмены операций.


### Таймер: Запланированное уведомление

С помощью каналов легко реализовать одноразовые таймеры.

Пример реализации кастомного одноразового таймера:


```go
package main

import (
	"fmt"
	"time"
)

func AfterDuration(d time.Duration) <- chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()
	return c
}

func main() {
	fmt.Println("Hi!")
	<- AfterDuration(time.Second)
	fmt.Println("Hello!")
	<- AfterDuration(time.Second)
	fmt.Println("Bye!")
}
```


Функция `After` из стандартного пакета `time` предоставляет ту же функциональность,
но с гораздо более эффективной реализацией.

Лучше использовать `time.After(aDuration)`,
так код будет выглядеть чище и понятнее.

Важно:
Выражение `<-time.After(aDuration)`
заблокирует выполнение текущей горутины,
в то время как вызов `time.Sleep(aDuration)` не создаёт блокирующей операции.

Механизм `<-time.After(aDuration)` часто используется для реализации таймаутов, о которых будет рассказано дальше.

## Использование каналов как мьютексов
Ранее уже упоминалось, что буферизированные каналы с ёмкостью 1 могут быть использованы как одноразовые [бинарные семафоры](https://en.wikipedia.org/wiki/Semaphore_(programming)).

На самом деле, такие каналы можно использовать и как мьютексы, но они менее эффективны, чем мьютексы из пакета `sync`.

Способы использования каналов как мьютексов:
- Блокировка через отправку (`send`), разблокировка через приём (`receive`).
- Блокировка через приём (`receive`), разблокировка через отправку (`send`).

Следующий пример демонстрирует вариант с блокировкой через отправку.

```go
package main

import "fmt"

func main() {
	// The capacity must be one.
	mutex := make(chan struct{}, 1)

	counter := 0
	increase := func() {
		mutex <- struct{}{} // lock
		counter++
		<-mutex // unlock
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done; <-done
	fmt.Println(counter) // 2000
}
```

Следующий пример демонстрирует вариант с блокировкой через приём (`receive`). Он показывает только изменённую часть по сравнению с предыдущим примером, где использовалась блокировка через отправку (`send`).

```go
...
func main() {
	mutex := make(chan struct{}, 1)
	mutex <- struct{}{} // this line is needed.

	counter := 0
	increase := func() {
		<-mutex // lock
		counter++
		mutex <- struct{}{} // unlock
	}
...
```

### Использование каналов как счётных семафоров
Буферизированные каналы можно использовать в качестве [счётных семафоров](https://en.wikipedia.org/wiki/Semaphore_(programming)). Счётные семафоры можно рассматривать как многопользовательские блокировки. Если ёмкость канала равна **N**, то его можно воспринимать как блокировку,
у которой в любой момент времени может быть не более **N** владельцев.

Бинарные семафоры (мьютексы) являются частным случаем счётных семафоров, где в каждый момент времени может быть не более одного владельца.

Счётные семафоры часто используются для ограничения количества одновременных запросов.

Способы использования каналов как семафоров:
Как и при использовании каналов в качестве мьютексов, есть два способа получения прав владения семафором:

- Получение права через **отправку** (**send**), освобождение через **приём** (**receive**).
- Получение права через **приём** (**receive**), освобождение через отправку (**send**).

Пример получения права владения через приём значений из канала:

```go
package main

import (
	"log"
	"time"
	"math/rand"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(c int) {
	log.Print("customer#", c, " enters the bar")
	seat := <- bar // need a seat to drink
	log.Print("++ customer#", c, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
	log.Print("-- customer#", c, " frees seat#", seat)
	bar <- seat // free seat and leave the bar
}

func main() {

	// the bar has 10 seats.
	bar24x7 := make(Bar, 10)
	// Place seats in an bar.
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		// None of the sends will block.
		bar24x7 <- Seat(seatId)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		go bar24x7.ServeCustomer(customerId)
	}

	// sleeping != blocking
	for {time.Sleep(time.Second)}
}
```

В приведённом выше примере только те клиенты, которые получили место, могут пить. Следовательно, в любой момент времени не более десяти клиентов могут пить.

Последний `for`-цикл в `main`-функции нужен, чтобы программа не завершилась раньше времени.
Существует более правильный способ решения этой задачи, который будет представлен далее.

Хотя одновременно пить могут не более десяти клиентов,
в баре может находиться большее количество клиентов.
Некоторые клиенты ждут свободного места.

Хотя каждая горутина клиента потребляет гораздо меньше ресурсов, чем системный поток, суммарные затраты ресурсов при большом количестве горутин могут быть значительными.

Поэтому лучше создавать горутину клиента только в том случае, если есть свободное место.

```go
... // same code as the above example

func (bar Bar) ServeCustomerAtSeat(c int, seat Seat) {
	log.Print("++ customer#", c, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
	log.Print("-- customer#", c, " frees seat#", seat)
	bar <- seat // free seat and leave the bar
}

func main() {
	bar24x7 := make(Bar, 10)
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		// Need a seat to serve next customer.
		seat := <- bar24x7
		go bar24x7.ServeCustomerAtSeat(customerId, seat)
	}
	for {time.Sleep(time.Second)}
}
```


В оптимизированной версии программы в любой момент времени будет существовать **не более десяти активных** горутин клиентов (но за время работы программы все равно **будет создано много горутин клиентов**).

В более эффективной реализации, представленной ниже,
за всё время работы программы будет создано не более десяти обслуживающих горутин клиентов.


```GO
... // same code as the above example

func (bar Bar) ServeCustomerAtSeat(consumers chan int) {
	for c := range consumers {
		seatId := <- bar
		log.Print("++ customer#", c, " drinks at seat#", seatId)
		time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
		log.Print("-- customer#", c, " frees seat#", seatId)
		bar <- seatId // free seat and leave the bar
	}
}

func main() {
	bar24x7 := make(Bar, 10)
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId)
	}

	consumers := make(chan int)
	for i := 0; i < cap(bar24x7); i++ {
		go bar24x7.ServeCustomerAtSeat(consumers)
	}
	
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumers <- customerId
	}
}
```

Не по теме: конечно, если нам не важны идентификаторы мест (что часто встречается на практике), то **семафор bar24x7 вообще не нужен**.

```go
... // same code as the above example

func ServeCustomer(consumers chan int) {
	for c := range consumers {
		log.Print("++ customer#", c, " drinks at the bar")
		time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
		log.Print("-- customer#", c, " leaves the bar")
	}
}

func main() {
	const BarSeatCount = 10
	consumers := make(chan int)
	for i := 0; i < BarSeatCount; i++ {
		go ServeCustomer(consumers)
	}
	
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumers <- customerId
	}
}
```


Способ получения владения семафором через отправку сравнительно проще. Шаг размещения мест **не требуется**.

```go
package main

import (
	"log"
	"time"
	"math/rand"
)

type Customer struct{id int}
type Bar chan Customer

func (bar Bar) ServeCustomer(c Customer) {
	log.Print("++ customer#", c.id, " starts drinking")
	time.Sleep(time.Second * time.Duration(3 + rand.Intn(16)))
	log.Print("-- customer#", c.id, " leaves the bar")
	<- bar // leaves the bar and save a space
}

func main() {
	// The bar can serve most 10 customers
	// at the same time.
	bar24x7 := make(Bar, 10)
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second * 2)
		customer := Customer{customerId}
		// Wait to enter the bar.
		bar24x7 <- customer
		go bar24x7.ServeCustomer(customer)
	}
	for {time.Sleep(time.Second)}
}
```

### Диалог (Пинг-Понг)
Две горутины могут взаимодействовать через канал. Ниже приведён пример, который выводит последовательность чисел Фибоначчи.

```go
package main

import "fmt"
import "time"
import "os"

type Ball uint64

func Play(playerName string, table chan Ball) {
	var lastValue Ball = 1
	for {
		ball := <- table // get the ball
		fmt.Println(playerName, ball)
		ball += lastValue
		if ball < lastValue { // overflow
			os.Exit(0)
		}
		lastValue = ball
		table <- ball // bat back the ball
		time.Sleep(time.Second)
	}
}

func main() {
	table := make(chan Ball)
	go func() {
		table <- 1 // throw ball on table
	}()
	go Play("A:", table)
	Play("B:", table)
}
```






