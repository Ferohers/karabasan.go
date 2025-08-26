package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	userName       string
	score          int
	previousJokeID int = -1
	errorCount     int
)

// ANSI escape codes for coloring terminal output.
const (
	ColorCyan   = "\033[36m"
	ColorGreen  = "\033[32m"
	ColorReset  = "\033[0m"
	
	// Assuming a standard terminal width for centering text.
	// You might need to adjust this value for different terminal sizes.
	terminalWidth = 80
	// Separator for the user input area.
	separator = "================================================================================"
)

// The global reader to handle all terminal input.
var reader = bufio.NewReader(os.Stdin)

// typewriterPrint simulates a typing effect by printing characters one by one.
// Bu fonksiyon, metni karakter karakter yazarak daktilo efekti oluşturur.
func typewriterPrint(s string) {
	typingSpeed := 15 * time.Millisecond // Yazma hızı. Değer küçüldükçe hız artar.
	for _, char := range s {
		fmt.Printf("%c", char)
		time.Sleep(typingSpeed)
	}
	fmt.Println()
}

// centerPrint calculates the padding and prints the text in the middle of the terminal.
// It also uses the typewriter effect.
func centerPrint(s string) {
	padding := (terminalWidth - countCharacters(s)) / 2
	if padding < 0 {
		padding = 0
	}
	fmt.Printf(strings.Repeat(" ", padding))
	typewriterPrint(s)
}

// blinkingCursor simulates a blinking cursor while the program "thinks."
// Bu fonksiyon, yapay zekanın "düşünüyor" hissi vermesi için yanıp sönen bir imleç oluşturur.
func blinkingCursor(duration time.Duration) {
	blinkingSpeed := 500 * time.Millisecond // Yanıp sönme hızı
	endTime := time.Now().Add(duration)
	for time.Now().Before(endTime) {
		fmt.Print("_")
		time.Sleep(blinkingSpeed)
		// Clear the cursor to make it blink.
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
	// ANSI codes add to the length but don't take up space on screen. We remove them.
	s = strings.ReplaceAll(s, ColorCyan, "")
	s = strings.ReplaceAll(s, ColorGreen, "")
	s = strings.ReplaceAll(s, ColorReset, "")
	return len(strings.ReplaceAll(s, " ", ""))
}

// sayJoke prints a random joke from a predefined list.
// It avoids repeating the previous joke.
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
	// "Düşünüyor" efekti ekleniyor.
	fmt.Print(ColorCyan + "..." + ColorReset)
	blinkingCursor(1 * time.Second)
	fmt.Println()
	centerPrint(ColorCyan + jokes[jokeIndex] + ColorReset)
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
	centerPrint(ColorCyan + laughs[getRandomInt(len(laughs))] + ColorReset)
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
	swears := []string{
		"EEE! mına korum böyle oyunun!! yıkıl köpek!",
		"bana bak! seni adam yerine koyduk karşımıza aldık,.. tööbe tööbee",
		"OHA! OHA! kırsaydın klavyeyi!!",
		"doğru oyna orospu!",
		"GÖT!",
	}

	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + swears[0] + ColorReset)
	}
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + swears[1] + ColorReset)
	}
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + swears[2] + ColorReset)
	}
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + swears[3] + ColorReset)
	}
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + swears[4] + ColorReset)
	}
}

// isVowel checks if a given character is a Turkish vowel.
func isVowel(r rune) bool {
	vowels := "aıueöüio"
	return strings.ContainsRune(vowels, r)
}

