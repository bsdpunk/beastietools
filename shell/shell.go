package shell

import "fmt"
import "os"
import "strings"
import "github.com/codegangsta/cli"
import "github.com/bsdpunk/beastietools/command"
import "github.com/gobs/readline"


var quit string = "quit"
var GlobalFlags = []cli.Flag{}

var found string = "no"
var Commands = []cli.Command{
    {
        Name:   "extip",
        Usage:  "",
        Action: command.CmdExtip,
        Flags:  []cli.Flag{},
    },
    {
        Name:   "arp",
        Usage:  "",
        Action: command.CmdArp,
        Flags:  []cli.Flag{},
    },

}

var words []string 


var matches = make([]string, 0, len(words))

func AttemptedCompletion(text string, start, end int) []string {
    if start == 0 { // this is the command to match
        return readline.CompletionMatches(text, CompletionEntry)
    } else {
        return nil
    }
}


func CompletionEntry(prefix string, index int) string {
    if index == 0 {
        matches = matches[:0]

        for _, w := range words {
            if strings.HasPrefix(w, prefix) {
                matches = append(matches, w)
            }
        }
    }

    if index < len(matches) {
        return matches[index]
    } else {
        return ""
    }
}




func CommandNotFound(c *cli.Context, command string) {
    fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
    os.Exit(2)
}

func Run() {

for _, c := range Commands {
    words = append(words, c.Name)
}
words = append(words, "quit")
words = append(words, "ls")

prompt := "beastietools> " 
matches = make([]string, 0, len(words))


L:
    for {
        found = "no"
        readline.SetCompletionEntryFunction(CompletionEntry)
        readline.SetAttemptedCompletionFunction(nil)
        result := readline.ReadLine(&prompt)
        if result == nil { // exit loop
            break L
        }
               
        input := *result
        input = strings.TrimSpace(input)
        if input  == quit {
            os.Exit(1)
        }else if input == "ls" {
            fmt.Println(Commands)
        }else{

            for _, c := range Commands {
                if c.HasName(input) {
                    
                    var command []string
                    command = append(command, "beastietools")
                    command = append(command, input)
                    app := cli.NewApp()
                    app.Author = "bsdpunk"
                    app.Email = ""
                    app.Usage = ""
                    app.Name = "beastietools"
                    app.Version = "0.1.0"

                    app.Flags = GlobalFlags
                    app.Commands = Commands
                    app.CommandNotFound = CommandNotFound

                    app.Run(command)
                    found = "yes"

                }
            }
            if found == "no" {
                fmt.Println("Invalid Command")
            }
            readline.AddHistory(input) 
        }

    }
}


func PrintSlice(slice []string) {
    fmt.Printf("Slice length = %d\r\n", len(slice))
    for i := 0; i < len(slice); i++ {
        fmt.Printf("[%d] := %s\r\n", i, slice[i])
    }
}




