func max(a int ,b int) int{
    if a>=b{
        return a
    }
    return b
}
func trap(height []int) int {
    
    
    var l,leftmax,rightmax,ans int
    var r int =len(height)-1
    for l<=r{
        
        if height[l]<=height[r]{
            
            if leftmax>height[l]{
                ans+=leftmax-height[l]
            }
            leftmax=max(leftmax,height[l])
            l++
        } else{
            
            if rightmax>height[r]{
                ans+=rightmax-height[r]
            }
            rightmax=max(rightmax,height[r])
            r--
        }
    }
    return ans
}