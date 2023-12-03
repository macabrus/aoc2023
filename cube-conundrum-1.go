package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main(){
    sum := 0
    maxNums := map[string]int{
        "red": 12,
        "green": 13,
        "blue": 14,
    }
    for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
        gameAndCubes := strings.Split(sc.Text(), ":")
        gamePart := strings.Split(strings.TrimSpace(gameAndCubes[0]), " ")
        gameNum, _ := strconv.Atoi(gamePart[1])
        sampleParts := strings.Split(strings.TrimSpace(gameAndCubes[1]), ";")
        bad := false
        for _, part := range sampleParts {
            cubes := strings.Split(strings.TrimSpace(part), ",")
            for _, cube := range cubes {
                colorNum := strings.Split(strings.TrimSpace(cube), " ")
                num, _ := strconv.Atoi(colorNum[0])
                color := strings.TrimSpace(colorNum[1])
                if maxNum, ok := maxNums[color]; ok && maxNum < num {
                    bad = true
                    break
                }
            }
            if bad {
                break
            }
        }
        if !bad {
            sum += gameNum
        }
    }
    fmt.Printf("%d\n", sum)
}
