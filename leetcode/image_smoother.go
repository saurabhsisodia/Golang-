package main
import "fmt"
func ok(i int,j int,r int,c int) bool{
    
    if i<r && i>=0 && j<c && j>=0{
        return true
	}
	return false
}

func imageSmoother(mat [][]int) [][]int {
    
    var i,j int
	var r=len(mat)
    var c=len(mat[0])
	ans:= make([][]int,r)
	for i=0;i<r;i++{
		ans[i]=make([]int,c)
		for j=0;j<c;j++{
			
			var total,non_zero int
			non_zero+=mat[i][j]
			total++
			if ok(i,j+1,r,c){
				non_zero+=mat[i][j+1]
				total++
			}
			if ok(i,j-1,r,c){
				non_zero+=mat[i][j-1]
				total++
			}
			if ok(i-1,j,r,c){
				non_zero+=mat[i-1][j]
				total++
			}
			if ok(i+1,j,r,c){
				non_zero+=mat[i+1][j]
				total++
			}
			if ok(i-1,j+1,r,c){
				non_zero+=mat[i-1][j+1]
				total++
			}
			if ok(i+1,j+1,r,c){
				non_zero+=mat[i+1][j+1]
				total++
			}
			if ok(i-1,j-1,r,c){
				non_zero+=mat[i-1][j-1]
				total++
			}
			if ok(i+1,j-1,r,c){
				non_zero+=mat[i+1][j-1]
				total++
			}
		ans[i][j]=non_zero/total

		}
	}
	return ans
    
}
func main(){
	var mat [][]int = [][]int{{0,0,0},{0,0,0},{0,0,0}}
	fmt.Println(ImageSmoother(mat))
}