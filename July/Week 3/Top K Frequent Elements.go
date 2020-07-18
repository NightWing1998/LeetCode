// Link - https://leetcode.com/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) []int {
    if k == len(nums) {return nums;}
    countMap := make(map[int]int);
    for _, value := range nums {
        count,_ := countMap[value];
        countMap[value] = count+1;
    }
    bucket := make([][]int,len(nums)+1);
    for num,count := range countMap {
        bucket[count] = append(bucket[count],num);
    }
    res := make([]int,0,k);
    for i:=len(bucket)-1; i>=0; i-- {
        for j := 0; j < len(bucket[i]) && k > 0; j++ {
            res = append(res,bucket[i][j]);
            k--;
        }
    }
    return res;
}

// Time - 12ms(beats 97.44%)
// Memory - 5.4MB(beats 98.85%)
