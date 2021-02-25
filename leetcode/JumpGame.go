package main
import "fmt"

func CanJump(slice []int)bool{
	if len(slice)==1{
		return true
	}
	var last int=len(slice)-1
	var i int
	for i=len(slice)-2;i>0;i--{
		if i+slice[i]>=last{
			last=i
		}
	}
	if slice[0]>=last{
		return true
	}
	return false
}
func main(){
	fmt.Println(CanJump([]int{1,0}))
}