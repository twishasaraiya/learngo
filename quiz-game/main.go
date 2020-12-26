package main

import (
	"os"
	"log"
	"encoding/csv"
	"fmt"
	"strconv"
	"flag"
	"strings"
	"math/rand"
	"time"
)


type problem struct{
	question string
	correctAnswer string
}

func ParseCSV(rows [][]string) []problem{
	data := make([]problem, len(rows))
	for idx,row := range rows {
		data[idx] = problem{
			question: row[0],
			correctAnswer: strings.TrimSpace(row[1]),
		}
	}
	return data
}

func Shuffle(problems []problem) []problem{
	var data = make([]problem, len(problems))
	random := rand.New(rand.NewSource(time.Now().Unix()))
	permutation := random.Perm(len(problems))
	for i,randomIdx := range permutation{
		data[i] = problems[randomIdx]
	}
	return data;
}
/*
	log: 21st December 2020
	What did we learn ?
	1. Install go
	2. Read a file
	3. parse a csv file
	4. convert string to integer
	5. Command line flag
*/
func main() {
	var filePath = flag.String("csv", "problems.csv","a csv file in the format of 'question,answer' (default 'problems.csv')")

	var shuffle = flag.Bool("shuffle", false, "a boolean flag to shuffle the order of questions")

	flag.Parse()
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err);
	}

	problems := ParseCSV(data)
	// fmt.Println(problems)
	if *shuffle {
		problems = Shuffle(problems) 
	}

	score := 0
	var answer int
	for idx, problem := range problems {
		fmt.Printf("Question %d: %s = ", idx + 1, problem.question)
		fmt.Scanf("%d", &answer)
		correctAnswer, _ := strconv.Atoi(problem.correctAnswer)

		if answer == correctAnswer {
			score += 1
		}
	}

	fmt.Print("Your score is ",score)
	file.Close()
}