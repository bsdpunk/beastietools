package shell

import "fmt"
import "os"
import "bufio"
import "strings"

var quit string = "quit"

func Run() {
    sum := 0
    for sum < 1{
        
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("beastietools> ")
        text, _ := reader.ReadString('\n')
        //fmt.Println(text)
        //fmt.Println(thing)
        input := strings.TrimSpace(text)
        if input  == quit {
                os.Exit(1)
        }
    }
}
