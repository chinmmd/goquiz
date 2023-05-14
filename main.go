package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)
func main(){
	// open file
    f, err := os.Open("problems.csv")
    if err != nil {
        log.Fatal(err)
    }
	round := 0
    score := 0
	// remember to close the file at the end of the program
    defer f.Close()
	// read csv values using csv.Reader
    csvReader := csv.NewReader(f)
	for {
        rec, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
		question := rec[0]
		answer := rec[1]
		var userAnswer string
        // do something with read line
        //fmt.Printf("%v %T\n", rec, rec)
		//fmt.Printf("%v %T\n", question, question)
		//fmt.Printf("%v %T\n", answer, answer)
		fmt.Print(question+" ")
		fmt.Scanln(&userAnswer)
        if userAnswer==answer {
			score++
		}
		round++
    }
	fmt.Print("Your score "+strconv.Itoa(score)+" out of "+strconv.Itoa(round))
}