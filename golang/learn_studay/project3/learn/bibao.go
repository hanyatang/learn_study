package main

import(
	"fmt"
	"strings"
)

func main(){
	f := makesuffix(".fdp")
	fmt.Println("文件名处理后=",f("winter"))
}
func makesuffix(suffix string) func (string) string {
	return func  (name string) string {
		if !strings.HasSuffix(name , suffix) {
			return name + suffix
		}
		return name
	}
}