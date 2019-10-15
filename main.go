package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in format of qustion and awnswers")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("unable to open CSV Filename: %s \n", *csvFileName))
	} else {
		fmt.Println("There you go!!.")
	}
	_ = file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to open the file")
	}
	_ = lines
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
