package main
import "fmt"

func FindShortestSubarray(slice []int)int{

	last:=make(map[int]int)
	first:=make(map[int]int)
	freq:=make(map[int]int)
	var i int
	for i=0;i<len(slice);i++{
		if first[slice[i]]==0{
			first[slice[i]]=i+1
		}
		last[slice[i]]=i+1
		freq[slice[i]]++
	}
	var ans=len(slice)
	var degree int
	var k,v int
	for _,v=range freq{
		if v>degree{
			degree=v
		}
	}
	for k,v=range freq{
		if v==degree{
			if last[k]-first[k]+1<ans{
				ans=last[k]-first[k]+1
			}
		}
	}
	return ans
}


func main(){

	var slice []int=[]int{1,2,3,4,5}
	fmt.Println(FindShortestSubarray(slice))


}