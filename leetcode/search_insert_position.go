package main
import "fmt"
func BS(arr []int,target int) int {
	var l int = 0
	var h int = len(arr)-1
	var mid int = (l+h)/2
	for l<=h{
		
		mid=(l+h)/2
		if arr[mid]==target{
			return mid
		}else if arr[mid]>target{
			h=mid-1
			continue
		} else{
		l=mid+1
		}
	}
	if arr[mid]<target{
		return mid+1
	}
	return mid
}

func main(){
	var arr []int = []int{2,4,5,6,7,9,12}
	fmt.Println(BS(arr,0))
}