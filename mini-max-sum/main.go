package main

// Given five positive integers, find the minimum and maximum values that can be calculated by summingexactly four of the five integers. Then print the respective minimum and maximum values as a single lineof two space-separated long integers.
// Output FormatPrint two space-separated long integers denoting the respective minimum and maximum values that canbe calculated by summing exactly four of the five integers. (The output can be greater than a 32 bitinteger.)

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	chainMax := make(chan float64)
	chainMin := make(chan float64)

	go func(c chan float64) {
		var sum float64 = 0
		for i := 1; i < len(arr); i++ {
			sum += float64(arr[i])
		}
		c <- sum
		close(c)
	}(chainMax)

	go func(c chan float64) {
		var sum float64 = 0
		for i := 0; i < len(arr)-1; i++ {
			sum += float64(arr[i])
		}
		c <- sum
		close(c)
	}(chainMin)

	fmt.Printf("%.f %.f", <-chainMin, <-chainMax)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
