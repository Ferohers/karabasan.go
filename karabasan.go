package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/term"
)

// ANSI escape codes for coloring and text formatting.
const (
	ColorCyan    = "\033[36m"
	ColorGreen   = "\033[32m"
	ColorMagenta = "\033[35m" // New color for AI responses
	ColorReset   = "\033[0m"

	// A new, clean prompt symbol for user input.
	promptSymbol = "> "
)

// Variables to store the terminal size and conversational content.
var (
	userName       string
	score          int
	previousJokeID int = -1
	errorCount     int
	terminalWidth  = 80
	separatorWidth = 80
	reader         = bufio.NewReader(os.Stdin)
	content        Content
)

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
	Stage1 Stage1  `json:"stage1"`
	Stage2 Stage2  `json:"stage2"`
	Stage3 Stage3  `json:"stage3"`
	Stage4 Stage4  `json:"stage4"`
	Stage5 Stage5  `json:"stage5"`
	Stage6 Stage6  `json:"stage6"`
	Stage7 Stage7  `json:"stage7"`
	Stage8 Stage8  `json:"stage8"`
	Stage9 Stage9  `json:"stage9"`
	Stage10 Stage10 `json:"stage10"`
}

type Stage1 struct {
	NamePrompt string `json:"namePrompt"`
	Responses  struct {
		Intro     string `json:"intro"`
		ShortName string `json:"shortName"`
		LongName  string `json:"longName"`
	} `json:"responses"`
}

type Stage2 struct {
	AgePrompt    string     `json:"agePrompt"`
	AgeResponse  string     `json:"ageResponse"`
	InvalidInput string     `json:"invalidInput"`
	AgeRanges    []AgeRange `json:"ageRanges"`
}

type AgeRange struct {
	Min  int    `json:"min"`
	Max  int    `json:"max"`
	Text string `json:"text"`
	Yes  string `json:"yes,omitempty"`
	No   string `json:"no,omitempty"`
}

type Stage3 struct {
	HeightPrompt   string      `json:"heightPrompt"`
	HeightResponse string      `json:"heightResponse"`
	InvalidInput   string      `json:"invalidInput"`
	HeightRanges   []RangeText `json:"heightRanges"`
}

type RangeText struct {
	Min  int    `json:"min"`
	Max  int    `json:"max"`
	Text string `json:"text"`
}

type Stage4 struct {
	WeightPrompt   string              `json:"weightPrompt"`
	WeightResponse string              `json:"weightResponse"`
	InvalidInput   string              `json:"invalidInput"`
	WeightRanges   []RangeTextVariants `json:"weightRanges"`
}

type RangeTextVariants struct {
	Min      int      `json:"min"`
	Max      int      `json:"max"`
	Text     string   `json:"text,omitempty"`
	Variants []string `json:"variants,omitempty"`
}

type Stage5 struct {
	Prompts []Prompt `json:"prompts"`
}

type Prompt struct {
	Text     string      `json:"text"`
	Yes      interface{} `json:"yes,omitempty"`
	No       interface{} `json:"no,omitempty"`
	Response string      `json:"response,omitempty"`
}

type Stage6 struct {
	JokeIntro string `json:"jokeIntro"`
}

type Stage7 struct {
	HometownPrompt string            `json:"hometownPrompt"`
	VowelResponses map[string]string `json:"vowelResponses"`
	Conclusion     string            `json:"conclusion"`
}

type Stage8 struct {
	Intro        string `json:"intro"`
	GuessPrompt  string `json:"guessPrompt"`
	InvalidGuess string `json:"invalidGuess"`
	TooLow       string `json:"tooLow"`
	TooLowFar    string `json:"tooLowFar"`
	TooHigh      string `json:"tooHigh"`
	TooHighFar   string `json:"tooHighFar"`
	OutOfBounds  string `json:"outOfBounds"`
	Success      struct {
		VeryGood string `json:"veryGood"`
		Good     string `json:"good"`
		Average  string `json:"average"`
		Poor     string `json:"poor"`
		VeryPoor string `json:"veryPoor"`
		Terrible string `json:"terrible"`
	} `json:"success"`
}

