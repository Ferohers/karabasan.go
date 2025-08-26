package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"math/rand"
)

var (
	userName       string
	score          int
	previousJokeID int = -1
	errorCount     int
	content        Content
)

// ANSI renkler gelsin macde dene bunu.
const (
	ColorCyan   = "\033[36m"
	ColorGreen  = "\033[32m"
	ColorReset  = "\033[0m"
	
	// Assuming a standard terminal width for centering text.
	terminalWidth = 80
	// Separator for the user input area.
	separator = "================================================================================"
)

// The global reader to handle all terminal input.
var reader = bufio.NewReader(os.Stdin)

// --- Structs to match the JSON data structure ---
type Content struct {
	Greetings []string `json:"greetings"`
	Jokes     []string `json:"jokes"`
	Laughs    []string `json:"laughs"`
	Swears    []string `json:"swears"`
	Proverbs  []string `json:"proverbs"`
	Stages    Stages   `json:"stages"`
}

type Stages struct {
	Stage1 Stage1 `json:"stage1"`
	Stage2 Stage2 `json:"stage2"`
	Stage3 Stage3 `json:"stage3"`
	Stage4 Stage4 `json:"stage4"`
	Stage5 Stage5 `json:"stage5"`
	Stage6 map[string]string `json:"stage6"`
	Stage7 Stage7 `json:"stage7"`
	Stage8 Stage8 `json:"stage8"`
	Stage9 Stage9 `json:"stage9"`
	Stage10 Stage10 `json:"stage10"`
}

type Stage1 struct {
	NamePrompt string `json:"namePrompt"`
	Responses  struct {
		ShortName string `json:"shortName"`
		LongName  string `json:"longName"`
		Intro     string `json:"intro"`
	} `json:"responses"`
}

type Stage2 struct {
	AgePrompt  string     `json:"agePrompt"`
	AgeRanges  []AgeRange `json:"ageRanges"`
}

type AgeRange struct {
	Min  int    `json:"min"`
	Max  int    `json:"max"`
	Text string `json:"text"`
	Yes  string `json:"yes,omitempty"`
	No   string `json:"no,omitempty"`
}

type Stage3 struct {
	HeightPrompt  string       `json:"heightPrompt"`
	HeightRanges  []RangeText  `json:"heightRanges"`
}

type Stage4 struct {
	WeightPrompt  string       `json:"weightPrompt"`
	WeightRanges  []RangeTextVariants `json:"weightRanges"`
}

type RangeText struct {
	Min  int    `json:"min"`
	Max  int    `json:"max"`
	Text string `json:"text"`
}

type RangeTextVariants struct {
	Min  int    `json:"min"`
	Max  int    `json:"max"`
	Text string `json:"text"`
	Variants []string `json:"variants,omitempty"`
}


type Stage5 struct {
	Prompts []Prompt `json:"prompts"`
}

type Prompt struct {
	Text     string   `json:"text"`
	Yes      interface{}   `json:"yes,omitempty"`
	No       interface{}   `json:"no,omitempty"`
	Response string   `json:"response,omitempty"`
}

type Stage7 struct {
	HometownPrompt  string   `json:"hometownPrompt"`
	VowelResponses  map[string]string `json:"vowelResponses"`
	Conclusion      string   `json:"conclusion"`
}

type Stage8 struct {
	Intro       string `json:"intro"`
	GuessPrompt string `json:"guessPrompt"`
	InvalidGuess string `json:"invalidGuess"`
	TooLow      string `json:"tooLow"`
	TooLowFar   string `json:"tooLowFar"`
	TooHigh     string `json:"tooHigh"`
	TooHighFar  string `json:"tooHighFar"`
	OutOfBounds string `json:"outOfBounds"`
	Success     struct {
		VeryGood string `json:"veryGood"`
		Good     string `json:"good"`
		Average  string `json:"average"`
		Poor     string `json:"poor"`
		VeryPoor string `json:"veryPoor"`
		Terrible string `json:"terrible"`
	} `json:"success"`
}

