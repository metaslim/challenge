package util

import (
	"bufio"
)

func ReadInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]

	return input
}
