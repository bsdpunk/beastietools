package command

import (
        "fmt"
        "io/ioutil"
        "github.com/codegangsta/cli"
    )



func CmdArp(c *cli.Context) error {
    contents, _ := ioutil.ReadFile("/proc/net/arp")
    fmt.Println(string(contents))
    return nil
}

