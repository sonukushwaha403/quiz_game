package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvfilename := flag.String("csv", "problem.csv", "provide a csv queston file in the following format 'ouestion,answer' formant only")
	timelimit := flag.Int("limit", 30, "enter time in seconds for each question")
	flag.Parse()
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	file, err := os.Open(*csvfilename)
	if err != nil {
		exit(fmt.Sprintf("cannot open the file named %s \n", *csvfilename))
	}
	r := csv.NewReader(file)
	line, err := r.ReadAll()
	if err != nil {
		exit("failed to read the provided csv file ")
	}
	//fmt.Println(line)

	problem := parseLine(line)
	correct := 0
	for i, p := range problem {
		fmt.Printf("Problem #%d is %s = ", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s \n", &ans)
			answerCh <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nthe number of correct answer out of %d is %d", len(line), correct)
			return
		case ans := <-answerCh:

			if ans == p.answer {
				//fmt.Println("correct")
				correct++
			}

		}
	}
	fmt.Printf("the number of correct answer out of %d is %d", len(line), correct)

	//fmt.Println(problem)

}

func parseLine(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{line[0], strings.TrimSpace(line[1])}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(name string) {
	fmt.Println(name)
	os.Exit(1)
}

func sonu(name string) *string {

	return &name
}
