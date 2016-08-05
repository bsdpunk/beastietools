package shell

import "fmt"
import "os"
import "bufio"
import "strings"
import "github.com/codegangsta/cli"
import "github.com/bsdpunk/beastietools/command"
var quit string = "quit"
var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
    {
        Name:   "extip",
        Usage:  "",
        Action: command.CmdExtip,
        Flags:  []cli.Flag{},
    },
    {
        Name:   "profile",
        Usage:  "",
        Action: command.CmdProfile,
        Flags:  []cli.Flag{},
    },
}

func CommandNotFound(c *cli.Context, command string) {
    fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
    os.Exit(2)
}

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
        }else if input == "ls" {
            fmt.Println(Commands)
        }else if input == "extip" {
            var command []string
            command = append(command, "beastietools")
            command = append(command, input)
            app := cli.NewApp()
            app.Author = "bsdpunk"
            app.Email = ""
            app.Usage = ""

            app.Flags = GlobalFlags
            app.Commands = Commands
            app.CommandNotFound = CommandNotFound

            app.Run(command)

        }
    }
}

func PrintSlice(slice []string) {
    fmt.Printf("Slice length = %d\r\n", len(slice))
    for i := 0; i < len(slice); i++ {
        fmt.Printf("[%d] := %s\r\n", i, slice[i])
    }
}
