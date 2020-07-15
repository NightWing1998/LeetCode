// Link - https://leetcode.com/problems/reverse-words-in-a-string/

import (
    "bytes"
    "strings"
)

func reverseWords(s string) string {
    var res bytes.Buffer;
    for i:=len(s)-1;i >= 0; i-- {
        for ;i >= 0 && s[i] == ' ';i-- {}
        j := i;
        for ;j >= 0 && s[j] != ' ';j-- {}
        res.WriteString(s[j+1:i+1]);
        res.WriteString(" ");
        i = j;
    }
    return strings.Trim(res.String()," ");
}

// Time - 4ms(beats 63.43%)
// Memory - 3.7MB(beats 63.77%)

