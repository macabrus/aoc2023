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
    for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
        gameAndCubes := strings.Split(sc.Text(), ":")
        gamePart := strings.Split(strings.TrimSpace(gameAndCubes[0]), " ")
        sampleParts := strings.Split(strings.TrimSpace(gameAndCubes[1]), ";")
        minCubes := map[string]int {
            "red": 0,
            "green": 0,
            "blue": 0,
        }
        for _, part := range sampleParts {
            cubes := strings.Split(strings.TrimSpace(part), ",")
            for _, cube := range cubes {
                colorNum := strings.Split(strings.TrimSpace(cube), " ")
                num, _ := strconv.Atoi(colorNum[0])
                color := strings.TrimSpace(colorNum[1])
                minCubes[color] = max(num, minCubes[color])
            }
        }
        sum += minCubes["red"] * minCubes["green"] * minCubes["blue"]
    }
    fmt.Printf("%d\n", sum)
}
