package main

import "fmt"

type customErr struct {
	msg string
}

func (ce customErr) Error() string {
	return ce.msg
}

func main() {
	c1 := customErr{"Error need more coffee"}
	foo(c1)
}
func foo(e error) {
	fmt.Println(e)
}
