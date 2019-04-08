package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	csvfile, _ := os.Open("problems.csv")
	reader := csv.NewReader(bufio.NewReader(csvfile))

	buf := bufio.NewReader(os.Stdin)

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		fmt.Println(line[0])
		sentence, _ := buf.ReadString('\n')
		fmt.Print(sentence)
		if strings.TrimRight(sentence, "\n") == string(line[1]) {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}
