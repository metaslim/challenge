package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/metaslim/challenge/lib/helper"
)

func main() {
	decorateParams, err := helper.LoadDecorateParams()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	commands := helper.PrepareCommand()

	fmt.Println("Enter Command:")
	fmt.Println("==================================")
	fmt.Println("\t`quit` to exit")
	fmt.Println("\t`help` to get help")
	fmt.Println("==================================")

	reader := bufio.NewReader(os.Stdin)
	input := helper.ReadInput(reader)

	for input != "quit" {
		for _, command := range commands {
			command.Run(input, decorateParams)
		}
		input = helper.ReadInput(reader)
	}

	fmt.Println("Bye and have a nice day ahead!")
}
