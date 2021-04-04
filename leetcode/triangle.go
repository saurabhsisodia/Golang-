package main
import "fmt"
func min(a,b int)int{
	if a>=b{
		return b
	}
	return a
}
func minimumPathSum(arr [][]int)int{
	var ans [][]int=make([][]int,len(arr))
	l:=len(arr)
	for i:=l-1;i>=0;i--{

		ans[i]=make([]int,len(arr[i]))

		for j:=0;j<len(arr[i]);j++{

			if i==l-1{
				ans[i][j]=arr[i][j]
			}else{
				if j+1<len(arr[i+1]){
					ans[i][j]=min(ans[i+1][j]+arr[i][j],ans[i+1][j+1]+arr[i][j])
				}else{
					ans[i][j]=ans[i+1][j]+arr[i][j]
				}
			}

		}

	}
	return ans[0][0]
}
func main(){
	var arr [][]int=[][]int{{-10}}
	fmt.Println(minimumPathSum(arr))
}