package main

// You are in charge of the cake for your niece's birthday and have decided the cake will have one candle for each year of her total age.
// When she blows out the candles, she’ll only be able to blow out the tallest ones. Your task is to find out how many candles she can successfully blow out.

// Input Format

// The first line contains a single integer, `n`, denoting the number of candles on the cake.
// The second line contains `n` space-separated integers, where each integer `i` describes the height of candle `i`.

// Constraints
// 1 <= n <= 10e5
// 1 <= ar[i] <= 10e7

// Output Format

// Return the number of candles that can be blown out on a new line.

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] >= arr[j] })

	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return 1
	}

	var cnt int32 = 1
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return cnt
		}
		cnt++
	}
	return cnt
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(arCount); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := birthdayCakeCandles(ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
