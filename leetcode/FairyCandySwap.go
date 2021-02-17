package main
import "fmt"

func FairCandySwap(a []int,b []int)[]int{

	var mpA map[int]int=make(map[int]int)
	var sumA,sumB,i int
	var v int

	for i=0;i<len(a);i++{

		sumA+=a[i]
	}
	for i=0;i<len(b);i++{
		sumB+=b[i]
	}
	for i=0;i<len(a);i++{
		v=sumA-2*a[i]
		mpA[v]=a[i]
	}
	for i=0;i<len(b);i++{

		v=sumB-2*b[i]
		if mpA[v]!=0{
			return []int{mpA[v],b[i]}
		}
	}
	return []int{1,2}
}
func main(){
	fmt.Println(FairCandySwap([]int{1,1},[]int{2,2}))
}