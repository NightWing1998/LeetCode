// Link - https://leetcode.com/problems/add-digits/

func addDigits(num int) int {
    if num%9 == 0 && num != 0 {
        return 9;
    }
    return num%9;
}

// Time - 0ms(beats 100%)
// Memory - 2.2MB(beats 100%)