type Stage9 struct {
	Prompts   []string `json:"prompts"`
	Responses struct {
		Win      string `json:"win"`
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
		fmt.Println("Error opening file: data.json. Please make sure the file exists and is in the same directory.")
		os.Exit(1)
	}
	defer jsonFile.Close()

	byteValue, _ := os.ReadFile("data.json")
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

// centerPrint calculates the necessary padding and prints text in the middle of the terminal.
// This function is updated to correctly handle multi-line strings by centering each line individually.
func centerPrint(s string) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		// Remove color codes for accurate width calculation
		cleanLine := strings.ReplaceAll(line, ColorCyan, "")
		cleanLine = strings.ReplaceAll(cleanLine, ColorGreen, "")
		cleanLine = strings.ReplaceAll(cleanLine, ColorMagenta, "")
		cleanLine = strings.ReplaceAll(cleanLine, ColorReset, "")

		padding := (terminalWidth - len(cleanLine)) / 2
		if padding < 0 {
			padding = 0
		}
		fmt.Printf(strings.Repeat(" ", padding))
		typewriterPrint(line)
	}
}

// aiResponse is a new function dedicated to AI conversational responses.
func aiResponse(s string) {
	fmt.Print(ColorCyan + "..." + ColorReset)
	blinkingCursor(1 * time.Second)
	fmt.Println()
	centerPrint(ColorMagenta + s + ColorReset)
}

// blinkingCursor simulates a blinking cursor to represent the program "thinking."
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

// userPrompt prints a separator and a clean prompt for the user.
func userPrompt(s string) {
	fmt.Println(ColorGreen + strings.Repeat("-", separatorWidth) + ColorReset)
	fmt.Println(ColorGreen + s + ColorReset)
	fmt.Print(ColorGreen + promptSymbol + ColorReset)
}

// countCharacters counts the number of visible characters in a string, ignoring ANSI codes.
func countCharacters(s string) int {
	s = strings.ReplaceAll(s, ColorCyan, "")
	s = strings.ReplaceAll(s, ColorGreen, "")
	s = strings.ReplaceAll(s, ColorMagenta, "")
	s = strings.ReplaceAll(s, ColorReset, "")
	return len(strings.ReplaceAll(s, " ", ""))
}

// getRandomInt returns a random integer up to the given maximum (exclusive).
func getRandomInt(max int) int {
    return rand.Intn(max)
}

// isVowel checks if a given character is a Turkish vowel.
func isVowel(r rune) bool {
    vowels := "aıueöüio"
    return strings.ContainsRune(vowels, r)
}

// sayJoke prints a random joke from a predefined list.
func sayJoke() {
    jokes := []string{
        "adamın biri soğuk çay istemiş...\nçaycı çayı getirmiş...\nadam da 'ISIT DA İÇELİM KARDEŞİM!' demiş!",
        "2 laz kuş avlamadaymış...\nbiri 'niye avlanamıyoz' diye dert yanmış...\nöbürü: 'BENCE KÖPEĞİ DAHA YUKARI ATMALIYIZ!",
        "bir grup laz yürüyen merdivenle çıkarken\nelektrikler kesilmiş...\n2 saat süreyle mahsur kalmışlar!!!",
        "30 yaşındaki bir Alman koskoca bir uçağı...\ntek eliyle kaldırmış..\nadam PİLOTMUŞ lan PİLOT!",
        "Temelle Dursun soygundadırlar...\nkaçarlarken polis arkalarından bağırır:\n'DUR KAÇMA OROSPU ÇOCUĞU!!'\nTemel Dursun'a dönerek:\n'Sen kaç! beni tanıdı!'",
    }
    var jokeIndex int
    for {
        jokeIndex = getRandomInt(len(jokes))
        if jokeIndex != previousJokeID {
            previousJokeID = jokeIndex
            break
        }
    }
    aiResponse(jokes[jokeIndex])
}

// laugh prints a random laughing phrase.
func laugh() {
    laughs := []string{
        "eki!eki!eki! köh!köh!köh! ayy nekadar neşeliyim!!",
        "neee? hahhahahahhahhhhayyyy!! kafadan kopardım gene!!   hehe!",
        "kah!keh!koh!küh! hahahahaha!!! hihihihi!! ve de hohoho!",
        "he he he he...",
        "hahahaha!! ay ben ölmiiim emi!",
    }
    aiResponse(laughs[getRandomInt(len(laughs))])
}

// actDumb has a 50% chance of printing a "dumb" joke.
func actDumb() {
    if getRandomInt(2) == 1 {
        aiResponse("\ngeri zekalı taklidi yap bakiim...\nTamam tamam bukadar yeter!!!\n")
        laugh()
    }
}

