package main

import (
    "os"
    "github.com/bsdpunk/beastietools/shell"
    "github.com/codegangsta/cli"
)

func main() {


    if len(os.Args) < 2{
        shell.Run()

    } else {

        app := cli.NewApp()
        app.Name = Name
        app.Version = Version
        app.Author = "bsdpunk"
        app.Email = ""
        app.Usage = ""

        app.Flags = GlobalFlags
        app.Commands = Commands
        app.CommandNotFound = CommandNotFound
        
        //fmt.Println(os.Args)
        app.Run(os.Args)
    }
}
