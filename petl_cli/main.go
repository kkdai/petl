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
	pipeline := flag.String("pipeline", "", "Pipeline list for this ETL process")
	flag.Parse()

	if *pipeline == "" {
		fmt.Println("Pattern argument is missing.")
		fmt.Println("Usage:")
		flag.PrintDefaults()
		return
	}

	info, _ := os.Stdin.Stat()

	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage:")
		fmt.Println("  cat yourfile.txt | petl_cli -pipelone=<your_pipeline>")
	} else if info.Size() > 0 {
		reader := bufio.NewReader(os.Stdin)
		processPipe(preparePipeline(*pipeline), reader)
	}
}

func preparePipeline(pline string) []etl.Pipeline {
	var retPipeline []etl.Pipeline

	for _, str := range pline {
		var targetPipe etl.Pipeline
		switch {
		case string(str) == "R" || string(str) == "r":
			targetPipe = etl.TransformRemoveSpace
		case string(str) == "L" || string(str) == "l":
			targetPipe = etl.TransformLower
		case string(str) == "U" || string(str) == "u":
			targetPipe = etl.TransformUpper
		default:
			targetPipe = etl.TransformDefault
		}
		retPipeline = append(retPipeline, targetPipe)
	}
	return retPipeline
}

func processPipe(plines []etl.Pipeline, reader *bufio.Reader) {
	line := 1
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}

		out := etl.PipeProcess(etl.Extract(input), plines)
		fmt.Printf("%2d: %s", line, <-out)
		line++

	}
}
