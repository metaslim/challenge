package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/metaslim/challenge/lib/setup"
	"github.com/metaslim/challenge/lib/util"
)

func main() {

	appConfig, err := util.LoadConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataSet, err := setup.LoadDataSet(appConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	commands := setup.PrepareCommand()

	fmt.Println("Enter Command:")
	fmt.Println("==================================")
	fmt.Println("\t`quit` to exit")
	fmt.Println("\t`help` to get help")
	fmt.Println("==================================")

	reader := bufio.NewReader(os.Stdin)
	input := util.ReadInput(reader)

	for input != "quit" {
		for _, command := range commands {
			command.SetInput(input)
			if command.Valid() {
				command.Run(dataSet)
				break
			}
		}
		input = util.ReadInput(reader)
	}

	fmt.Println("Bye and have a nice day ahead!")
}
