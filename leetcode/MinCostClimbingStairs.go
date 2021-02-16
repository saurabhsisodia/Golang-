package main
import "fmt"
func minCostClimbingStairs(cost []int)int{
	var n int =len(cost)
	var f1,f2 int
	var i int
	for i=n-1;i>=0;i--{
		var temp int=f2
		if f1<=f2{
			temp=f1
		}
		f1,f2=cost[i]+temp,f1
	}
	if f1<=f2{
		return f1
	}
	return f2
}
func main(){
	var cost []int =[]int {0,2,2,1}
	fmt.Println(minCostClimbingStairs(cost))
}