package main
import "fmt"

func PositionOfLargeGroup(s string)[][]int{

	var ans [][]int=make([][]int,len(s)/3)
	var count,i,j,k int

	for j=0;j<len(s);j++{

		if j==0{
			continue
		}
		if s[j]!=s[j-1]{
			if j-i>=3{
				ans[k]=[]int{i,j-1}
				k++
				count++
			}
			i=j
		}
	}
	if j-i>=3{
		ans[k]=[]int{i,j-1}
		count++
	}
    return ans[:count]
}

func main(){
	fmt.Println(PositionOfLargeGroup("aaaa"))
}
