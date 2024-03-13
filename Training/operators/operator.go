package main

import "fmt"

func main() {
	a := 10 // 10 - 1010
	b := 12 // 12 -1100
	c := a & b
	fmt.Println("bitwise", c)

	d := 10 // 1010
	e := 12 // 1100
	f := d | e
	fmt.Println("Bitwise OR", f)

	g := 10
	h := 12
	i := g ^ h
	fmt.Println("bitwise XOR", i)

	l := 10
	r := 2
	m := l << r
	fmt.Println("left operator", m)

	x := 10
	z := 2
	y := x >> z
	fmt.Println("right operator", y)

	modulos := 10 % 2
	fmt.Println("modulos", modulos)

	modulos1 := 8 % 3 // modulous basically gives us the remainder
	fmt.Println("modulous1", modulos1)

	// ternary operator

}
