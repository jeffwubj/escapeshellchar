# escapeshellchar
escape/unescape shell special chars

```
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
```

* original string `a bc'|%;<>$"`
* escaped string `a\%20bc\%27\%7C%25\%3B\%3C\%3E\%24\"`
* unescaped string `a\ bc\'\|%\;\<\>\$\"`
