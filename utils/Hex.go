package utils

import "fmt"

//PrintBytes - Prints bytes out as hex
func PrintBytes(name string, bytes []byte) {
	fmt.Printf("%s: [% x]\n", name, bytes)
}
