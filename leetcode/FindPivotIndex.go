package main
import "fmt"

func PivotIndex(slice []int)int{
	var PrefixSum []int=make([]int,len(slice))

	var i int
	var n int=len(slice)
	for i=0;i<n;i++{
		if i==0{
			PrefixSum[i]=slice[i]
		} else{
			PrefixSum[i]=PrefixSum[i-1]+slice[i]
		}
	}

	for i=0;i<n;i++{

		if i==0{
			if PrefixSum[n-1]-slice[0]==0{
				return 0
			}
		} else if i==n-1{
			if PrefixSum[n-1]-slice[n-1]==0{
				return n-1
			}
		} else{
			if PrefixSum[i-1]==PrefixSum[n-1]-PrefixSum[i]{
				return i
			}
		}

	}
	return -1
}

func main(){
	var slice []int=[]int{2}
	fmt.Println(PivotIndex(slice))
}