type Stage9 struct {
	Prompts []string `json:"prompts"`
	Responses struct {
		Win     string `json:"win"`
		Cheating string `json:"cheating"`
		Equal    string `json:"equal"`
	} `json:"responses"`
}

type Stage10 struct {
	JokeIntro  string `json:"jokeIntro"`
	ExitPrompt string `json:"exitPrompt"`
}


// loadContent reads the JSON file and unmarshals it into the Content struct.
func loadContent() {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	byteValue, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	err = json.Unmarshal(byteValue, &content)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		os.Exit(1)
	}
}


// typewriterPrint simulates a typing effect by printing characters one by one.
func typewriterPrint(s string) {
	typingSpeed := 15 * time.Millisecond 
	for _, char := range s {
		fmt.Printf("%c", char)
		time.Sleep(typingSpeed)
	}
	fmt.Println()
}

// centerPrint calculates the padding and prints the text in the middle of the terminal.
func centerPrint(s string) {
	padding := (terminalWidth - countCharacters(s)) / 2
	if padding < 0 {
		padding = 0
	}
	fmt.Printf(strings.Repeat(" ", padding))
	typewriterPrint(s)
}

// blinkingCursor simulates a blinking cursor while the program "thinks."
func blinkingCursor(duration time.Duration) {
	blinkingSpeed := 500 * time.Millisecond 
	endTime := time.Now().Add(duration)
	for time.Now().Before(endTime) {
		fmt.Print("_")
		time.Sleep(blinkingSpeed)
		fmt.Print("\b \b")
		time.Sleep(blinkingSpeed)
	}
}

// userPrompt prints a separator and then the user prompt on the left.
func userPrompt(s string) {
	fmt.Println(ColorGreen + separator + ColorReset)
	fmt.Print(ColorGreen + s + ColorReset)
}

// getRandomInt returns a random integer up to the given maximum (exclusive).
func getRandomInt(max int) int {
	return rand.Intn(max)
}

// countCharacters counts the number of non-space characters in a string.
func countCharacters(s string) int {
	s = strings.ReplaceAll(s, ColorCyan, "")
	s = strings.ReplaceAll(s, ColorGreen, "")
	s = strings.ReplaceAll(s, ColorReset, "")
	return len(strings.ReplaceAll(s, " ", ""))
}

// sayJoke prints a random joke from a predefined list.
func sayJoke() {
	var jokeIndex int
	for {
		jokeIndex = getRandomInt(len(content.Jokes))
		if jokeIndex != previousJokeID {
			previousJokeID = jokeIndex
			break
		}
	}
	fmt.Print(ColorCyan + "..." + ColorReset)
	blinkingCursor(1 * time.Second)
	fmt.Println()
	centerPrint(ColorCyan + content.Jokes[jokeIndex] + ColorReset)
}

// laugh prints a random laughing phrase.
func laugh() {
	centerPrint(ColorCyan + content.Laughs[getRandomInt(len(content.Laughs))] + ColorReset)
}

// actDumb has a 50% chance of printing a "dumb" joke.
func actDumb() {
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + "\ngeri zekalı taklidi yap bakiim...\nTamam tamam bukadar yeter!!!\n" + ColorReset)
		laugh()
	}
}

// swear has a chance to print one or more rude phrases.
func swear() {
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + content.Swears[0] + ColorReset)
	}
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + content.Swears[1] + ColorReset)
	}
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + content.Swears[2] + ColorReset)
	}
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + content.Swears[3] + ColorReset)
	}
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + content.Swears[4] + ColorReset)
	}
}

// isVowel checks if a given character is a Turkish vowel.
func isVowel(r rune) bool {
	vowels := "aıueöüio"
	return strings.ContainsRune(vowels, r)
}

