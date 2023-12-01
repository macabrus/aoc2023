package main
import ("fmt"
        "os"
        "bufio"
        "strings")

func main(){
    sum := 0
    words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    num_val := make(map[string]int)
    for i, _ := range words {
        num_val[words[i]] = i
    }
    for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
        line := sc.Text()
        for i := range line {
            if line[i] >  '0' && line[i] <= '9' {
                sum += 10 * int(line[i] - '0')
                break
            }
            found := false
            for _, word := range words {
                if strings.HasPrefix(line[i:], word) {
                    sum += 10 * num_val[word]
                    found = true
                    break
                }
            }
            if found {
                break
            }
        }
        for i := len(line) - 1; i >= 0; i-- {
            if line[i] >  '0' && line[i] <= '9' {
                sum += int(line[i] - '0')
                break
            }
            found := false
            for _, word := range words {
                if strings.HasSuffix(line[:i+1], word) {
                    sum += num_val[word]
                    found = true
                    break
                }
            }
            if found {
                break
            }
        }
    }
    fmt.Print(sum)
}
