package main

import (
	"fmt"
	"kulang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome to kulang, %s: \n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
