package main
import ("fmt"
        "os"
        "bufio")

func main(){
    sum := 0
    for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
        line := sc.Text()
        for i := range line {
            if line[i] >  '0' && line[i] <= '9' {
                sum += 10 * int(line[i] - '0')
                break
            }
        }
        for i := len(line) - 1; i >= 0; i-- {
            if line[i] >  '0' && line[i] <= '9' {
                sum += int(line[i] - '0')
                break
            }
        }
    }
    fmt.Println(sum)
}
