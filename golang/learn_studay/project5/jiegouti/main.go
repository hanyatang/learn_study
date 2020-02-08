package main

import (
	"fmt"
)

type People struct {
	Name string
	Weight int
}

//func main() {
	//peo :=People{"小明",19}//值传递结构体
	//fmt.Println(peo)
	//peo1 := &People{"小明",18}//结构体的指针
	//fmt.Println(peo1)
//}
func (p *People) run() {
	fmt.Println(p.Name,"正在跑步，当前体重是：",p.Weight)
	p.Weight-=1
}

func main() {
	peo :=People{"张三",100}
	peo.run()
	fmt.Println("跑步后的体重是：",peo.Weight)
}
