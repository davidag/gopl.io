// Units converts its numeric argument to different temperatures
package main

import (
	"bufio"
	"fmt"
	"gopl.io/ch2/tempconv"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		// Read from standard input
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			values := strings.Split(input.Text(), " ")
			for _, value := range values {
				if len(value) > 0 {
					process(value)
				}
			}
		}
	} else {
		// Read from command arguments
		for _, arg := range os.Args[1:] {
			process(arg)
		}
	}
}

func process(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "units: %v\n", err)
		return
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	k := tempconv.Kelvin(t)
	fmt.Printf("%s = %s, %s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c), k, tempconv.KToC(k))
}
