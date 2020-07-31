// Link - https://leetcode.com/problems/word-break-ii/

func wordBreak(s string, wordDict []string) []string {
    wordMap := make(map[string][]string);
    return wordSearch(s,wordDict,wordMap);
}

func wordSearch(s string, wordDict []string, wordMap map[string][]string) []string {
    if list,exists := wordMap[s]; exists {
        return list;
    }
    list := make([]string,0);
    for _,word := range wordDict {
        if word == s {
            list = append(list,s);
            break;
        }
    }
    for i := 1; i < len(s); i++ {
        leftSubStr := string(s[0:i]);
        for _, word := range wordDict {
            if word == leftSubStr {
                for _,right := range wordSearch(string(s[i:]),wordDict,wordMap) {
                    list = append(list,leftSubStr + " " + right);
                }
            }
        }
    }
    wordMap[s] = list;
    return list;
}

// Time - 12ms(beats 10.53%)
// Memory - 2.8MB(-)
