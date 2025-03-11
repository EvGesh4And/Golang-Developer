# Как изящно закрыть каналы

[Статья](https://go101.org/article/channel-closing.html)


Несколько дней назад я написал статью, в которой объясняются [правила работы с каналами в Go](https://go101.org/article/channel.html). Эта статья получила много голосов на [Reddit](https://www.reddit.com/r/golang/comments/5k489v/the_full_list_of_channel_rules_in_golang/) и [HN](https://news.ycombinator.com/item?id=13252416), но также вызвала некоторые критические замечания относительно деталей дизайна каналов в Go.

Я собрал несколько распространённых критических замечаний о следующих аспектах каналов в Go:

1. Нет простого и универсального способа проверить, закрыт ли канал, без изменения его состояния.
2. Закрытие уже закрытого канала вызовет панику, поэтому опасно закрывать канал, если вызывающий не уверен в его статусе.
3. Отправка значений в закрытый канал вызовет панику, поэтому опасно отправлять данные в канал, если отправитель не знает, закрыт он или нет.

Эти замечания кажутся обоснованными, но на самом деле это не так. Действительно, в Go нет встроенной функции для проверки, закрыт ли канал.

Однако существует простой метод, позволяющий проверить, закрыт ли канал, при условии, что в него не отправлялись (и не будут отправляться) значения. Этот метод уже был показан в [предыдущей статье](https://go101.org/article/channel-use-cases.html#check-closed-status). Для удобства он приведён в следующем примере.

```go
package main

import "fmt"

type T int

func IsClosed(ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func main() {
	c := make(chan T)
	fmt.Println(IsClosed(c)) // false
	close(c)
	fmt.Println(IsClosed(c)) // true
}
```

Как упомянуто выше, это не универсальный способ проверки, закрыт ли канал.

Фактически, даже если бы в Go существовала встроенная функция `closed`, позволяющая проверять, закрыт ли канал, её полезность была бы очень ограниченной. Это похоже на встроенную функцию `len`, которая показывает количество элементов в буфере канала. Проблема в том, что состояние канала может измениться сразу после вызова такой функции, и возвращённое значение уже не будет актуальным.

Хотя допустимо прекратить отправку данных в канал `ch`, если `closed(ch)` возвращает `true`, нельзя безопасно продолжать отправку данных или закрывать канал, если `closed(ch)` возвращает `false`.

## Принцип закрытия каналов

Общий принцип работы с каналами в Go: **не закрывайте канал на стороне получателя и не закрывайте канал, если у него несколько отправителей**. Иными словами, закрывать **канал должна только та горутина, которая отправляет данные**, и только если **она является единственным отправителем**.

(Далее этот принцип будем называть **принципом закрытия каналов** - **channel closing principle**)

Разумеется, это не единственный возможный подход. Универсальный принцип: **не закрывать (и не отправлять значения) закрытые каналы**. Если можно гарантировать, что ни одна горутина больше не будет закрывать и отправлять данные в не закрытый и не nil-канал, тогда канал можно закрыть безопасно. Однако обеспечение таких гарантий со стороны получателя или одного из нескольких отправителей требует больших усилий и усложняет код. Напротив, придерживаться вышеупомянутого принципа закрытия каналов намного проще.

## Решения с "грубым" закрытием каналов
Если вам всё же необходимо закрыть канал на стороне получателя или в одной из нескольких горутин-отправителей, можно использовать [механизм recover](https://go101.org/article/control-flows-more.html#panic-recover), чтобы предотвратить панику при закрытии уже закрытого канала.

Пример (предположим, что канал передаёт элементы типа `T`):

```go
func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// The return result can be altered
			// in a defer function call.
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch)   // panic if ch is closed
	return true // <=> justClosed = true; return
}
```

Это решение, очевидно, нарушает **принцип закрытия каналов**.

Ту же идею можно использовать при отправке значений в потенциально закрытый канал.

```go
func SafeSend(ch chan T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()

	ch <- value  // panic if ch is closed
	return false // <=> closed = false; return
}
```

Грубое решение не только нарушает **принцип закрытия каналов**, но и может привести к **гонке данных (data races)** в процессе.

## Решения, которые вежливо (**politely**) закрывают каналы
Многие разработчики предпочитают использовать `sync.Once` для закрытия каналов:

```go
type MyChannel struct {
	C    chan T
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}
```

Конечно, мы также можем использовать `sync.Mutex`, чтобы избежать многократного закрытия канала:

```go
type MyChannel struct {
	C      chan T
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
}

func (mc *MyChannel) IsClosed() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}
```

Эти методы **могут быть аккуратными, но они не гарантируют отсутствие гонки данных (data races)**. В текущей спецификации Go нет гарантии, что при одновременном выполнении операций закрытия канала и отправки данных в канал не возникнет гонки данных. Если функция `SafeClose` вызывается одновременно с операцией отправки данных в тот же канал, может возникнуть гонка данных (хотя обычно оно не приводит к критическим ошибкам).

## Решения, которые закрывают каналы более элегантно

Один из недостатков функции `SafeSend`, описанной выше, заключается в том, что её вызовы нельзя использовать в качестве операций отправки, следующих за ключевым словом `case` в `select`-блоках. Другой недостаток функций `SafeSend` и `SafeClose` в том, что многие разработчики, включая меня, считают, что использование `panic/recover` и пакета `sync` в этих решениях не является элегантным.

Далее будут представлены решения, основанные только на механизме каналов, без использования `panic/recover` и пакета `sync`, подходящие для различных ситуаций.
(В следующих примерах используется `sync.WaitGroup`, чтобы сделать примеры завершёнными. Однако в реальной практике его использование не всегда обязательно.)

### 1. M получателей, один отправитель, отправитель сообщает "больше не будет отправок", закрывая канал данных

Это самый простой случай — просто позволить отправителю закрыть канал данных, когда он больше не хочет отправлять данные.

```go
package main

import (
	"time"
	"math/rand"
	"sync"
	"log"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)

	// the sender
	go func() {
		for {
			if value := rand.Intn(Max); value == 0 {
				// The only sender can close the
				// channel at any time safely.
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			// Receive values until dataCh is
			// closed and the value buffer queue
			// of dataCh becomes empty.
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}
```

### Один получатель, N отправителей, единственный получатель сообщает "пожалуйста, прекратите отправку" через закрытие дополнительного сигнального канала

Этот случай немного сложнее предыдущего. Мы не можем позволить получателю закрыть канал данных для прекращения передачи данных, так как это нарушит **принцип закрытия каналов**. Однако мы можем позволить получателю закрыть дополнительный сигнальный канал, чтобы уведомить отправителей о необходимости прекратить отправку данных.

```go
package main

import (
	"time"
	"math/rand"
	"sync"
	"log"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
		// stopCh is an additional signal channel.
		// Its sender is the receiver of channel
		// dataCh, and its receivers are the
		// senders of channel dataCh.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				// The try-receive operation is to try
				// to exit the goroutine as early as
				// possible. For this specified example,
				// it is not essential.
				select {
				case <- stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first
				// branch in the second select may be
				// still not selected for some loops if
				// the send to dataCh is also unblocked.
				// But this is acceptable for this
				// example, so the first select block
				// above can be omitted.
				select {
				case <- stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == Max-1 {
				// The receiver of channel dataCh is
				// also the sender of stopCh. It is
				// safe to close the stop channel here.
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()
}
```

Как упомянуто в комментариях, для дополнительного сигнального канала его отправителем является получатель данных из основного канала. Дополнительный сигнальный канал закрывается своим единственным отправителем, что соответствует принципу закрытия каналов.

В этом примере канал `dataCh` никогда не закрывается. Да, закрывать каналы вовсе не обязательно. Канал в конечном итоге будет удалён сборщиком мусора, если на него больше нет ссылок в горутинах, независимо от того, закрыт он или нет. Таким образом, в данном случае изящество закрытия канала заключается в том, чтобы вовсе не закрывать его.

### 3. M получателей, N отправителей, любой из них может сказать "давайте закончим игру", уведомив модератора о закрытии дополнительного сигнального канала

Это самая сложная ситуация. Мы **не можем позволить ни одному из получателей или отправителей закрыть канал данных**. Также **мы не можем позволить какому-либо из получателей закрыть дополнительный сигнальный канал для уведомления всех участников о завершении работы**, так как это нарушит **принцип закрытия каналов**. Однако мы можем **ввести роль модератора, который будет отвечать за закрытие сигнального канала**.

Одним из ключевых приёмов в следующем примере является использование попытки отправки (**try-send**), чтобы уведомить модератора о необходимости закрытия дополнительного сигнального канала.

```go
package main

import (
	"time"
	"math/rand"
	"sync"
	"log"
	"strconv"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
		// stopCh is an additional signal channel.
		// Its sender is the moderator goroutine shown
		// below, and its receivers are all senders
		// and receivers of dataCh.
	toStop := make(chan string, 1)
		// The channel toStop is used to notify the
		// moderator to close the additional signal
		// channel (stopCh). Its senders are any senders
		// and receivers of dataCh, and its receiver is
		// the moderator goroutine shown below.
		// It must be a buffered channel.

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					// Here, the try-send operation is
					// to notify the moderator to close
					// the additional signal channel.
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// The try-receive operation here is to
				// try to exit the sender goroutine as
				// early as possible. Try-receive and
				// try-send select blocks are specially
				// optimized by the standard Go
				// compiler, so they are very efficient.
				select {
				case <- stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first
				// branch in this select block might be
				// still not selected for some loops
				// (and for ever in theory) if the send
				// to dataCh is also non-blocking. If
				// this is unacceptable, then the above
				// try-receive operation is essential.
				select {
				case <- stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// Same as the sender goroutine, the
				// try-receive operation here is to
				// try to exit the receiver goroutine
				// as early as possible.
				select {
				case <- stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first
				// branch in this select block might be
				// still not selected for some loops
				// (and forever in theory) if the receive
				// from dataCh is also non-blocking. If
				// this is not acceptable, then the above
				// try-receive operation is essential.
				select {
				case <- stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						// Here, the same trick is
						// used to notify the moderator
						// to close the additional
						// signal channel.
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}
```

В этом примере **принцип закрытия канала всё ещё соблюдается**.

Обратите внимание, что размер буфера канала `toStop` равен одному. Это делается для того, чтобы предотвратить потерю первого уведомления, если оно будет отправлено до того, как горутина модератора будет готова принять уведомление от канала `toStop`.

Также можно установить размер буфера канала `toStop` равным сумме числа отправителей и получателей. В этом случае нам не нужно будет использовать блок select с попыткой отправки (try-send) для уведомления модератора.

```go
...
toStop := make(chan string, NumReceivers + NumSenders)
...
			value := rand.Intn(Max)
			if value == 0 {
				toStop <- "sender#" + id
				return
			}
...
				if value == Max-1 {
					toStop <- "receiver#" + id
					return
				}
...
```

### 4. Вариант ситуации "M получателей, один отправитель": запрос на закрытие делается третьей горутиной

Иногда требуется, чтобы сигнал о закрытии был отправлен третьей горутиной. В таких случаях можно использовать дополнительный сигнализирующий канал, чтобы уведомить отправителя о необходимости закрыть основной канал данных. Например:

```go
package main

import (
	"time"
	"math/rand"
	"sync"
	"log"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 100
	const NumThirdParties = 15

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)
	closing := make(chan struct{}) // signal channel
	closed := make(chan struct{})
	
	// The stop function can be called
	// multiple times safely.
	stop := func() {
		select {
		case closing<-struct{}{}:
			<-closed
		case <-closed:
		}
	}
	
	// some third-party goroutines
	for i := 0; i < NumThirdParties; i++ {
		go func() {
			r := 1 + rand.Intn(3)
			time.Sleep(time.Duration(r) * time.Second)
			stop()
		}()
	}

	// the sender
	go func() {
		defer func() {
			close(closed)
			close(dataCh)
		}()

		for {
			select{
			case <-closing: return
			default:
			}

			select{
			case <-closing: return
			case dataCh <- rand.Intn(Max):
			}
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}
```

Идея, использованная в функции stop, была взята [из комментария](https://groups.google.com/g/golang-nuts/c/lEKehHH7kZY/m/SRmCtXDZAAAJ) Роджера Пеппе.

## 5. Вариант ситуации с "N отправителями": необходимо закрыть канал данных, чтобы уведомить получателей о завершении передачи данных
В приведённых выше решениях для ситуаций с несколькими отправителями (`N-sender`) мы избегали закрытия каналов данных, чтобы соблюдать **принцип закрытия каналов**. Однако иногда требуется явно закрыть канал данных, чтобы получатели знали, что отправка данных завершена.

В таких случаях можно преобразовать ситуацию с N отправителями в ситуацию с одним отправителем, используя промежуточный канал. Этот промежуточный канал будет иметь только одного отправителя, и мы сможем закрыть его вместо закрытия исходного канала данных.

```go
package main

import (
	"time"
	"math/rand"
	"sync"
	"log"
	"strconv"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	log.SetFlags(0)

	// ...
	const Max = 1000000
	const NumReceivers = 10
	const NumSenders = 1000
	const NumThirdParties = 15

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)     // will be closed
	middleCh := make(chan int)   // will never be closed
	closing := make(chan string) // signal channel
	closed := make(chan struct{})

	var stoppedBy string

	// The stop function can be called
	// multiple times safely.
	stop := func(by string) {
		select {
		case closing <- by:
			<-closed
		case <-closed:
		}
	}
	
	// the middle layer
	go func() {
		exit := func(v int, needSend bool) {
			close(closed)
			if needSend {
				dataCh <- v
			}
			close(dataCh)
		}

		for {
			select {
			case stoppedBy = <-closing:
				exit(0, false)
				return
			case v := <- middleCh:
				select {
				case stoppedBy = <-closing:
					exit(v, true)
					return
				case dataCh <- v:
				}
			}
		}
	}()
	
	// some third-party goroutines
	for i := 0; i < NumThirdParties; i++ {
		go func(id string) {
			r := 1 + rand.Intn(3)
			time.Sleep(time.Duration(r) * time.Second)
			stop("3rd-party#" + id)
		}(strconv.Itoa(i))
	}

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					stop("sender#" + id)
					return
				}

				select {
				case <- closed:
					return
				default:
				}

				select {
				case <- closed:
					return
				case middleCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for range [NumReceivers]struct{}{} {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}
```

## Ещё варианты ситуаций?

Наверняка существуют и другие варианты ситуаций, но приведённые выше являются самыми распространёнными и базовыми. Используя каналы (а также другие техники конкурентного программирования) разумно, можно найти решение, которое будет соблюдать принцип закрытия каналов, для каждой возможной ситуации.

## Заключение
Не существует ситуаций, которые заставляют нарушать принцип закрытия каналов. Если вы столкнулись с такой ситуацией, пересмотрите ваш дизайн и перепишите код.

Программирование с использованием каналов в Go похоже на искусство.







