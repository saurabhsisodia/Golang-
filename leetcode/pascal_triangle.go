package main
import "fmt"
func main(){

	n:=3
	arr:=make([][]int,n+1)
	i:=1
	for i<=n{
		arr[i]=make([]int,i)
		arr[i][0]=1
		if i==1{
			i++
			continue
		} else{
			for j:=0;j<i-2;j++{
				arr[i][j+1]=arr[i-1][j]+arr[i-1][j+1]
			}
			arr[i][i-1]=1
			i++
		}
	}
	fmt.Println(arr[1:])
}
