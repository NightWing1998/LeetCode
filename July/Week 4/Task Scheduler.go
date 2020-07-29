// Link - https://leetcode.com/problems/task-scheduler/

func max(a,b int) int {
    if a > b { 
        return a;
    }
    return b;
     // if a < b {return b;}return a; ---> less optimal for max = diff is 8ms
}

func leastInterval(tasks []byte, n int) int {
    var count = make([]int,26);
    maxCount := 0;
    for _,task := range tasks {
        count[task-'A']++;
        if maxCount < count[task-'A'] {
            maxCount = count[task-'A'];
        }
    }
    maxTaskCount := 0;
    for i := 0; i < 26; i++ {
        if count[i] == maxCount {maxTaskCount++}
    }
    return max(len(tasks),(maxCount-1)*(n+1) + maxTaskCount);
}

// Time - 4ms(beats 97.50%)
// Memory - 6MB(beats 50%)


