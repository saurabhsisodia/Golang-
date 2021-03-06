import "math"
func maxArea(height []int) int {
    
    var l,ans int
    var r int=len(height)-1
    
    for l<r{
        
        temp:=(r-l)*int(math.Min(float64(height[r]),float64(height[l])))
        ans=int(math.Max(float64(ans),float64(temp)))
        
        if height[l]<height[r]{
            l++
        } else {
            r--
        }
    }
    return ans
    
}