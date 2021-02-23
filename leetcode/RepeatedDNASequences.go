package main
import "fmt"

func findRepeatedDNASequence(s string)[]string{

	var check=make(map[string]int)
	var len int=len(s)
	var temp []string
	if len<=10{
		return temp
	}

	var i int =0
	var ans []string=make([]string,len-10+1)
	var k int=0
	for i=0;i<len-10+1;i++{
		if check[s[i:i+10]]==1{
			ans[k]=s[i:i+10]
			check[s[i:i+10]]=2
			k++
		} else if check[s[i:i+10]]==2{
			continue
		} else {
			check[s[i:i+10]]=1
		}
	}

	return ans[:k]
}

func main(){
	var s string="hello"
	fmt.Println(findRepeatedDNASequence(s))
}