// stage10 concludes the game with a final joke.
func stage10() {
	centerPrint(ColorCyan + content.Stages.Stage10.JokeIntro + ColorReset)
	sayJoke()
	userPrompt(content.Stages.Stage10.ExitPrompt)
	reader.ReadString('\n')
	os.Exit(0)
}

// stage9 is the number guessing game where the computer guesses the user's number.
func stage9() {
	var guess int = getRandomInt(100) + 1
	upperLimit := 100
	lowerLimit := 1
	errorCount = 0
	guessCount := 0

	centerPrint(ColorCyan + content.Stages.Stage9.Prompts[0] + ColorReset)
	centerPrint(ColorCyan + content.Stages.Stage9.Prompts[1] + ColorReset)
	centerPrint(ColorCyan + content.Stages.Stage9.Prompts[2] + ColorReset)

	for {
		guessCount++
		fmt.Print(ColorCyan + "..." + ColorReset)
		blinkingCursor(500 * time.Millisecond)
		fmt.Println()
		
		centerPrint(ColorCyan + fmt.Sprintf(" %d  ??\n", guess) + ColorReset)
		userPrompt("? ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "y" {
			if upperLimit-1 == guess && lowerLimit+1 == guess {
				swear()
				errorCount++
				if errorCount > 5 {
					break
				}
			} else {
				lowerLimit = guess
				guess = getRandomInt(upperLimit-lowerLimit-1) + lowerLimit + 1
			}
		} else if input == "d" {
			if upperLimit-1 == guess && lowerLimit+1 == guess {
				swear()
				errorCount++
				if errorCount > 5 {
					break
				}
			} else {
				upperLimit = guess
				guess = getRandomInt(upperLimit-lowerLimit-1) + lowerLimit + 1
			}
		} else if input == "b" {
			break
		}
	}

	centerPrint(ColorCyan + fmt.Sprintf(content.Stages.Stage9.Responses.Win, guessCount) + ColorReset)
	if guessCount < score {
		centerPrint(ColorCyan + content.Stages.Stage9.Responses.Win + ColorReset)
	} else if guessCount > score {
		centerPrint(ColorCyan + content.Stages.Stage9.Responses.Cheating + ColorReset)
	} else {
		centerPrint(ColorCyan + content.Stages.Stage9.Responses.Equal + ColorReset)
	}

	stage10()
}

