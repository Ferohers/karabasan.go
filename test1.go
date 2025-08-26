package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// ANSI escape codes for coloring and text formatting.
const (
	ColorCyan   = "\033[36m"
	ColorGreen  = "\033[32m"
	ColorReset  = "\033[0m"

	// Separator width is now a constant, allowing for dynamic separators.
	separatorWidth = 80
	// A new, clean prompt symbol for user input.
	promptSymbol = "> "
	
	// Assuming a standard terminal width for centering text.
	terminalWidth = 80
)

// reader is the global reader to handle terminal input.
var reader = bufio.NewReader(os.Stdin)

// typewriterPrint simulates a typing effect by printing characters one by one.
// This function creates the dynamic, "thinking" feel of a modern chat interface.
func typewriterPrint(s string) {
	typingSpeed := 15 * time.Millisecond 
	for _, char := range s {
		fmt.Printf("%c", char)
		time.Sleep(typingSpeed)
	}
	fmt.Println()
}

// centerPrint calculates the necessary padding and prints text in the middle of the terminal.
func centerPrint(s string) {
	padding := (terminalWidth - countCharacters(s)) / 2
	if padding < 0 {
		padding = 0
	}
	fmt.Printf(strings.Repeat(" ", padding))
	typewriterPrint(s)
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
// The user will type their input on a separate line below the prompt.
func userPrompt(s string) {
	fmt.Println(ColorGreen + strings.Repeat("-", separatorWidth) + ColorReset)
	fmt.Println(ColorGreen + s + ColorReset)
	fmt.Print(ColorGreen + promptSymbol + ColorReset)
}

// countCharacters counts the number of visible characters in a string, ignoring ANSI codes.
func countCharacters(s string) int {
	// Only remove the color codes from the string before counting.
	s = strings.ReplaceAll(s, ColorCyan, "")
	s = strings.ReplaceAll(s, ColorGreen, "")
	s = strings.ReplaceAll(s, ColorReset, "")
	return len(strings.ReplaceAll(s, " ", ""))
}

// stage0 is the initial welcome and introduction.
func stage0() {
	// AI's first message, centered and with a typing effect.
	centerPrint(ColorCyan + "Merhaba, hoş geldin." + ColorReset)
	time.Sleep(1 * time.Second)

	// Thinking effect.
	fmt.Print(ColorCyan + "..." + ColorReset)
	blinkingCursor(1 * time.Second)
	fmt.Println()

	centerPrint(ColorCyan + "Ben yeni nesil bir terminal arayüzüyüm." + ColorReset)
	time.Sleep(1 * time.Second)

	// User input prompt.
	userPrompt("Adın ne?")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// AI responds to user input, demonstrating the chat flow.
	fmt.Print(ColorCyan + "..." + ColorReset)
	blinkingCursor(1 * time.Second)
	fmt.Println()
	centerPrint(ColorCyan + fmt.Sprintf("Tanıştığıma memnun oldum, %s. Nasıl yardımcı olabilirim?", name) + ColorReset)
	
	userPrompt("...") // An open-ended prompt for the user to continue the conversation.
	reader.ReadString('\n')
	
	// Exit the program after the conversation ends.
	fmt.Println("Programdan çıkılıyor...")
}

// main is the entry point of the Go application.
func main() {
	stage0()
}

