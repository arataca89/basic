/*
a89 BASIC v 0.0.0
20210514
arataca89@gmail.com
*/
package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	var linha string
	os := runtime.GOOS
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("a89BASIC  0.0.0 rodando em sistema:", os)
	fmt.Println("ajuda: HELP; sair:QUIT.")
	fmt.Println(strings.Repeat("-", 50))
	for {
		fmt.Print(">")
		fmt.Scanf("%s\r", &linha)
		if linha == "QUIT" {
			return
		}

	}

}
