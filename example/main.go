package main

import (
	"fmt"

	"github.com/jeffwubj/escapeshellchar"
)

func main() {
	escapedstr := escapeshellchar.EscapeShellString("';lkgfdq1-]\\=-00-[0")
	fmt.Println(escapedstr)
	fmt.Println(escapeshellchar.UnEscapeShellString(escapedstr))
}
