package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "bytes"
)

type Board struct {
    Board [][]rune
    Used  [][]bool
}

func IsDigit(c rune) bool {
    return c >= '0' && c <= '9'
}

func IsOnBoard(b Board, x int, y int) bool {
    return !(x < 0 || x >= len(b.Board) || y < 0 || y >= len(b.Board[0]))
}

var directions = [][]int {
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


func IsPartNum(b Board, i int, j int) bool {
    isPartNum := false
    for k := j; IsOnBoard(b, i, k) && IsDigit(b.Board[i][k]); k++ {
        for _, d := range directions {
            x := i + d[0]
            y := k + d[1]
            if !IsOnBoard(b, x, y) || IsDigit(b.Board[x][y]) {
                continue
            }
            if b.Board[x][y] != '.' {
                isPartNum = true
            }
        }
    }
    return isPartNum
}


func GetNumLen(b Board, i int, j int) int {
    i, j = GetStart(b, i, j)
    return len(ReadDigits(b, i, j, false))
}

func GetStart(b Board, i int, j int) (int, int) {
    // re-position to the start if needed
    if !IsOnBoard(b, i, j) || !IsDigit(b.Board[i][j]) {
        return i, j
    }
    for ; IsOnBoard(b, i, j - 1) && IsDigit(b.Board[i][j - 1]) ; {
        j--
    }
    return i, j
}

// return multiply result of connected gears
func IsGears(b Board, i int, j int) bool {
    i, j = GetStart(b, i, j)
    length := GetNumLen(b, i, j)
    fmt.Printf("GEEE: %v\n", ReadDigits(b, i, j, false))
    for k := j; k < j + length; k++ {
        for _, d := range directions {
            x := i + d[0]
            y := k + d[1]
            if !IsOnBoard(b, x, y) || b.Board[x][y] != '*' {
                continue
            }
            for _, d := range directions {
                x := x + d[0]
                y := y + d[1]
                // check if this is part of outer loop number
                if x == i && j <= y && y < j + length {
                    fmt.Println("outer loop num!")
                    continue
                }
                if IsOnBoard(b, x, y) && !b.Used[x][y] && IsDigit(b.Board[x][y]) {
                    return true
                }
            }
        }
    }
    return false
}


func ReadDigits(b Board, i int, j int, markUsed bool) string {
    i, j = GetStart(b, i, j)
    s := bytes.NewBufferString("")
    for k := j; IsOnBoard(b, i, k) && IsDigit(b.Board[i][k]); k++ {
        s.WriteRune(b.Board[i][k])
        if markUsed {
            b.Used[i][k] = true
        }
    }
    return s.String()
}


func GetPartNum(b Board, i int, j int) int {
    i, j = GetStart(b, i, j)
    digits := ReadDigits(b, i, j, true)
    fmt.Printf("GetPartNum: %s\n", digits)
    if digits == "" {
        return 0
    }

    // multiplier for gear pair
    pair := 0
    // check surroundings (flood fill)
    // if * is connecting two numbers
    for k := j; k < j + len(digits); k++ {
        for _, d := range directions {
            x := i + d[0]
            y := k + d[1]
            //fmt.Printf("check (%d, %d)\n", x, y)
            if !IsOnBoard(b, x, y) || b.Board[x][y] != '*' {
                continue
            }
            // fmt.Printf("found * at (%d, %d)\n", x, y)
            for _, d := range directions {
                x := x + d[0]
                y := y + d[1]
                // fmt.Printf("scan surrounding (%d, %d)\n", x, y)
                // checking Used
                // prevents from reading already
                // consumed number in infinite recursion
                if IsOnBoard(b, x, y) && !b.Used[x][y] && IsDigit(b.Board[x][y]) {
                    pair = GetPartNum(b, x, y)
                    fmt.Printf("found pair %d \n", pair)
                }
            }
        }
    }
    num, _ := strconv.Atoi(digits)
    return num * pair
}

func main() {
    b := Board {
        Board: make([][]rune, 0),
        Used: make([][]bool, 0),
    }
    for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
        b.Board = append(b.Board, []rune(sc.Text()))
        b.Used = append(b.Used, make([]bool, len(b.Board[len(b.Board) - 1])))
    }
    // fmt.Printf("%v\n", b.Board)
    // fmt.Printf("%v\n", b.Used)
    sum, i, j := 0, 0, 0
    for ; i < len(b.Board); i++ {
        for ; j < len(b.Board[0]); {
            l := GetNumLen(b, i, j)
            pn := GetPartNum(b, i, j)
            isGears := IsGears(b, i, j)
            fmt.Printf("Is Gears? %v\n", isGears)
            if isGears {
                fmt.Printf("Gear val: %v\n", pn)
                sum += pn
            } else {
                fmt.Println("Gear val: -")
            }
            if l != 0 {
                j += l
            } else {
                j++
            }
        }
        j = 0
    }
    fmt.Printf("%d\n", sum)
}