// stage10 concludes the game with a final joke.
func stage10() {
	centerPrint(ColorCyan + "\nşimdik sana bi fıkra daha:\n" + ColorReset)
	sayJoke()
	userPrompt("Çıkmak için bir tuşa basın.")
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

	centerPrint(ColorCyan + "şimdik sen bi sayı tut, ben bulmaya çalışiim. Ama dürüst ol." + ColorReset)
	centerPrint(ColorCyan + "tahminimde yükselmen gerekirse 'y', düşmem gerekirse 'd' ile yanıt ver." + ColorReset)
	centerPrint(ColorCyan + "sayıyı bulursam 'b' ile yanıt vermen yeterli." + ColorReset)

	for {
		guessCount++
		// "Düşünüyor" efekti ekleniyor.
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

	centerPrint(ColorCyan + fmt.Sprintf(" %d  tahminde bildim...\n", guessCount) + ColorReset)
	if guessCount < score {
		centerPrint(ColorCyan + "kodum! kodum! kodum! hehehehe!" + ColorReset)
	} else if guessCount > score {
		centerPrint(ColorCyan + "lanet olsun! beni geçtin! %100 hile yapmışsındır!" + ColorReset)
	} else {
		centerPrint(ColorCyan + "hmm... eşitiz galiba..." + ColorReset)
	}

	stage10()
}

// stage8 is the number guessing game where the user guesses the computer's number.
func stage8() {
	target := getRandomInt(100) + 1
	var guess int
	guessCount := 0

	centerPrint(ColorCyan + fmt.Sprintf("%s,\n gel senlen oyun oynayak...\nben şimdik 1 ilen 100 arası bi sayı tutiim...\ntuttum.\n", userName) + ColorReset)

	for {
		guessCount++
		userPrompt("tahmin et bakalım..? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var err error
		guess, err = strconv.Atoi(input)
		if err != nil {
			typewriterPrint(ColorCyan + "Geçersiz giriş. Lütfen bir sayı girin." + ColorReset)
			continue
		}

		if guess == target {
			if guessCount <= 3 {
				centerPrint(ColorCyan + fmt.Sprintf(" %d  tahminde nası bildin lan? walla brawo!!\n", guessCount) + ColorReset)
			} else if guessCount <= 5 {
				centerPrint(ColorCyan + fmt.Sprintf(" %d . denemede buldun!! tebrik etmek lazım şindi seni...\n", guessCount) + ColorReset)
			} else if guessCount <= 10 {
				centerPrint(ColorCyan + fmt.Sprintf(" %d tahminde buldun.. eh..\n", guessCount) + ColorReset)
			} else if guessCount <= 20 {
				centerPrint(ColorCyan + fmt.Sprintf("NİHAYET!!!  bişey  %d  kere sorulmaz ki ama, dimi?!\n", guessCount) + ColorReset)
			} else if guessCount <= 30 {
				centerPrint(ColorCyan + fmt.Sprintf("bir an ümidimi kesmiştim! neytse ki  %d  kerede buldun! aferin!\n", guessCount) + ColorReset)
			} else {
				centerPrint(ColorCyan + fmt.Sprintf(" %d \ntahminde bulundun...  sen,\n1- Türkçe bilmiyorsun...\n2- Klavye kullanmasını bilmiyorsun...\n3- ya da cinsel yönden bazısorunların var!!!\nE M B E S İ L !\n", guessCount) + ColorReset)
			}
			score = guessCount
			stage9()
			break
		} else {
			if guess < 1 || guess > 100 {
				centerPrint(ColorCyan + "Abartma! abartma!  1-100 arası dedik!" + ColorReset)
			} else if guess < target {
				if target-guess > 20 {
					centerPrint(ColorCyan + "çık çık" + ColorReset)
				} else {
					centerPrint(ColorCyan + "yaklaştın, acık daa çık!" + ColorReset)
				}
			} else { // guess > target
				if guess-target > 20 {
					centerPrint(ColorCyan + "aşşalara gel aşşalara" + ColorReset)
				} else {
					centerPrint(ColorCyan + "biraz daa düş!" + ColorReset)
				}
			}
		}
	}
}

// stage7 asks for the user's hometown and responds based on the last vowel.
func stage7() {
	userPrompt(fmt.Sprintf("memleket nere %s?\n? ", userName))
	hometown, _ := reader.ReadString('\n')
	hometown = strings.TrimSpace(hometown)

	runes := []rune(hometown)
	lastVowel := ' '
	foundVowel := false

	// Iterate backwards to find the last vowel
	for i := len(runes) - 1; i >= 0; i-- {
		if isVowel(runes[i]) {
			lastVowel = runes[i]
			foundVowel = true
			break
		}
	}

	if foundVowel {
		// "Düşünüyor" efekti ekleniyor.
		fmt.Print(ColorCyan + "..." + ColorReset)
		blinkingCursor(1 * time.Second)
		fmt.Println()
		
		switch lastVowel {
		case 'u', 'o':
			centerPrint(ColorCyan + fmt.Sprintf("madem %slusun,\n buralara ne b*k yemeye geldin?! Ayrıca\n%sdan\n   adam falan çıkmaz!\n", hometown, hometown) + ColorReset)
		case 'ü', 'ö':
			centerPrint(ColorCyan + fmt.Sprintf("heheheh!%sden\n top çıkarmış diyolar!?!", hometown) + ColorReset)
		case 'a', 'ı':
			centerPrint(ColorCyan + fmt.Sprintf("naaaber pis\n%slı!\n", hometown) + ColorReset)
		case 'e', 'i':
			centerPrint(ColorCyan + fmt.Sprintf("nea!? %sden\n     adam çıkmaz ki beah!!!  hihöhöhö!!\n", hometown) + ColorReset)
		}
	}

	laugh()
	centerPrint(ColorCyan + fmt.Sprintf("\nneyse %s,\n kusura bakma...\n", userName) + ColorReset)

	stage8()
}

// stage6 prints a joke and a proverb.
func stage6() {
	centerPrint(ColorCyan + "bak sana şindi konuyla ilgili bir fıkra..." + ColorReset)
	sayJoke()
	laugh()

	proverbs := []string{
		"yani sakla samanı gelir zamanı.",
		"yani arkadaşlarımızı dikkatli seçmemiz lazım.",
		"buradan alınacak ders: Göte giren şemsiye açılmaz..",
	}

	centerPrint(ColorCyan + fmt.Sprintf("\n%s\n", proverbs[getRandomInt(len(proverbs))]) + ColorReset)
	laugh()
	centerPrint("") // Newline
	stage7()
}

// stage5 contains a series of random questions.
func stage5() {
	// Question 1: Eyes
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + fmt.Sprintf("%s!\n", userName) + ColorReset)
		userPrompt("sana gözlerinin çok güzel olduğunu söyleyen olmuşmuydu hiç\n(e/h)? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			centerPrint(ColorCyan + "yalan söylemiş!" + ColorReset)
			laugh()
		} else {
			centerPrint(ColorCyan + "doğrudur. çünkü gözlerin güzel diil!" + ColorReset)
			laugh()
		}
	}

	// Question 2: Money
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + fmt.Sprintf("\nyavrum\n%s\n", userName) + ColorReset)
		userPrompt("ayda 50 milyon kazanmak istermisin?\n(e/h)? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			centerPrint(ColorCyan + "o zaman Ay'a gitmen lazım..." + ColorReset)
			laugh()
		} else {
			centerPrint(ColorCyan + "iyi... zaten Ay'da sağlıklı çalışabileceğini sanmıyordum." + ColorReset)
			laugh()
		}
	}

	// Question 3: Name Origin
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + fmt.Sprintf("\n%s\n", userName) + ColorReset)
		userPrompt("adı nerden geliyo?\n? ")
		reader.ReadString('\n') // Just read and discard the user's input
		centerPrint(ColorCyan + "üüüü! baya uzaktan geliyomuş!" + ColorReset)
		laugh()
	}

	// Question 4: Holding a number
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + fmt.Sprintf("\n%s\n", userName) + ColorReset)
		userPrompt("bi sayı tut.\ntuttunmu (e/h)\n? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			centerPrint(ColorCyan + "şimdi de bırak!" + ColorReset)
			laugh()
		} else {
			centerPrint(ColorCyan + "bi sayıyı tutamadın allah belanı versin" + ColorReset)
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
		centerPrint(ColorCyan + fmt.Sprintf("\nnasılsınız lan\n%s?\n", userName) + ColorReset)
		userPrompt("iyimisin ki (e/h)\n? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			randChoice := getRandomInt(3)
			if randChoice == 0 {
				centerPrint(ColorCyan + "niye iyisin? oturduğun yere bir bak bakiim...\njoysitick falan unutmuş olmasınlar?" + ColorReset)
			} else if randChoice == 1 {
				centerPrint(ColorCyan + fmt.Sprintf("iyi iyi... sen iyi olmaya devam et\n%s!\nuyu da büyü!\n", userName) + ColorReset)
			} else {
				centerPrint(ColorCyan + fmt.Sprintf("böyle bir hayatta nasıl iyi oluyorsunuz ki lan\n%s?\nbize de söyle yolunu biz de iyi olalım..\n", userName) + ColorReset)
			}
		} else {
			randChoice := getRandomInt(3)
			if randChoice == 0 {
				centerPrint(ColorCyan + "bana ne lan! geber!" + ColorReset)
			} else if randChoice == 1 {
				centerPrint(ColorCyan + "iyi iyi allah kötülük versin! he he he !!" + ColorReset)
			} else {
				centerPrint(ColorCyan + fmt.Sprintf("derdini anlat bana! açıl bana yavrucuum! utanma ben doktorum...\nKötü olmana sebep olan şey nedir %s", userName) + ColorReset)
				reader.ReadString('\n') // Read and discard
				centerPrint(ColorCyan + "\n??\nhahahahahahahaha!!! git allasen yaw! dert  ettiğin şeye bak!" + ColorReset)
			}
		}
		laugh()
	}

	// Question 7: Student
	if getRandomInt(2) == 1 {
		centerPrint(ColorCyan + fmt.Sprintf("\nneyse... %s\n", userName) + ColorReset)
		userPrompt("      öğrencimisin?\n? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			randChoice := getRandomInt(2)
			if randChoice == 0 {
				centerPrint(ColorCyan + "wah! wah! wah! çok üzüldüm.. ailenin haberi varmı? ha!haha!!hohoho!!!\n" + ColorReset)
			} else {
				centerPrint(ColorCyan + "nerde öğrencisin? okulda mı?? hihohohohhohohooo!!!\nespri konuşlandırdım!!\n" + ColorReset)
			}
		} else {
			randChoice := getRandomInt(2)
			if randChoice == 0 {
				centerPrint(ColorCyan + "ulan insan en azından askerden yırtmak için öğrenci olur! Ama sen, tıss!\n" + ColorReset)
			} else {
				userPrompt("hangi işle meşgulsun o vakit?\n? ")
				reader.ReadString('\n') // Read and discard
				centerPrint(ColorCyan + "siktir lan göt! cümle alem senin ne mal olduğunu biliyor.\n" + ColorReset)
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
		userPrompt("oldu olcak kilonu da söyle bari... çok umurumda ya...\n? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		weight, err = strconv.Atoi(input)
		if err != nil {
			typewriterPrint(ColorCyan + "Geçersiz giriş. Lütfen bir sayı girin." + ColorReset)
			continue
		}

		// "Düşünüyor" efekti ekleniyor.
		fmt.Print(ColorCyan + "..." + ColorReset)
		blinkingCursor(1 * time.Second)
		fmt.Println()
		
		if weight <= 39 {
			centerPrint(ColorCyan + "Rüzgarlı havada dışarı falan çıkma hehehe!" + ColorReset)
			actDumb()
		} else if weight >= 40 && weight <= 59 {
			centerPrint(ColorCyan + "o kadar yemiş yersen ishal de olursun, kabız da!" + ColorReset)
			actDumb()
		} else if weight >= 60 && weight <= 79 {
			centerPrint(ColorCyan + "sen normalsin o yüzden dalga geçmiicem... noormaal! noormaal! hehehe!!" + ColorReset)
			actDumb()
		} else if weight >= 80 && weight <= 99 {
			randChoice := getRandomInt(3)
			if randChoice == 0 {
				centerPrint(ColorCyan + "Lütfen, oturduğun koltuk sağlam kalsın!" + ColorReset)
			} else if randChoice == 1 {
				centerPrint(ColorCyan + "Maaşşallaaah! damızlıkmısın? hangi çiftlikte yetiştin? keh!keh!keh!!." + ColorReset)
			} else {
				centerPrint(ColorCyan + "Duba! dikkat et benim üstüme düşme!" + ColorReset)
			}
			actDumb()
		} else if weight >= 100 {
			centerPrint(ColorCyan + "Anlamıştım... 2 saattir klavyenin anasını ağlattın" + ColorReset)
			actDumb()
		}
		centerPrint("") // Newline
		break
	}
	stage5()
}

// stage3 asks for the user's height and responds accordingly.
func stage3() {
	var height int
	for {
		userPrompt("boyun kaç cm senin?\n? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		height, err = strconv.Atoi(input)
		if err != nil {
			typewriterPrint(ColorCyan + "Geçersiz giriş. Lütfen bir sayı girin." + ColorReset)
			continue
		}

		// "Düşünüyor" efekti ekleniyor.
		fmt.Print(ColorCyan + "..." + ColorReset)
		blinkingCursor(1 * time.Second)
		fmt.Println()
		
		if height <= 99 {
			centerPrint(ColorCyan + "Deden pigmelerin hangi kavminden lan?" + ColorReset)
		} else if height >= 100 && height <= 149 {
			centerPrint(ColorCyan + "Kısa boylu olman önemli diil, diyeceğimi sanıyorsan yanılıyorsun pis cüce!" + ColorReset)
		} else if height >= 150 && height <= 169 {
			centerPrint(ColorCyan + "Bacaklarına biraz gübre ektir. Faydası olur. kah!kih!koh!" + ColorReset)
		} else if height >= 170 && height <= 189 {
			centerPrint(ColorCyan + "iyi... bana ne... sorduk mu?!" + ColorReset)
		} else if height >= 190 && height <= 209 {
			centerPrint(ColorCyan + "Oha! fasülye sırığı!" + ColorReset)
		} else if height >= 210 {
			centerPrint(ColorCyan + "Yok deve!! kaç santim dedik, milim demedik!" + ColorReset)
			continue
		}
		centerPrint("") // Newline
		break
	}
	stage4()
}

// stage2 asks for the user's age and responds accordingly.
func stage2() {
	var age int
	for {
		userPrompt("kaç yaşındasın?\n? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		age, err = strconv.Atoi(input)
		if err != nil {
			typewriterPrint(ColorCyan + "Geçersiz giriş. Lütfen bir sayı girin." + ColorReset)
			continue
		}
		
		// "Düşünüyor" efekti ekleniyor.
		fmt.Print(ColorCyan + "..." + ColorReset)
		blinkingCursor(1 * time.Second)
		fmt.Println()

		if age <= 4 {
			centerPrint(ColorCyan + "çok küçükmüşsün be! sen git anan gelsin lan lavuk!" + ColorReset)
		} else if age >= 5 && age <= 9 {
			userPrompt("sütünü içtin mi yavrum?\n(e/h)? ")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(strings.ToLower(choice))
			if choice == "e" {
				centerPrint(ColorCyan + "Beynine pek etkisi olmamış, git biraz da PEPSı iç!" + ColorReset)
			} else {
				centerPrint(ColorCyan + "bok iç o zaman!" + ColorReset)
			}
		} else if age >= 10 && age <= 17 {
			centerPrint(ColorCyan + "iyi iyi 18ine pek bişi kalmamış... Uyu da büyü!" + ColorReset)
		} else if age >= 18 && age <= 24 {
			userPrompt("Oy kullancanmı genç?\n(e/h)? ")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(strings.ToLower(choice))
			if choice == "e" {
				centerPrint(ColorCyan + "ver de gör ebeninkini!" + ColorReset)
			} else {
				centerPrint(ColorCyan + "Ulan sen ne biçim Tee.Cee vatandaşısın? Hayvan!..." + ColorReset)
			}
		} else if age >= 25 && age <= 39 {
			centerPrint(ColorCyan + "vayy! naber morruk? Nerde eski programcılar dimi mirim?" + ColorReset)
		} else if age >= 40 && age <= 59 {
			centerPrint(ColorCyan + "Yuh! bayağı yaşlısın... yaşlılar muhattabım diildir.. Git estetik yaptır gel..." + ColorReset)
		} else if age >= 60 && age <= 98 {
			centerPrint(ColorCyan + "Ulan bunak! Klavyeyi nası görüyon? Geber de helvanı yiyelim. hehehe!" + ColorReset)
		} else if age >= 99 {
			centerPrint(ColorCyan + "Kafa bulma lan göt" + ColorReset)
			continue
		}
		centerPrint("") // Newline
		break
	}
	stage3()
}

// stage1 asks for the user's name and starts the conversation.
func stage1() {
	userPrompt("senin adın ne güzelim\n? ")
	input, _ := reader.ReadString('\n')
	userName = strings.TrimSpace(input)

	// "Düşünüyor" efekti ekleniyor.
	fmt.Print(ColorCyan + "..." + ColorReset)
	blinkingCursor(1 * time.Second)
	fmt.Println()

	charCount := countCharacters(userName)

	if charCount <= 2 {
		centerPrint(ColorCyan + fmt.Sprintf("Uzak doğudan mısın yoksa başka bir gezegenden mi?\n %d\n harfli ismini biraz zor telafuz ediyorum da...\n%c...\n%ch%s!!!\neee.. olmadı galiba... hehehehehee!\n", charCount, userName[0], userName[0], userName) + ColorReset)
	} else if charCount >= 8 {
		centerPrint(ColorCyan + "maaşşallaaaah!\nnüfus memuru ananı babanı pek sevmiyormuş galiba!!!" + ColorReset)
		laugh()
	}

	centerPrint(ColorCyan + fmt.Sprintf("%s...\n", userName) + ColorReset)
	stage2()
}

// stage0 is the initial welcome message.
func stage0() {
	centerPrint(ColorCyan + "merhaba" + ColorReset)
	centerPrint(ColorCyan + "ben karabasan..." + ColorReset)
	stage1()
}

// main is the entry point of the Go application.
func main() {
	rand.Seed(time.Now().UnixNano())
	stage0()
}
