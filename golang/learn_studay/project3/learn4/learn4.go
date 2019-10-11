package main

import (
	"fmt"
)
func BinaryFind(arr *[6]int,leftindex int,rightindex int,findval int) {
	if leftindex > rightindex{
		fmt.Println("找不到")
		return
	}
	middle := (leftindex + rightindex) / 2
	
	if (*arr)[middle] > findval {
		BinaryFind(arr,leftindex,middle - 1,findval)
	}else if (*arr)[middle] < findval {
		BinaryFind(arr,middle + 1,rightindex,findval)
	}else {
		fmt.Printf("找到了，下标是%v",middle)
	}
}
func main () {
	arr :=[6]int{24,8,34,133,64,412}
	shunxu(&arr)
	BinaryFind(&arr,0,len(arr),34)
}
func shunxu(arr *[6]int) {
	fmt.Println("排序前=",arr)
	temp :=0
	for i:=0;i<len(*arr) - 1 ;i++{
		for j:=0;j<len(*arr)-1;j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j +1]
				(*arr)[j+ 1] = temp
			}
		}
	}
	fmt.Println("排序后=",arr)
}