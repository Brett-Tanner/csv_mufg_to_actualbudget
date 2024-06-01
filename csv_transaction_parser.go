package main

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	path := "/Users/brett/Downloads/new_transactions"
	inputCSVs := getCSVs(path)
	err := processCSVs(inputCSVs, path)
	handleErr(err)

	os.Exit(0)
}

func getCSVs(path string) []fs.DirEntry {
	files, err := os.ReadDir(path)
	handleErr(err)

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

func processCSVs(inputCSVs []fs.DirEntry, inputPath string) error {
	outputPath := getOutputPath(inputPath)
	dirErr := os.MkdirAll(outputPath, 0777)
	handleErr(dirErr)

	for _, inputCSV := range inputCSVs {
		inputIO, inputIOErr := os.Open(filepathFor(inputCSV.Name(), inputPath))
		defer inputIO.Close()
		handleErr(inputIOErr)

		outputIO, outputIOErr := os.Create(filepathFor(inputCSV.Name(), outputPath))
		defer outputIO.Close()
		handleErr(outputIOErr)

		rows, err := csv.NewReader(inputIO).ReadAll()
		handleErr(err)
		writer := csv.NewWriter(outputIO)
		createHeaders(writer)
		for i, row := range rows {
			if i == 0 {
				continue
			}

			transformRow(writer, row)
		}
		writer.Flush()
		handleErr(writer.Error())
	}

	return nil
}

func getOutputPath(inputPath string) string {
	pathSegments := strings.Split(inputPath, "/")

	parent := pathSegments[0 : len(pathSegments)-1]
	parent = append(parent, "processed_transactions")

	return strings.Join(parent, "/")
}

func filepathFor(filename, path string) string {
	var filepath strings.Builder

	filepath.WriteString(path)
	filepath.WriteString("/")
	filepath.WriteString(filename)

	return filepath.String()
}

func createHeaders(writer *csv.Writer) [4]string {
	headers := [4]string{"Date", "Payee", "Outflow", "Inflow"}
	writer.Write(headers[:])
	return headers
}

func transformRow(writer *csv.Writer, inputHeaders []string) [4]string {
	date := inputHeaders[0]
	payee := asUTF8(inputHeaders[2])
	outflow := inputHeaders[3]
	inflow := inputHeaders[4]

	outputRow := [4]string{date, payee, outflow, inflow}
	writer.Write(outputRow[:])
	return outputRow
}

func asUTF8(japaneseString string) string {
	var result strings.Builder

	winUTF8 := transform.NewWriter(&result, japanese.ShiftJIS.NewDecoder())
	winUTF8.Write([]byte(japaneseString))
	winUTF8.Close()
	fmt.Println(result.String())
	return result.String()
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
