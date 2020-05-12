package main

// Given a time in 12-hour AM/PM format, convert it to military (24-hour) time.

// Note: Midnight is 12:00:00AM on a 12-hour clock, and 00:00:00 on a 24-hour clock. Noon is 12:00:00PM on a 12-hour clock, and 12:00:00 on a 24-hour clock.

// Function Description

// Complete the timeConversion function in the editor below. It should return a new string representing the input time in 24 hour format.

// timeConversion has the following parameter(s):

//     s: a string representing time in 12-hour format

// Input Format

// A single string
// containing a time in 12-hour clock format (i.e.: hh:mm:ssAM or hh:mm:ssPM),
// where 01<=hh<=12 and 0 <= mm,ss <= 59

// Constraints

//     All input times are valid

// Output Format

// Convert and print the given time in 24-hour format, where 0 <= hh <= 23.

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	/*
	 * Write your code here.
	 */

	tm := strings.Split(s[:8], ":")
	switch s[8:] {
	case "AM":
		func(st string) {
			if st == "12" {
				tm[0] = "00"
			}
		}(tm[0])
	case "PM":
		func(st string) {
			switch st {
			case "01":
				tm[0] = "13"
			case "02":
				tm[0] = "14"
			case "03":
				tm[0] = "15"
			case "04":
				tm[0] = "16"
			case "05":
				tm[0] = "17"
			case "06":
				tm[0] = "18"
			case "07":
				tm[0] = "19"
			case "08":
				tm[0] = "20"
			case "09":
				tm[0] = "21"
			case "10":
				tm[0] = "22"
			case "11":
				tm[0] = "23"
			}
		}(tm[0])
	}
	return strings.Join(tm, ":")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)
	fmt.Println(result)

	fmt.Fprintf(writer, "%s\n", result)

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
