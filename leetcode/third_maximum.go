package main
import (
	"fmt"
	"sort"
)

func ThirdMaximum(arr [] int ) int {
	// sort the given slice

	sort.Ints(arr)
	last:=arr[len(arr)-1]
	count:=2
    //fmt.Println(arr)
	for i:=len(arr)-1;i>=0;i--{
		if arr[i]!=last{
			last=arr[i]
			count--
		}
        if count==0{
			return last
		}
	}
	return arr[len(arr)-1]
}
func main(){

}