// stage8 is the number guessing game where the user guesses the computer's number.
func stage8() {
	target := getRandomInt(100) + 1
	var guess int
	guessCount := 0

	centerPrint(ColorCyan + fmt.Sprintf(content.Stages.Stage8.Intro, userName) + ColorReset)

	for {
		guessCount++
		userPrompt(content.Stages.Stage8.GuessPrompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var err error
		guess, err = strconv.Atoi(input)
		if err != nil {
			typewriterPrint(ColorCyan + content.Stages.Stage8.InvalidGuess + ColorReset)
			continue
		}

		if guess == target {
			var successMsg string
			if guessCount <= 3 {
				successMsg = content.Stages.Stage8.Success.VeryGood
			} else if guessCount <= 5 {
				successMsg = content.Stages.Stage8.Success.Good
			} else if guessCount <= 10 {
				successMsg = content.Stages.Stage8.Success.Average
			} else if guessCount <= 20 {
				successMsg = content.Stages.Stage8.Success.Poor
			} else if guessCount <= 30 {
				successMsg = content.Stages.Stage8.Success.VeryPoor
			} else {
				successMsg = content.Stages.Stage8.Success.Terrible
			}
			centerPrint(ColorCyan + fmt.Sprintf(successMsg, guessCount) + ColorReset)
			score = guessCount
			stage9()
			break
		} else {
			if guess < 1 || guess > 100 {
				centerPrint(ColorCyan + content.Stages.Stage8.OutOfBounds + ColorReset)
			} else if guess < target {
				if target-guess > 20 {
					centerPrint(ColorCyan + content.Stages.Stage8.TooLowFar + ColorReset)
				} else {
					centerPrint(ColorCyan + content.Stages.Stage8.TooLow + ColorReset)
				}
			} else { // guess > target
				if guess-target > 20 {
					centerPrint(ColorCyan + content.Stages.Stage8.TooHighFar + ColorReset)
				} else {
					centerPrint(ColorCyan + content.Stages.Stage8.TooHigh + ColorReset)
				}
			}
		}
	}
}

// stage7 asks for the user's hometown and responds based on the last vowel.
func stage7() {
	centerPrint(ColorCyan + fmt.Sprintf(content.Stages.Stage7.HometownPrompt, userName) + ColorReset)
	userPrompt("? ")
	hometown, _ := reader.ReadString('\n')
	hometown = strings.TrimSpace(hometown)

	runes := []rune(hometown)
	lastVowel := ' '
	foundVowel := false

	for i := len(runes) - 1; i >= 0; i-- {
		if isVowel(runes[i]) {
			lastVowel = runes[i]
			foundVowel = true
			break
		}
	}

	if foundVowel {
		fmt.Print(ColorCyan + "..." + ColorReset)
		blinkingCursor(1 * time.Second)
		fmt.Println()
		
		switch lastVowel {
		case 'u', 'o':
			centerPrint(ColorCyan + fmt.Sprintf(content.Stages.Stage7.VowelResponses["u o"], hometown, hometown) + ColorReset)
		case 'ü', 'ö':
			centerPrint(ColorCyan + fmt.Sprintf(content.Stages.Stage7.VowelResponses["ü ö"], hometown) + ColorReset)
		case 'a', 'ı':
			centerPrint(ColorCyan + fmt.Sprintf(content.Stages.Stage7.VowelResponses["a ı"], hometown) + ColorReset)
		case 'e', 'i':
			centerPrint(ColorCyan + fmt.Sprintf(content.Stages.Stage7.VowelResponses["e i"], hometown) + ColorReset)
		}
	}

	laugh()
	centerPrint(ColorCyan + fmt.Sprintf(content.Stages.Stage7.Conclusion, userName) + ColorReset)

	stage8()
}

// stage6 prints a joke and a proverb.
func stage6() {
	centerPrint(ColorCyan + "bak sana şindi konuyla ilgili bir fıkra..." + ColorReset)
	sayJoke()
	laugh()
	centerPrint(ColorCyan + fmt.Sprintf("\n%s\n", content.Proverbs[getRandomInt(len(content.Proverbs))]) + ColorReset)
	laugh()
	centerPrint("") 
	stage7()
}

// stage5 contains a series of random questions.
func stage5() {
	// Question 1: Eyes
	if getRandomInt(2) == 1 {
		prompt := content.Stages.Stage5.Prompts[0]
		centerPrint(ColorCyan + fmt.Sprintf(prompt.Text, userName) + ColorReset)
		userPrompt("(e/h)? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			centerPrint(ColorCyan + prompt.Yes.(string) + ColorReset)
			laugh()
		} else {
			centerPrint(ColorCyan + prompt.No.(string) + ColorReset)
			laugh()
		}
	}

	// Question 2: Money
	if getRandomInt(2) == 1 {
		prompt := content.Stages.Stage5.Prompts[1]
		centerPrint(ColorCyan + fmt.Sprintf(prompt.Text, userName) + ColorReset)
		userPrompt("(e/h)? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			centerPrint(ColorCyan + prompt.Yes.(string) + ColorReset)
			laugh()
		} else {
			centerPrint(ColorCyan + prompt.No.(string) + ColorReset)
			laugh()
		}
	}

	// Question 3: Name Origin
	if getRandomInt(2) == 1 {
		prompt := content.Stages.Stage5.Prompts[2]
		centerPrint(ColorCyan + fmt.Sprintf(prompt.Text, userName) + ColorReset)
		userPrompt("? ")
		reader.ReadString('\n') 
		centerPrint(ColorCyan + prompt.Response + ColorReset)
		laugh()
	}

	// Question 4: Holding a number
	if getRandomInt(2) == 1 {
		prompt := content.Stages.Stage5.Prompts[3]
		centerPrint(ColorCyan + fmt.Sprintf(prompt.Text, userName) + ColorReset)
		userPrompt("(e/h)? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			centerPrint(ColorCyan + prompt.Yes.(string) + ColorReset)
			laugh()
		} else {
			centerPrint(ColorCyan + prompt.No.(string) + ColorReset)
			laugh()
		}
	}

	// Question 5: Nickname (This prompt needs special handling as it has a mix of AI intro and user prompt)
	if getRandomInt(2) == 1 {
		runes := []rune(userName)
		var nickname string
		if len(runes) >= 2 && isVowel(runes[1]) {
			nickname = fmt.Sprintf("%c%c%coş", runes[0], runes[1], runes[2])
		} else if len(runes) >= 2 {
			nickname = fmt.Sprintf("%c%coş", runes[0], runes[1])
		}
		if nickname != "" {
			centerPrint(ColorCyan + fmt.Sprintf("\n%s, sana kısaca %s diyebilirmiyim??\n", userName, nickname) + ColorReset)
			userPrompt("(e/h)? ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(strings.ToLower(input))
			if input == "e" {
				centerPrint(ColorCyan + "iyi... ama ben demek istemiyorum!" + ColorReset)
				laugh()
			} else {
				centerPrint(ColorCyan + fmt.Sprintf("%s! %s! %s!\n", nickname, nickname, nickname) + ColorReset)
				laugh()
			}
		}
	}
	
	// Question 6: How are you?
	if getRandomInt(2) == 1 {
		prompt := content.Stages.Stage5.Prompts[4]
		centerPrint(ColorCyan + fmt.Sprintf(prompt.Text, userName) + ColorReset)
		userPrompt("(e/h)? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			randChoice := getRandomInt(3)
			centerPrint(ColorCyan + fmt.Sprintf(prompt.Yes.([]interface{})[randChoice].(string), userName) + ColorReset)
		} else {
			randChoice := getRandomInt(3)
			if randChoice < 2 {
				centerPrint(ColorCyan + prompt.No.([]interface{})[randChoice].(string) + ColorReset)
			} else {
				centerPrint(ColorCyan + fmt.Sprintf(prompt.No.([]interface{})[randChoice].(string), userName) + ColorReset)
				reader.ReadString('\n')
				centerPrint(ColorCyan + prompt.No.([]interface{})[3].(string) + ColorReset)
			}
		}
		laugh()
	}

	// Question 7: Student
	if getRandomInt(2) == 1 {
		prompt := content.Stages.Stage5.Prompts[5]
		centerPrint(ColorCyan + fmt.Sprintf(prompt.Text, userName) + ColorReset)
		userPrompt("? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			randChoice := getRandomInt(2)
			centerPrint(ColorCyan + prompt.Yes.([]interface{})[randChoice].(string) + ColorReset)
		} else {
			randChoice := getRandomInt(2)
			if randChoice == 0 {
				centerPrint(ColorCyan + prompt.No.([]interface{})[randChoice].(string) + ColorReset)
			} else {
				userPrompt(prompt.No.([]interface{})[1].(string))
				reader.ReadString('\n')
				centerPrint(ColorCyan + prompt.No.([]interface{})[2].(string) + ColorReset)
			}
		}
		laugh()
	}
	stage6()
}

// stage4 asks for the user's weight and responds accordingly.
func stage4() {
	var weight int
	for {
		userPrompt(content.Stages.Stage4.WeightPrompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		weight, err = strconv.Atoi(input)
		if err != nil {
			typewriterPrint(ColorCyan + "Geçersiz giriş. Lütfen bir sayı girin." + ColorReset)
			continue
		}

		fmt.Print(ColorCyan + "..." + ColorReset)
		blinkingCursor(1 * time.Second)
		fmt.Println()
		
		for _, r := range content.Stages.Stage4.WeightRanges {
			if weight >= r.Min && weight <= r.Max {
				if len(r.Variants) > 0 {
					centerPrint(ColorCyan + r.Variants[getRandomInt(len(r.Variants))] + ColorReset)
				} else {
					centerPrint(ColorCyan + r.Text + ColorReset)
				}
				actDumb()
				break
			}
		}
		centerPrint("") 
		break
	}
	stage5()
}

// stage3 asks for the user's height and responds accordingly.
func stage3() {
	userPrompt(content.Stages.Stage3.HeightPrompt)
	var height int
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		height, err = strconv.Atoi(input)
		if err != nil {
			typewriterPrint(ColorCyan + "Geçersiz giriş. Lütfen bir sayı girin." + ColorReset)
			continue
		}

		fmt.Print(ColorCyan + "..." + ColorReset)
		blinkingCursor(1 * time.Second)
		fmt.Println()
		
		for _, r := range content.Stages.Stage3.HeightRanges {
			if height >= r.Min && height <= r.Max {
				centerPrint(ColorCyan + r.Text + ColorReset)
				break
			}
		}
		centerPrint("") 
		break
	}
	stage4()
}

// stage2 asks for the user's age and responds accordingly.
func stage2() {
	userPrompt(content.Stages.Stage2.AgePrompt)
	var age int
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		age, err = strconv.Atoi(input)
		if err != nil {
			typewriterPrint(ColorCyan + "Geçersiz giriş. Lütfen bir sayı girin." + ColorReset)
			continue
		}
		
		fmt.Print(ColorCyan + "..." + ColorReset)
		blinkingCursor(1 * time.Second)
		fmt.Println()

		for _, r := range content.Stages.Stage2.AgeRanges {
			if age >= r.Min && age <= r.Max {
				if r.Yes != "" || r.No != "" {
					userPrompt(r.Text)
					choice, _ := reader.ReadString('\n')
					choice = strings.TrimSpace(strings.ToLower(choice))
					if choice == "e" {
						centerPrint(ColorCyan + r.Yes + ColorReset)
					} else {
						centerPrint(ColorCyan + r.No + ColorReset)
					}
				} else {
					centerPrint(ColorCyan + r.Text + ColorReset)
				}
				break
			}
		}
		centerPrint("")
		break
	}
	stage3()
}

// stage1 asks for the user's name and starts the conversation.
func stage1() {
	userPrompt(content.Stages.Stage1.NamePrompt)
	input, _ := reader.ReadString('\n')
	userName = strings.TrimSpace(input)

	fmt.Print(ColorCyan + "..." + ColorReset)
	blinkingCursor(1 * time.Second)
	fmt.Println()

	charCount := countCharacters(userName)

	if charCount <= 2 {
		centerPrint(ColorCyan + fmt.Sprintf(content.Stages.Stage1.Responses.ShortName, charCount, userName[0], userName[0], userName) + ColorReset)
	} else if charCount >= 8 {
		centerPrint(ColorCyan + content.Stages.Stage1.Responses.LongName + ColorReset)
		laugh()
	}

	centerPrint(ColorCyan + fmt.Sprintf(content.Stages.Stage1.Responses.Intro, userName) + ColorReset)
	stage2()
}

// stage0 is the initial welcome message.
func stage0() {
	for _, greeting := range content.Greetings {
		centerPrint(ColorCyan + greeting + ColorReset)
	}
	stage1()
}

// main is the entry point of the Go application.
func main() {
	rand.Seed(time.Now().UnixNano())
	loadContent()
	stage0()
}

