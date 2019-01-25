package main

import (
	"fmt"

	"github.com/jeffwubj/escapeshellchar"
)

func main() {
	escapedstr := escapeshellchar.EscapeShellString("a bc'|%;<>$\"")
	fmt.Println(escapedstr)
	fmt.Println(escapeshellchar.UnEscapeShellString(escapedstr))
}
