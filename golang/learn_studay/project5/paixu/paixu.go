package main
 
import (
	"fmt"
	"sort"
	

)



func main(){
	list :=[]int{2,5,63,4,7}
	for i:=0 ;i<len(list)-1;i++{
		for j:=0;j<len(list)-1-i;j++{
			if list[j]>list[j+1]{
				list[j],list[j+1] = list[j+1],list[j]
			}
		}
	}
	fmt.Println(list)
	//sort.Ints(list)
	sort.Sort(sort.IntSlice(list))
	fmt.Println(list)
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	fmt.Println(list)
	

}