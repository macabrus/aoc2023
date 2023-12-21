package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "bytes"
)

func IsDigit(c rune) bool {
    return c >= '0' && c <= '9'
}

func IsOnBoard(board [][]rune, x int, y int) bool {
    return !(x < 0 || x >= len(board) || y < 0 || y >= len(board[0]))
}

func IsPartNum(board [][]rune, i int, j int) bool {
    directions := [][]int {
        {-1, -1},
        {-1,  0},
        {-1,  1},
        { 0, -1},
        // { 0,  0},
        { 0,  1},
        { 1, -1},
        { 1,  0},
        { 1,  1},
    }
    isPartNum := false
    for k := j; IsOnBoard(board, i, k) && IsDigit(board[i][k]); k++ {
        for _, direction := range directions {
            x := i + direction[0]
            y := k + direction[1]
            fmt.Printf("i=%d, j=%d, x=%d, y=%d\n", i, k, x, y)
            if !IsOnBoard(board, x, y) || IsDigit(board[x][y]) {
                continue
            }
            if board[x][y] != '.' {
                isPartNum = true
            }
        }
    }
    return isPartNum
}

func GetPartNum(board [][]rune, i int, j int) string {
    buff := bytes.NewBufferString("")
    for k := j; IsOnBoard(board, i, k) && IsDigit(board[i][k]); k++ {
        buff.WriteRune(board[i][k])
    }
    return buff.String()
}

func main() {
    board := make([][]rune, 0)
    for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
        board = append(board, []rune(sc.Text()))
    }
    fmt.Printf("%v\n", board)
    sum, i, j := 0, 0, 0
    for ; i < len(board); i++ {
        for ; j < len(board[0]); {
            numStr := GetPartNum(board, i, j)
            fmt.Println(numStr)
            if IsPartNum(board, i, j) {
                num, _ := strconv.Atoi(numStr)
                sum += num
            }
            if len(numStr) != 0 {
                j += len(numStr)
                continue
            }
            j++
        }
        j = 0
    }
    fmt.Printf("%d\n", sum)
}
