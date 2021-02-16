package main
import (
	"fmt"
)
func main(){

	var arr []int=[]int{3,2,3}
	var x int = len(arr)/2

	mp:=make(map[int]int)
	for _,ele:=range arr{
		mp[ele]+=1
	}
	for k,v:=range mp{
		if v>x{
			fmt.Println(k)
			break
		}

		
	}
}