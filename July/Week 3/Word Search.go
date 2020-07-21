// Link - https://leetcode.com/problems/word-search/

func exist(board [][]byte, word string) bool {
    for i := 0; i < len(board);i++ {
        for j := 0; j < len(board[i]); j++ {
            if board[i][j] == word[0] &&
                    checkCellWithAlpha(i,j,len(board),len(board[i]),0,board,word){
                return true;
            }
        }
    }
    return false;
}

func checkCellWithAlpha(i,j,n,m,k int,board [][]byte,word string) bool {
    if k >= len(word) {return true;}
    if i < 0 || j < 0 || i >= n || j >= m {return false;}
    if board[i][j] != word[k] {
        return false;
    }
    temp := board[i][j];
    board[i][j] = '$';
    res :=  checkCellWithAlpha(i,j+1,n,m,k+1,board,word) ||
            checkCellWithAlpha(i+1,j,n,m,k+1,board,word) || 
            checkCellWithAlpha(i,j-1,n,m,k+1,board,word) || 
            checkCellWithAlpha(i-1,j,n,m,k+1,board,word);
    board[i][j] = temp;
    return res;
}

// Time - 4ms(beats 97.17%)
// Memory - 3.6MB(beats 61.43%)

