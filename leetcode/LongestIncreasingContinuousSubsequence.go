package main
import "fmt"

func LongIncConSubsequence(slice []int) int{

	var i,j int
	var ans int

	if len(slice)==0{
		return 0
	}
	for j=1;j<len(slice);j++{

		if slice[j]<=slice[j-1]{
			if j-i>ans{
				ans=j-i
			}
			i=j
		}
		
	}
	if j-i>ans{
		ans=j-i
	}
	return ans
}

func main(){
	var slice []int=[]int{5,4,3,2,1}
	fmt.Println(LongIncConSubsequence(slice))
}