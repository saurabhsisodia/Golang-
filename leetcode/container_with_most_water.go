func min(a int, b int) int{
    if a>b{
        return b
    }
    return a
}
func maxArea(height []int) int {
    
    var l,ans int
    var r int=len(height)-1
    
    for l<r{
        
        temp:=(r-l)*min(height[l],height[r])
        ans=-min(-ans,-temp)
        
        if height[l]<height[r]{
            l++
        } else {
            r--
        }
    }
    return ans
    
}