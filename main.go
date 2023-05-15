package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func reName(newFileName string ,fileName string){
	err:= os.Rename(fileName, newFileName)
	if err!= nil {
		log.Fatal(err)
	}
}
func openFile(fileName string) *os.File{
	// open file
	f, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
	return f
}

func checkAnswer(answer string, userAnswer string,score int) int{
	if userAnswer==answer {
	 	score++
	}
	return score
}

func main(){
	fmt.Println("press 1 to rename a file or press other to do nothing.")
	var wRename string
	fmt.Scanln(&wRename)
	const fileName = "problems.csv"
	var newFileName string
	if wRename=="1" {
		fmt.Scanln(&newFileName)
		reName(newFileName, fileName)
	}
	round := 0
    score := 0
	// read csv values using csv.Reader
	f := openFile(newFileName)
	// remember to close the file at the end of the program
	defer f.Close()
    csvReader := csv.NewReader(f)
	for {
		//read each line
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
		fmt.Print(question+" ")
		//input
		fmt.Scanln(&userAnswer)
		score = checkAnswer(answer, userAnswer, score)
		round++
    }
	//final score
	fmt.Print("Your score "+strconv.Itoa(score)+" out of "+strconv.Itoa(round))
	if wRename == "1" {
		reName(fileName, newFileName)
	}
}