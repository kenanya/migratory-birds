package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the migratoryBirds function below.
func migratoryBirds(arr []int32) int32 {
	birdsGroup := make(map[int32]int32)

	// var ok bool = false
	for ii:=0; ii<len(arr); ii++ {
		if _, ok := birdsGroup[arr[ii]]; ok {
			birdsGroup[arr[ii]]++
		} else {
			birdsGroup[arr[ii]] = 1
		}
	}

	birdTypeMost, valueMost := int32(-1), int32(-1)
	for k, v := range birdsGroup {
		if v >= valueMost {
			if v == valueMost {
				if k < birdTypeMost {
					birdTypeMost = k					
				}				
			} else {
				birdTypeMost = k
			}			
			valueMost = v
		}
	}	

	return birdTypeMost
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(arrCount); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := migratoryBirds(arr)

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
