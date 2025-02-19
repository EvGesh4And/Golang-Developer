package main

import "fmt"

func main() {

	var s []byte

	s = make([]byte, 5)

	fmt.Println(s)

	x := [3]string{"Лайка", "Белка", "Стрелка"}
	z := x[:] // a slice referencing the storage of x

	fmt.Println(z, len(z), cap(z))

	arr := [...]int{1, 2, 3, 4, 5, 6, 7}

	j := arr[2:]

	k := j[:cap(j)]

	fmt.Println(j, k)

	a := []string{"John", "Paul"}
	b := []string{"George", "Ringo", "Pete"}
	a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
	// a == []string{"John", "Paul", "George", "Ringo", "Pete"}

	mn := [...]int{1, 2, 3, 4, 5, 6, 7}

	l := mn[:3:4]

	fmt.Println(mn, l)
	l = append(l, 100)
	fmt.Println(mn, l)
	l = append(l, 1000)
	fmt.Println(mn, l)

	h := []int{1, 2, 3, 4, 5, 6}

	copy(h[0:4], h[1:5])

	fmt.Println(h)

	aa := []int{1, 2, 3, 4, 5, 6, 7}

	aa = append(aa[:5], aa[6:]...)

	fmt.Println(aa[:cap(aa)])

}
