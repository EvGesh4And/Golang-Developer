# Атомарные операции, предоставляемые в стандартном пакете `sync/atomic`

[Статья](https://go101.org/article/concurrent-atomic-operation.html)

Атомарные операции являются более примитивными, чем другие методы синхронизации. Они не используют блокировки и, как правило, реализуются непосредственно на аппаратном уровне. Фактически, они часто применяются для реализации других методов синхронизации.

Обратите внимание, что многие примеры ниже не являются конкурентными программами. Они предназначены исключительно для демонстрации и объяснения принципов работы с атомарными функциями, предоставляемыми стандартным пакетом `sync/atomic`.

## Обзор атомарных операций, доступных до Go 1.19

Стандартный пакет `sync/atomic` предоставляет следующие пять атомарных функций для целочисленного типа `T`, где `T` должен быть одним из: `int32`, `int64`, `uint32`, `uint64` и `uintptr`.

```go
func AddT(addr *T, delta T)(new T)
func LoadT(addr *T) (val T)
func StoreT(addr *T, val T)
func SwapT(addr *T, new T) (old T)
func CompareAndSwapT(addr *T, old, new T) (swapped bool)
```

Например, для типа `int32` предоставляются следующие пять функций:

```go
func AddInt32(addr *int32, delta int32)(new int32)
func LoadInt32(addr *int32) (val int32)
func StoreInt32(addr *int32, val int32)
func SwapInt32(addr *int32, new int32) (old int32)
func CompareAndSwapInt32(addr *int32,
				old, new int32) (swapped bool)
```

Следующие четыре атомарные функции предоставляются для (безопасных) типов указателей. Когда эти функции были добавлены в стандартную библиотеку, Go ещё не поддерживал пользовательские дженерики, поэтому они были реализованы через [небезопасный указатель](https://go101.org/article/unsafe.html) `unsafe.Pointer` (Go-аналог void* в C).

```go
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer,
				) (old unsafe.Pointer)
func CompareAndSwapPointer(addr *unsafe.Pointer,
				old, new unsafe.Pointer) (swapped bool)
```

Функции `AddPointer` для указателей не существует, так как (безопасные) указатели в Go не поддерживают арифметические операции.

Стандартный пакет `sync/atomic` также предоставляет тип `Value`, соответствующий указательный тип `*Value` имеет четыре метода (перечисленные ниже, два последних были добавлены в Go 1.17). Эти методы позволяют выполнять атомарные операции с значениями любого типа.

```go
func (*Value) Load() (x interface{})
func (*Value) Store(x interface{})
func (*Value) Swap(new interface{}) (old interface{})
func (*Value) CompareAndSwap(old, new interface{}) (swapped bool)
```

## Обзор новых атомарных операций, добавленных в Go 1.19
В Go 1.19 были введены несколько типов, каждый из которых содержит набор методов для атомарных операций. Эти методы обеспечивают тот же эффект, что и функции уровня пакета, перечисленные в предыдущем разделе.

Среди этих типов:

`Int32`, `Int64`, `Uint32`, `Uint64` и `Uintptr` предназначены для атомарных операций с целыми числами.
Методы типа `atomic.Int32` приведены ниже. Методы других четырех типов реализованы аналогичным образом.


```go
func (*Int32) Add(delta int32) (new int32)
func (*Int32) Load() int32
func (*Int32) Store(val int32)
func (*Int32) Swap(new int32) (old int32)
func (*Int32) CompareAndSwap(old, new int32) (swapped bool)
```

Начиная с Go 1.18, язык поддерживает обобщённые (generic) типы. Некоторые стандартные пакеты начали использовать generics с версии Go 1.19, и `sync/atomic` — один из таких пакетов.

В Go 1.19 в этом пакете был введён обобщённый тип `Pointer[T any]`, который позволяет работать с указателями атомарно. Его методы приведены ниже.

```go
(*Pointer[T]) Load() *T
(*Pointer[T]) Store(val *T)
(*Pointer[T]) Swap(new *T) (old *T)
(*Pointer[T]) CompareAndSwap(old, new *T) (swapped bool)
```

В Go 1.19 также был введён тип `Bool` в пакете `sync/atomic`, который позволяет выполнять атомарные операции с булевыми значениями. Этот тип обеспечивает потокобезопасное чтение и запись true/false, что полезно при организации синхронизации в многопоточных программах.

## Атомарные операции для целых чисел

Оставшаяся часть этой статьи покажет несколько примеров использования атомарных операций, предоставленных в Go.

Следующий пример демонстрирует, как выполнить атомарную операцию `Add` над значением типа `int32` с использованием функции `AddInt32`. В этом примере главная горутина создает 1000 новых конкурирующих горутин. Каждая из этих горутин увеличивает целочисленное значение `n` на единицу. Атомарные операции гарантируют отсутствие условий гонки среди этих горутин. В конечном итоге будет выведено значение 1000.

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var n int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&n, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(atomic.LoadInt32(&n)) // 1000
}
```

Если заменить выражение `atomic.AddInt32(&n, 1)` на `n++`, то вывод может быть не 1000.

Следующий код переиспользует тип `atomic.Int32` и его методы (начиная с Go 1.19). Этот код выглядит немного аккуратнее.

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var n atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			n.Add(1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(n.Load()) // 1000
}
```

Функции/методы `StoreT` и `LoadT` атомарного типа часто используются для реализации методов установки и получения значения (соответствующего указателя на тип), если значения типа необходимо использовать одновременно в нескольких горутинах. Например, версия функции:

```go
type Page struct {
	views uint32
}

func (page *Page) SetViews(n uint32) {
	atomic.StoreUint32(&page.views, n)
}

func (page *Page) Views() uint32 {
	return atomic.LoadUint32(&page.views)
}
```

А вот версия с типом и методами (с Go 1.19):

```go
type Page struct {
	views atomic.Uint32
}

func (page *Page) SetViews(n uint32) {
	page.views.Store(n)
}

func (page *Page) Views() uint32 {
	return page.views.Load()
}
```

Для знакового целочисленного типа T (`int32` или `int64`) вторым аргументом вызова функции `AddT` может быть отрицательное значение, чтобы выполнить атомарную операцию уменьшения. Но как выполнить атомарные операции уменьшения для значений беззнаковых типов `T`, таких как `uint32`, `uint64` и `uintptr`? Для вторых аргументов беззнаковых типов существуют два случая.

1. Для беззнаковой переменной `v` типа `T` значение `-v` является допустимым в Go. Таким образом, мы можем передать `-v` как второй аргумент вызова `AddT`.
2. Для положительной константы целого числа `c` значение `-c` является недопустимым вторым аргументом вызова `AddT` (где `T` обозначает беззнаковый целочисленный тип). В этом случае мы можем использовать `^T(c-1)` в качестве второго аргумента.

Этот трюк `^T(v-1)` также работает для беззнаковой переменной `v`, но `^T(v-1)` менее эффективен, чем `T(-v)`.

В случае с трюком `^T(c-1)`, если `c` является типизированным значением и его тип точно совпадает с `T`, то форму можно сократить до `^(c-1)`.

Пример:

```go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var (
		n uint64 = 97
		m uint64 = 1
		k int    = 2
	)
	const (
		a        = 3
		b uint64 = 4
		c uint32 = 5
		d int    = 6
	)

	show := fmt.Println
	atomic.AddUint64(&n, -m)
	show(n) // 96 (97 - 1)
	atomic.AddUint64(&n, -uint64(k))
	show(n) // 94 (96 - 2)
	atomic.AddUint64(&n, ^uint64(a - 1))
	show(n) // 91 (94 - 3)
	atomic.AddUint64(&n, ^(b - 1))
	show(n) // 87 (91 - 4)
	atomic.AddUint64(&n, ^uint64(c - 1))
	show(n) // 82 (87 - 5)
	atomic.AddUint64(&n, ^uint64(d - 1))
	show(n) // 76 (82 - 6)
	x := b; atomic.AddUint64(&n, -x)
	show(n) // 72 (76 - 4)
	atomic.AddUint64(&n, ^(m - 1))
	show(n) // 71 (72 - 1)
	atomic.AddUint64(&n, ^uint64(k - 1))
	show(n) // 69 (71 - 2)
}
```