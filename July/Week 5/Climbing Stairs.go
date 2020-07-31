// Link - https://leetcode.com/problems/climbing-stairs/

func climbStairs(n int) int {
    if n <= 2 {
        return n;
    }
//     waysToClimbDp := make([]int,n+1);
//     waysToClimbDp[1],waysToClimbDp[2] = 1,2;
    
//     for i:= 3;i <=n;i++ {
//         waysToClimbDp[i] = waysToClimbDp[i-2] + waysToClimbDp[i-1];
//     }
//     return waysToClimbDp[n];
//     ---------requires too much memory--------
    res := 2;
    prev := 1;
    for i := 3; i <= n; i++ {
        prev,res = res,res+prev;
    }
    return res;
}

// Time - 0ms(beats 100%)
// Memory - 2MB(beats 50%)

