package main

import(
	"fmt"
	
)
func fln(n int)([]uint64){
	flb :=make([]uint64,n)
	flb[0] = 1
	flb[1] = 1
	for i:=2 ;i<n;i++{
		flb[i] = flb[i-1] + flb[i-2]	
	}
	return flb
}
func main() {
	
	abc := fln(10)
	fmt.Println(abc)
	
}