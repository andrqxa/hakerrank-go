package main

import (
	// "bufio"
	"fmt"
	"math"
	// "io"
	// "os"
	// "strconv"
	// "strings"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) (res int32) {
	// Write your code here
	sz := len(arr[0])
	cMain := make(chan int32)
	cOpps := make(chan int32)
	go func(c chan int32) {
		var mainDiag int32 = 0
		for i := 0; i < sz; i++ {
			mainDiag += arr[i][i]
		}
		c <- mainDiag
		close(c)
	}(cMain)
	go func(c chan int32) {
		var oppDiag int32 = 0
		j := 0
		for i := sz - 1; i >= 0; i-- {
			oppDiag += arr[i][j]
			j++
		}
		c <- oppDiag
		close(c)
	}(cOpps)
	a := float64(<-cMain)
	b := float64(<-cOpps)
	return int32(math.Abs(a - b))
}

func main() {
	a0 := []int32{11, 2, 4}
	a1 := []int32{4, 5, 6}
	a2 := []int32{10, 8, -12}
	arr := [][]int32{}
	arr = append(arr, a0)
	arr = append(arr, a1)
	arr = append(arr, a2)
	fmt.Println(arr)
	fmt.Println(diagonalDifference(arr))

	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv(""))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	// nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// checkError(err)
	// n := int32(nTemp)

	// var arr [][]int32
	// for i := 0; i < int(n); i++ {
	// 	arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

	// 	var arrRow []int32
	// 	for _, arrRowItem := range arrRowTemp {
	// 		arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
	// 		checkError(err)
	// 		arrItem := int32(arrItemTemp)
	// 		arrRow = append(arrRow, arrItem)
	// 	}

	// 	if len(arrRow) != int(n) {
	// 		panic("Bad input")
	// 	}

	// 	arr = append(arr, arrRow)
	// }

	// result := diagonalDifference(arr)

	// fmt.Fprintf(writer, "%d\n", result)

	// writer.Flush()
}

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
