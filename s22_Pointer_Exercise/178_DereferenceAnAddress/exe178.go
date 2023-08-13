package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Println(*a)
	fmt.Printf("%T %v \n", a, a)
	fmt.Println(*a)
	fmt.Printf("%T %v \n", b, b)
	fmt.Println(*c)
	fmt.Printf("%T %v \n", c, c)
	fmt.Println(*d)
	fmt.Printf("%T %v \n", d, d)

}
