package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func main() {
	inputCSVs := getCSVs("/Users/brett/Downloads/new_transactions")
	fmt.Println(inputCSVs)
	// ParseTransactions(inputCSVs)
}

func getCSVs(path string) []fs.DirEntry {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	var inputCSVs []fs.DirEntry
	for _, file := range files {
		if isCSV(file.Name()) {
			inputCSVs = append(inputCSVs, file)
		}
	}

	return inputCSVs
}

func isCSV(filename string) bool {
	segments := strings.Split(filename, ".")
	if segments[len(segments)-1] == "csv" {
		return true
	}

	return false
}

// func ParseTransactions(inputFiles []fs.DirEntry) {
// 	for _, file := range inputFiles {
//
// 		inputFile, err := os.Open(file.Name())
// 		defer inputFile.Close()
// 		outputFile, err := os.Create(file.Name() + "output")
// 		outputWriter := csv.NewWriter(outputFile)
// 		defer outputFile.Close()
// 		if err != nil {
// 			fmt.Println(err)
// 		}
//
// 		rows, err := csv.NewReader(inputFile).ReadAll()
// 		for i, row := range rows {
// 			if i == 0 {
// 				translateHeaders(outputWriter, row)
// 			}
//
// 			translateRow(outputWriter, row)
// 		}
// 	}
// }
//
// func translateHeaders(w *csv.Writer, row []string) [4]string {
// 	return [4]string{"", "", "", ""}
// }
//
// func translateRow(w *csv.Writer, row []string) [4]string {
// 	for i, cell := range row {
// 		switch i {
// 		case 0:
// 		}
// 	}
// 	return [4]string{"", "", "", ""}
// }