// swear has a chance to print one or more rude phrases.
func swear() {
    swears := []string{
        "EEE! mına korum böyle oyunun!! yıkıl köpek!",
        "bana bak! seni adam yerine koyduk karşımıza aldık,.. tööbe tööbee",
        "OHA! OHA! kırsaydın klavyeyi!!",
        "doğru oyna orospu!",
        "GÖT!",
    }
    if getRandomInt(2) == 1 {
        aiResponse(swears[0])
    }
    if getRandomInt(2) == 1 {
        aiResponse(swears[1])
    }
    if getRandomInt(2) == 1 {
        aiResponse(swears[2])
    }
    if getRandomInt(2) == 1 {
        aiResponse(swears[3])
    }
    if getRandomInt(2) == 1 {
        aiResponse(swears[4])
    }
}

// stage10 concludes the game with a final joke.
func stage10() {
    aiResponse(content.Stages.Stage10.JokeIntro)
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
    aiResponse(content.Stages.Stage9.Prompts[0])
    aiResponse(content.Stages.Stage9.Prompts[1])
    aiResponse(content.Stages.Stage9.Prompts[2])
    for {
        guessCount++
        aiResponse(fmt.Sprintf(" %d  ??\n", guess))
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
    
    // Fixed: The final response is now handled in a single, cohesive block.
    if guessCount < score {
        aiResponse(fmt.Sprintf(content.Stages.Stage9.Responses.Win, guessCount))
    } else if guessCount > score {
        aiResponse(content.Stages.Stage9.Responses.Cheating)
    } else {
        aiResponse(content.Stages.Stage9.Responses.Equal)
    }

    stage10()
}

// stage8 is the number guessing game where the user guesses the computer's number.
func stage8() {
    target := getRandomInt(100) + 1
    var guess int
    guessCount := 0
    aiResponse(fmt.Sprintf(content.Stages.Stage8.Intro, userName))
    for {
        guessCount++
        userPrompt(content.Stages.Stage8.GuessPrompt)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        var err error
        guess, err = strconv.Atoi(input)
        if err != nil {
            aiResponse(content.Stages.Stage8.InvalidGuess)
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
            aiResponse(fmt.Sprintf(successMsg, guessCount))
            score = guessCount
            stage9()
            break
        } else {
            if guess < 1 || guess > 100 {
                aiResponse(content.Stages.Stage8.OutOfBounds)
            } else if guess < target {
                if target-guess > 20 {
                    aiResponse(content.Stages.Stage8.TooLowFar)
                } else {
                    aiResponse(content.Stages.Stage8.TooLow)
                }
            } else { // guess > target
                if guess-target > 20 {
                    aiResponse(content.Stages.Stage8.TooHighFar)
                } else {
                    aiResponse(content.Stages.Stage8.TooHigh)
                }
            }
        }
    }
}

// stage7 asks for the user's hometown and responds based on the last vowel.
func stage7() {
    userPrompt(fmt.Sprintf(content.Stages.Stage7.HometownPrompt, userName))
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
        switch lastVowel {
        case 'u', 'o':
            aiResponse(fmt.Sprintf(content.Stages.Stage7.VowelResponses["u o"], hometown, hometown))
        case 'ü', 'ö':
            aiResponse(fmt.Sprintf(content.Stages.Stage7.VowelResponses["ü ö"], hometown))
        case 'a', 'ı':
            aiResponse(fmt.Sprintf(content.Stages.Stage7.VowelResponses["a ı"], hometown))
        case 'e', 'i':
            aiResponse(fmt.Sprintf(content.Stages.Stage7.VowelResponses["e i"], hometown))
        }
    }
    laugh()
    aiResponse(fmt.Sprintf(content.Stages.Stage7.Conclusion, userName))
    stage8()
}

// stage6 prints a joke and a proverb.
func stage6() {
    aiResponse(content.Stages.Stage6.JokeIntro)
    sayJoke()
    laugh()
    proverbs := []string{
        "yani sakla samanı gelir zamanı.",
        "yani arkadaşlarımızı dikkatli seçmemiz lazım.",
        "buradan alınacak ders: Göte giren şemsiye açılmaz..",
    }
    aiResponse(fmt.Sprintf("\n%s\n", proverbs[getRandomInt(len(proverbs))]))
    laugh()
    centerPrint("")
    stage7()
}

