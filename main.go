package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/semihsemih/go-terminal-ui"
	"github.com/semihsemih/go-terminal-ui/color"
	"github.com/semihsemih/go-terminal-ui/font"
	"github.com/semihsemih/go-terminal-ui/screen"
)

type Problem struct {
	question string
	answer   string
}

func init() {
	screen.ClearScreen()
}

func main() {
	fmt.Println(font.Underline(color.BackgroundBrightCyan("______________Questions______________")))

	csvFilename := flag.String("csv", "questions.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	shuffleQuestions := flag.String("shuffle", "on", "Shuffle the order of the questions in each exam.")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("%s: %s\n", color.BackgroundBrightRed("Failed to open CSV file: "), *csvFilename))
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(color.BackgroundBrightRed("Failed to parse the provided CSV file."))
	}

	problems := parseLines(lines)
	if *shuffleQuestions == "on" {
		problems = shuffleProblems(problems)
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

problemLoop:
	for index, problem := range problems {
		fmt.Print(font.Bold("Problem #" + strconv.Itoa(index+1)))
		fmt.Printf(" %s = ", problem.question)

		answerChannel := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerChannel:
			if answer == problem.answer {
				correct++
			}
		}
	}

	fmt.Print(color.Green("\nYou scored " + strconv.Itoa(correct) + " out of " + strconv.Itoa(len(problems)) + "\n"))
	screen.ToExitKeyPress()
	screen.RestoreScreen()
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{line[0], strings.TrimSpace(line[1])}
	}

	return ret
}

func shuffleProblems(problems []Problem) []Problem {
	rand.Seed(time.Now().UnixNano())
	currentIndex := len(problems)

	// While there remain elements to shuffle
	for 0 != currentIndex {
		// Pick a remaining element...
		randomIndex := rand.Intn(currentIndex)
		currentIndex--

		// And swap it with the current element.
		temporaryValue := problems[currentIndex]
		problems[currentIndex] = problems[randomIndex]
		problems[randomIndex] = temporaryValue
	}

	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
