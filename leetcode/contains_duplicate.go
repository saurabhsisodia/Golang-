package main
import "fmt"
func main(){
	var arr []int=[]int{1,2,3,4}

	for i:=0;i<len(arr);i++{
		var x int
		if arr[i]<0{
			x=-1*arr[i]
		}else{
			x=arr[i]
		}

		if arr[x-1]<0{
			fmt.Println(true)
			break
		}else{
			arr[x-1]=-1*arr[i]
		}
	}
	fmt.Println(false)

}
