package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Question struct {
	Text   string `json:"text"`
	Answer string `json:"answer"`
	// Alternatives []string `json:"alternatives"`
}

func loadQuestionsFromJSON(filePath string) ([]Question, error) {
	byteValue, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	var questions []Question
	err = json.Unmarshal(byteValue, &questions)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON from %s: %w", filePath, err)
	}

	return questions, nil
}

func main() {
	quizFile := flag.String("file", "quiz.json", "Path to the JSON file containing quiz questions")
	numQuestionsToAsk := flag.Int("num", -1, "Total number of questions to ask (0 for all unique, -1 for all unique once, >0 for repetitions)")
	flag.Parse()

	fmt.Println("Welcome to the Go Quiz!")
	fmt.Printf("Loading questions from: %s\n", *quizFile)
	fmt.Println("------------------------------------------")
	time.Sleep(1 * time.Second)

	questions, err := loadQuestionsFromJSON(*quizFile)
	if err != nil {
		fmt.Printf("Failed to load quiz questions: %v\n", err)
		os.Exit(1)
	}

	if len(questions) == 0 {
		fmt.Println("No questions found in the quiz file. Exiting.")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	totalQuestions := len(questions)
	if *numQuestionsToAsk == -1 || *numQuestionsToAsk == 0 {
		rand.Shuffle(len(questions), func(i, j int) {
			questions[i], questions[j] = questions[j], questions[i]
		})
		totalQuestions = len(questions)
	} else if *numQuestionsToAsk > 0 {
		totalQuestions = *numQuestionsToAsk
	} else {
		fmt.Println("Invalid value for --num. Please use -1, 0, or a positive integer.")
		os.Exit(1)
	}

	score := 0
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < totalQuestions; i++ {
		randomIndex := rand.Intn(len(questions))
		q := questions[randomIndex]

		fmt.Printf("\nQuestion %d (of %d): %s\n", i+1, totalQuestions, q.Text)
		fmt.Print("Your answer: ")

		input, _ := reader.ReadString('\n')
		userAnswer := strings.TrimSpace(input)

		if strings.EqualFold(userAnswer, q.Answer) {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Printf("Incorrect. The correct answer was: %s\n", q.Answer)
		}
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("\n------------------------------------------")
	fmt.Println("Quiz finished!")
	fmt.Printf("You scored %d out of %d questions asked.\n", score, totalQuestions)
	fmt.Printf("There were %d unique questions in the file.\n", len(questions))

	if score == totalQuestions {
		fmt.Println("Excellent! You got all questions correct!")
	} else if score > totalQuestions/2 {
		fmt.Println("Good job! You passed the quiz.")
	} else {
		fmt.Println("Keep practicing! You can do better next time.")
	}
	fmt.Println("------------------------------------------")
}
