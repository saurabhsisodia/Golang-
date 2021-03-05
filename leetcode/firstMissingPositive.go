package main
import (
	"fmt"
	"math"
)

func partition(arr []int)int{
	size:=len(arr)
	var left,right int
	left=-1
	for right<size{
		if arr[right]>0{
			left++
			arr[left],arr[right]=arr[right],arr[left]
		}
		right++
	}
	return left
}
func main(){
	var arr []int=[]int{3,4,-1,1}
	leftpositive:=partition(arr)+1
	for i:=0;i<leftpositive;i++{
		x:=int(math.Abs(float64(arr[i])))-1
		if x>=leftpositive{
			continue
		}
		if arr[x]>0{
			arr[x]=-1*arr[x]
		}

	}
	ok:=false
	for i:=0;i<leftpositive;i++{
		if arr[i]>0{
			fmt.Println(i+1)
			ok=true
			break
		}
	}
	if ok==false{
	fmt.Println(leftpositive+1)
	}
}