// stage5 contains a series of random questions.
func stage5() {
    // Question 1: Eyes
    if getRandomInt(2) == 1 {
        prompt := content.Stages.Stage5.Prompts[0]
        userPrompt(fmt.Sprintf(prompt.Text, userName))
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(strings.ToLower(input))
        if input == "e" {
            aiResponse(prompt.Yes.(string))
            laugh()
        } else {
            aiResponse(prompt.No.(string))
            laugh()
        }
    }
    // Question 2: Money
    if getRandomInt(2) == 1 {
        prompt := content.Stages.Stage5.Prompts[1]
        userPrompt(fmt.Sprintf(prompt.Text, userName))
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(strings.ToLower(input))
        if input == "e" {
            aiResponse(prompt.Yes.(string))
            laugh()
        } else {
            aiResponse(prompt.No.(string))
            laugh()
        }
    }
    // Question 3: Name Origin
    if getRandomInt(2) == 1 {
        prompt := content.Stages.Stage5.Prompts[2]
        aiResponse(fmt.Sprintf(prompt.Text, userName))
        userPrompt("? ")
        reader.ReadString('\n') 
        aiResponse(prompt.Response)
        laugh()
    }
    // Question 4: Holding a number
    if getRandomInt(2) == 1 {
        prompt := content.Stages.Stage5.Prompts[3]
        aiResponse(fmt.Sprintf(prompt.Text, userName))
        userPrompt(prompt.Text)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(strings.ToLower(input))
        if input == "e" {
            aiResponse(prompt.Yes.(string))
            laugh()
        } else {
            aiResponse(prompt.No.(string))
            laugh()
        }
    }
    // Question 5: Nickname
    if getRandomInt(2) == 1 {
        runes := []rune(userName)
        var nickname string
        if len(runes) >= 2 && isVowel(runes[1]) {
            nickname = fmt.Sprintf("%c%c%coş", runes[0], runes[1], runes[2])
        } else if len(runes) >= 2 {
            nickname = fmt.Sprintf("%c%coş", runes[0], runes[1])
        }
        if nickname != "" {
            aiResponse(fmt.Sprintf("\n%s, sana kısaca %s diyebilirmiyim??\n", userName, nickname))
            userPrompt("? ")
            input, _ := reader.ReadString('\n')
            input = strings.TrimSpace(strings.ToLower(input))
            if input == "e" {
                aiResponse("iyi... ama ben demek istemiyorum!")
                laugh()
            } else {
                aiResponse(fmt.Sprintf("%s! %s! %s!\n", nickname, nickname, nickname))
                laugh()
            }
        }
    }
    // Question 6: How are you?
    if getRandomInt(2) == 1 {
        prompt := content.Stages.Stage5.Prompts[4]
        aiResponse(fmt.Sprintf(prompt.Text, userName))
        userPrompt("? ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(strings.ToLower(input))
        if input == "e" {
            randChoice := getRandomInt(3)
            if randChoice == 0 {
                aiResponse(prompt.Yes.([]interface{})[0].(string))
            } else if randChoice == 1 {
                aiResponse(fmt.Sprintf(prompt.Yes.([]interface{})[1].(string), userName))
            } else {
                aiResponse(fmt.Sprintf(prompt.Yes.([]interface{})[2].(string), userName))
            }
        } else {
            randChoice := getRandomInt(3)
            if randChoice == 0 {
                aiResponse(prompt.No.([]interface{})[0].(string))
            } else if randChoice == 1 {
                aiResponse(prompt.No.([]interface{})[1].(string))
            } else {
                aiResponse(fmt.Sprintf(prompt.No.([]interface{})[2].(string), userName))
                reader.ReadString('\n') 
                aiResponse(prompt.No.([]interface{})[3].(string))
            }
        }
        laugh()
    }
    // Question 7: Student
    if getRandomInt(2) == 1 {
        prompt := content.Stages.Stage5.Prompts[5]
        aiResponse(fmt.Sprintf(prompt.Text, userName))
        userPrompt("? ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(strings.ToLower(input))
        if input == "e" {
            randChoice := getRandomInt(2)
            aiResponse(prompt.Yes.([]interface{})[randChoice].(string))
        } else {
            randChoice := getRandomInt(2)
            if randChoice == 0 {
                aiResponse(prompt.No.([]interface{})[0].(string))
            } else {
                userPrompt(prompt.No.([]interface{})[1].(string))
                reader.ReadString('\n') 
                aiResponse(prompt.No.([]interface{})[2].(string))
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
            aiResponse(content.Stages.Stage4.InvalidInput)
            continue
        }
        aiResponse(fmt.Sprintf(content.Stages.Stage4.WeightResponse, weight))
        if weight <= 39 {
            aiResponse(content.Stages.Stage4.WeightRanges[0].Text)
            actDumb()
        } else if weight >= 40 && weight <= 59 {
            aiResponse(content.Stages.Stage4.WeightRanges[1].Text)
            actDumb()
        } else if weight >= 60 && weight <= 79 {
            aiResponse(content.Stages.Stage4.WeightRanges[2].Text)
            actDumb()
        } else if weight >= 80 && weight <= 99 {
            randChoice := getRandomInt(len(content.Stages.Stage4.WeightRanges[3].Variants))
            aiResponse(content.Stages.Stage4.WeightRanges[3].Variants[randChoice])
            actDumb()
        } else if weight >= 100 {
            aiResponse(content.Stages.Stage4.WeightRanges[4].Text)
            actDumb()
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
            aiResponse(content.Stages.Stage3.InvalidInput)
            continue
        }
        aiResponse(fmt.Sprintf(content.Stages.Stage3.HeightResponse, height))
        if height <= 99 {
            aiResponse(content.Stages.Stage3.HeightRanges[0].Text)
        } else if height >= 100 && height <= 149 {
            aiResponse(content.Stages.Stage3.HeightRanges[1].Text)
        } else if height >= 150 && height <= 169 {
            aiResponse(content.Stages.Stage3.HeightRanges[2].Text)
        } else if height >= 170 && height <= 189 {
            aiResponse(content.Stages.Stage3.HeightRanges[3].Text)
        } else if height >= 190 && height <= 209 {
            aiResponse(content.Stages.Stage3.HeightRanges[4].Text)
        } else if height >= 210 {
            aiResponse(content.Stages.Stage3.HeightRanges[5].Text)
            continue
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
            aiResponse(content.Stages.Stage2.InvalidInput)
            continue
        }
        aiResponse(fmt.Sprintf(content.Stages.Stage2.AgeResponse, age))
        if age <= 4 {
            aiResponse(content.Stages.Stage2.AgeRanges[0].Text)
        } else if age >= 5 && age <= 9 {
            userPrompt(content.Stages.Stage2.AgeRanges[1].Text)
            choice, _ := reader.ReadString('\n')
            choice = strings.TrimSpace(strings.ToLower(choice))
            if choice == "e" {
                aiResponse(content.Stages.Stage2.AgeRanges[1].Yes)
            } else {
                aiResponse(content.Stages.Stage2.AgeRanges[1].No)
            }
        } else if age >= 10 && age <= 17 {
            aiResponse(content.Stages.Stage2.AgeRanges[2].Text)
        } else if age >= 18 && age <= 24 {
            userPrompt(content.Stages.Stage2.AgeRanges[3].Text)
            choice, _ := reader.ReadString('\n')
            choice = strings.TrimSpace(strings.ToLower(choice))
            if choice == "e" {
                aiResponse(content.Stages.Stage2.AgeRanges[3].Yes)
            } else {
                aiResponse(content.Stages.Stage2.AgeRanges[3].No)
            }
        } else if age >= 25 && age <= 39 {
            aiResponse(content.Stages.Stage2.AgeRanges[4].Text)
        } else if age >= 40 && age <= 59 {
            aiResponse(content.Stages.Stage2.AgeRanges[5].Text)
        } else if age >= 60 && age <= 98 {
            aiResponse(content.Stages.Stage2.AgeRanges[6].Text)
        } else if age >= 99 {
            aiResponse(content.Stages.Stage2.AgeRanges[7].Text)
            continue
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
    aiResponse(fmt.Sprintf(content.Stages.Stage1.Responses.Intro, userName))
    stage2()
}

// stage0 is the initial welcome and introduction.
func stage0() {
	fmt.Println()
	centerPrint(ColorCyan + "Merhaba, hoş geldin." + ColorReset)
	time.Sleep(1 * time.Second)
	centerPrint(ColorCyan + "Ben yeni nesil bir terminal arayüzüyüm." + ColorReset)
	time.Sleep(1 * time.Second)
	stage1()
}

// main is the entry point of the Go application.
func main() {
	rand.Seed(time.Now().UnixNano())
	width, _, err := term.GetSize(int(os.Stdin.Fd()))
	if err == nil {
		terminalWidth = width
		separatorWidth = width
	}
	loadContent()
	stage0()
}
