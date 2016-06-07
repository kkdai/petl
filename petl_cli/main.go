package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	etl "github.com/kkdai/petl"
)

func main() {
	pattern := flag.String("pattern", "", "Pattern definition to look for")
	flag.Parse()

	if *pattern == "" {
		fmt.Println("Pattern argument is missing.")
		fmt.Println("Usage:")
		flag.PrintDefaults()
		return
	}

	info, _ := os.Stdin.Stat()

	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage:")
		fmt.Println("  cat yourfile.txt | searchr -pattern=<your_pattern>")
	} else if info.Size() > 0 {
		reader := bufio.NewReader(os.Stdin)
		match(*pattern, reader)
	}
}

func match(pattern string, reader *bufio.Reader) {
	line := 1
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}

		color := "\x1b[39m"
		// if strings.Contains(input, pattern) {
		// 	color = "\x1b[31m"
		// }

		// fmt.Printf("%s%2d: %s", color, line, input)
		fmt.Printf("%s%2d: %s", color, line, <-etl.TransformRemoveSpace(etl.Extract(input)))
		line++

	}
}
