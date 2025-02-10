# Пакет `unsafe`

Пакет `unsafe` в Go даёт доступ к низкоуровневым возможностям работы с памятью, которые обычно скрыты в языке.

Его можно использовать для:

- Определения размера и выравнивания структур
- Получения указателей на поля структур
- Преобразования типов без копирования

Но осторожно — использование `unsafe` может привести к некорректной работе программы, если не учитывать специфику платформы и внутренние детали реализации Go.

## Основные функции `unsafe`

1. `unsafe.Sizeof()`

    Возвращает размер переменной или типа **в байтах**.

    ```go
    package main

    import (
        "fmt"
        "unsafe"
    )

    func main() {
        var x int64
        fmt.Println(unsafe.Sizeof(x)) // 8 байт

        var y struct {
            a int8
            b int64
            c int8
        }
        fmt.Println(unsafe.Sizeof(y)) // 24 байта (из-за выравнивания)
    }
    ```
2. `unsafe.Alignof()`

   Показывает **требуемое выравнивание** для типа (обычно совпадает с размером самого большого примитивного типа в структуре).

   ```go
    package main

    import (
        "fmt"
        "unsafe"
    )

    func main() {
        fmt.Println(unsafe.Alignof(int8(0)))    // 1
        fmt.Println(unsafe.Alignof(int64(0)))   // 8
        fmt.Println(unsafe.Alignof(float64(0))) // 8

        type Example struct {
            a int8  // 1 байт
            b int64 // 8 байт (должен быть выровнен по 8)
            c int8  // 1 байт
        }
        fmt.Println(unsafe.Alignof(Example{})) // 8 (по самому большому типу)
    }
   ```

3. `unsafe.Offsetof()`

    Возвращает **смещение** поля внутри структуры (в байтах от начала структуры).

    ```go
    package main

    import (
        "fmt"
        "unsafe"
    )

    type Example struct {
        a int8  // 0-й байт
        b int64 // 8-й байт (из-за выравнивания)
        c int8  // 16-й байт
    }

    func main() {
        fmt.Println(unsafe.Offsetof(Example{}.a)) // 0
        fmt.Println(unsafe.Offsetof(Example{}.b)) // 8
        fmt.Println(unsafe.Offsetof(Example{}.c)) // 16
    }
    ```
4. 
    `unsafe.Pointer` — универсальный указатель, который можно привести к любому типу.

    📌 Пример: преобразование `*int` в `*float64`:

    ```go
    package main

    import (
        "fmt"
        "unsafe"
    )

    func main() {
        var i int = 42
        ptr := unsafe.Pointer(&i)   // Приведение *int → unsafe.Pointer
        fptr := (*float64)(ptr)     // Приведение unsafe.Pointer → *float64

        fmt.Println(*fptr)          // Читаем память как float64 (неожиданные значения!)
    }
    ```

5. `uintptr`

    `uintptr` — это целочисленный аналог указателя. Можно использовать для арифметики с адресами.

    📌 **Пример: получение адреса и сдвиг на 8 байт:**

    ```go
    package main

    import (
        "fmt"
        "unsafe"
    )

    func main() {
        var arr = [2]int{42, 99}
        ptr := unsafe.Pointer(&arr[0]) // Получаем указатель на arr[0]
        ptr = unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(arr[0])) // Смещаемся на 8 байт (int64)

        fmt.Println(*(*int)(ptr))       // Читаем arr[1] (99)
    }
    ```


❗ Когда стоит использовать unsafe?

- ✔ Для оптимизации работы с памятью
- ✔ Для работы с системными API (например, взаимодействие с C-кодом через cgo)
- ✔ Для высокопроизводительных структур данных

❌ Когда НЕ стоит использовать unsafe?
- ❌ В обычных приложениях (может сломаться на другой архитектуре)
- ❌ Когда есть безопасные альтернативы (например, reflect)
- ❌ Если ты не полностью понимаешь, как работает память