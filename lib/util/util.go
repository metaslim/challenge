package util

import (
	"bufio"
)

//ReadInput will read user keyboard input and remove the trialing newline
func ReadInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]

	return input
}
