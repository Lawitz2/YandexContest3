package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const utfshiftcapital = -64
const utfshiftnumber = -48

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("failed to open a file: %v", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var result int

	fo, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf("failed to create a file: %v", err)
		return
	}

	writer := bufio.NewWriter(fo)

	for scanner.Scan() {
		text := scanner.Text()
		result = 0
		textslices := strings.Split(text, ",")

		uniquechars := make(map[byte]struct{})
		for _, i := range textslices[:3] { // count the amount of unique characters
			for z := 0; z < len(i); z++ {
				uniquechars[i[z]] = struct{}{}
			}
		}

		textslices[0] = strings.ToUpper(textslices[0])

		result += (int(textslices[0][0])+utfshiftcapital)*256 + len(uniquechars)

		for _, i := range textslices[3:5] {
			for _, d := range i {
				result += int(d+utfshiftnumber) * 64
			}
		}

		hexresult := strconv.FormatInt(int64(result), 16)
		hexupper := strings.ToUpper(hexresult)

		writer.WriteString(hexupper[max(len(hexupper)-3, 0):])
		writer.WriteString(" ")
	}
	writer.Flush()
}
