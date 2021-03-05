package main
import "fmt"

func main(){
	var arr []int=[]int{1,2,3,4,5}
	var ans []int=make([]int,len(arr))
	ans[0]=1
	for i:=1;i<len(arr);i++{
		ans[i]=ans[i-1]*arr[i-1]
	}

	r:=1
	for i:=len(arr)-1;i>=0;i--{
		ans[i]=ans[i]*r
		r*=arr[i]
	}
	fmt.Println(ans)
}