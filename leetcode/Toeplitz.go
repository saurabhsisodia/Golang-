package main
import "fmt"
func isToeplitzMatrix(mat [][]int) bool{

	var mp map[int]string=make(map[int]string)
	
	var r,c,val int
	var row []int
	for r,row=range mat{
		for c,val=range row{
			if mp[r-c]==""{
				mp[r-c]=string(val)
			} else if mp[r-c]!=string(val){
				return false
			}
		}
	}
	return true
}
func main(){
	var mat [][]int =[][]int{{0,33,98},{34,22,33}}
	fmt.Println(isToeplitzMatrix(mat))

}
