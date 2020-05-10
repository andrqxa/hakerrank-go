package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getSimbols(n int, simb string) string {
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		arr = append(arr, simb)
	}
	return strings.Join(arr, "")
}

// Complete the staircase function below.
func staircase(n int32) {
	spice := int(n - 1)
	for i := 1; i <= int(n); i++ {
		simbs := getSimbols(i, "#")
		spaces := getSimbols(spice, " ")
		fmt.Printf("%s%s\n", spaces, simbs)
		spice--
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
