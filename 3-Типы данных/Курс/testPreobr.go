package main

func main() {
	var i int32 = 42
	var j uint32 = i         // ошибка
	var k uint32 = uint32(i) // верно
	var n int64 = i          // ошибка!
	var m int64 = int64(i)   // верно
	var r rune = i           // верно

	var ff float64 = 546.45645
	var in int64 = int64(ff)
}
