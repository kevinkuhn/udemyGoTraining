package main

import "fmt"

func main() {
	for i := 32; i < 121; i++ {

		// explenations: https://golang.org/pkg/fmt/

		// %d = decimal (0-9), base 10
		// %b = binary (0/1), base 2
		// %x = hex (16), base 16, with lower-case letters for a-f
		// %X = base 16, with upper-case letters for A-F
		// %q = a single-quoted character literal safely escaped with Go syntax.
		// %s = the uninterpreted bytes of the string or slice
		// %U = Unicode format: U+1234; same as "U+%04X"

		fmt.Printf("%d \t %b \t %x \t %X \t %q \t %s \t %U \n", i, i,
			i, i, i, string(i), i)
	}
}
