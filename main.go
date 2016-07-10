package main

import (
    "github.com/joehandzik/git-go-quip/cmd"
    "fmt"
    "os"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}
