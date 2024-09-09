package monkeylang

import (
	"fmt"
	"os"
	"os/user"

	"github.com/perryfranks/monkey-lang/repl"
)

func Entry() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)

	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func EvalErrors(input string) (progErrors []string, err error) {

	return []string{}, nil
}