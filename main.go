/*
a89 BASIC v 0.0.0
20210514
arataca89@gmail.com
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

var programa string

func main() {
	var lineStr string
	osName := runtime.GOOS
	lineScanner := bufio.NewScanner(os.Stdin)

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("a89BASIC  0.0.0 rodando em sistema:", osName)
	fmt.Println("ajuda: HELP; sair:QUIT.")
	fmt.Println("PROGRAMA CARREGADO:", programa)
	fmt.Println(strings.Repeat("-", 50))

	for {
		fmt.Print(">")
		for lineScanner.Scan() {
			lineStr = lineScanner.Text()
			tokens := strings.Split(lineStr, " ")
			if tokens[0] == "QUIT" {
				return
			}
			fmt.Println(tokens)
			fmt.Print(">")

		}
		if err := lineScanner.Err(); err != nil {
			log.Println(err)
		}

	}

}
