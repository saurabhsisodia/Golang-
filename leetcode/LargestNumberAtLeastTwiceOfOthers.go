package main
import "fmt"
func DominantIndex(slice []int)int{

	var f1,f2,max int
	var i,index int
	for i=0;i<len(slice);i++{

		if slice[i]>max{
			max=slice[i]
			index=i
		}
		if 2*slice[i] >f2{
			f1=f2
			f2=2*slice[i]
		} else if 2*slice[i]>f1{
			f1=2*slice[i]
		}

	}
	if max>=f1{
		return index
	}
	return -1
}
func main(){
	var slice []int=[]int{1}
	fmt.Println(DominantIndex(slice))
}