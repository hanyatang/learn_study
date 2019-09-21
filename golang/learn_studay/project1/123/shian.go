package main 

import (
	"fmt"
)

func nubmer() func() int {
	var x int
	return func() int {
		x++
		return x *x
	}
}

func main(){
	f := nubmer()
		fmt.Println(f())
		fmt.Println(f())
		fmt.Println(f())
		fmt.Println(f())
		fmt.Println(f())
}