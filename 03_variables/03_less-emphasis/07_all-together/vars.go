package main

import "fmt"

// package scope variables
var a = "the quick brown fox stores in variable a"
var b, c string = "stores in b", "stores in c"
var d string

func main() {
	d = "this also package scope variable"
	var e = 18 // function scope variable
	f := 19
	g := "this is function scope variable"
	h, i := "another function", "scope variable"
	j, k, l, m := 17.03, false, true, 'y'
	n := "n"
	o := `o`

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)
}
