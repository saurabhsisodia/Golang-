package main
import "fmt"

func SingleNumber(slice []int16)int16{
	var ans int16 =0
	var i int =0
	for i=0;i<len(slice);i++{

		ans=ans^slice[i]
	}
	return ans


}

func main(){
	var slice []int16=[]int16{1,1,2,2,3,3,-1,-1,2}
	fmt.Println(SingleNumber(slice))



}

