package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/metaslim/challenge/lib/setup"
	"github.com/metaslim/challenge/lib/textcolor"
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

	textcolor.Yellow("\n<Command> [help|quit] ➜ ")
	reader := bufio.NewReader(os.Stdin)
	input := util.ReadInput(reader)

	for input != "quit" {
		if input == "" {
			input = "help"
		}

		for _, command := range commands {
			command.SetInput(input)
			if command.Valid() {
				command.Run(dataSet)
				break
			}
		}
		textcolor.Yellow("\n<Command> [help|quit] ➜ ")
		input = util.ReadInput(reader)
	}

	fmt.Println("Bye and have a nice day ahead!")
}
