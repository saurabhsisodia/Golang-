package main
import "fmt"
func temp(arr []int,mp map[int]int,start int)bool{
	if start<0 || start>=len(arr){
		return false
	}
	if arr[start]==0{
		return true
	}
	if mp[start]!=0{
		return false
	}

	mp[start]=1
	return temp(arr,mp,start+arr[start]) || temp(arr,mp,start-arr[start])
}
func canReach(arr []int,start int)bool{
	var mp map[int]int=make(map[int]int)
	return temp(arr,mp,start)

}

func main(){
	var arr []int=[]int{3,0,2,1,2}
	fmt.Println(canReach(arr,2))
}