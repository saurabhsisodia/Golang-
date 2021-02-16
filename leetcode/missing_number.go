package main
import "fmt"
func Missing_Number(slice []int) int{

	xor:=0
	for i:=0;i<=len(slice);i++{
		xor^=i
	}
	for i:=0;i<len(slice);i++{
		xor^=slice[i]
	}
	return xor

}
func main(){
	var slice=[]int{0}
	fmt.Println(Missing_Number(slice))
}