package main
import "fmt"

func MonotonicArray(slice []int)bool{
	var inc,dec bool=true,true
	var i int
	for i=0;i<len(slice);i++{

		if i==0{
			continue
		}

		if slice[i]<slice[i-1]{
			inc=false
		}
		if slice[i]>slice[i-1]{
			dec=false
		}
	}

	if inc==false && dec==false{
		return inc
	}
	return true
}
func main(){
	fmt.Println(MonotonicArray([]int{1,1,1}))
}