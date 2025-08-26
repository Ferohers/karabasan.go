package main

import (
	"bufio"
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
)

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
    aiResponse("\nşimdik sana bi fıkra daha:\n")
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
    aiResponse("şimdik sen bi sayı tut, ben bulmaya çalışiim. Ama dürüst ol.")
    aiResponse("tahminimde yükselmen gerekirse 'y', düşmem gerekirse 'd' ile yanıt ver.")
    aiResponse("sayıyı bulursam 'b' ile yanıt vermen yeterli.")
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
    aiResponse(fmt.Sprintf(" %d  tahminde bildim...\n", guessCount))
    if guessCount < score {
        aiResponse("kodum! kodum! kodum! hehehehe!")
    } else if guessCount > score {
        aiResponse("lanet olsun! beni geçtin! %100 hile yapmışsındır!")
    } else {
        aiResponse("hmm... eşitiz galiba...")
    }
    stage10()
}

// stage8 is the number guessing game where the user guesses the computer's number.
func stage8() {
    target := getRandomInt(100) + 1
    var guess int
    guessCount := 0
    aiResponse(fmt.Sprintf("%s,\n gel senlen oyun oynayak...\nben şimdik 1 ilen 100 arası bi sayı tutiim...\ntuttum.\n", userName))
    for {
        guessCount++
        userPrompt("tahmin et bakalım..? ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        var err error
        guess, err = strconv.Atoi(input)
        if err != nil {
            aiResponse("Geçersiz giriş. Lütfen bir sayı girin.")
            continue
        }
        if guess == target {
            if guessCount <= 3 {
                aiResponse(fmt.Sprintf(" %d  tahminde nası bildin lan? walla brawo!!\n", guessCount))
            } else if guessCount <= 5 {
                aiResponse(fmt.Sprintf(" %d . denemede buldun!! tebrik etmek lazım şindi seni...\n", guessCount))
            } else if guessCount <= 10 {
                aiResponse(fmt.Sprintf(" %d tahminde buldun.. eh..\n", guessCount))
            } else if guessCount <= 20 {
                aiResponse(fmt.Sprintf("NİHAYET!!!  bişey  %d  kere sorulmaz ki ama, dimi?!\n", guessCount))
            } else if guessCount <= 30 {
                aiResponse(fmt.Sprintf("bir an ümidimi kesmiştim! neytse ki  %d  kerede buldun! aferin!\n", guessCount))
            } else {
                aiResponse(fmt.Sprintf(" %d \ntahminde bulundun...  sen,\n1- Türkçe bilmiyorsun...\n2- Klavye kullanmasını bilmiyorsun...\n3- ya da cinsel yönden bazısorunların var!!!\nE M B E S İ L !\n", guessCount))
            }
            score = guessCount
            stage9()
            break
        } else {
            if guess < 1 || guess > 100 {
                aiResponse("Abartma! abartma!  1-100 arası dedik!")
            } else if guess < target {
                if target-guess > 20 {
                    aiResponse("çık çık")
                } else {
                    aiResponse("yaklaştın, acık daa çık!")
                }
            } else { // guess > target
                if guess-target > 20 {
                    aiResponse("aşşalara gel aşşalara")
                } else {
                    aiResponse("biraz daa düş!")
                }
            }
        }
    }
}

// stage7 asks for the user's hometown and responds based on the last vowel.
func stage7() {
    userPrompt(fmt.Sprintf("memleket nere %s?", userName))
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
            aiResponse(fmt.Sprintf("madem %slusun,\n buralara ne b*k yemeye geldin?! Ayrıca\n%sdan\n   adam falan çıkmaz!\n", hometown, hometown))
        case 'ü', 'ö':
            aiResponse(fmt.Sprintf("heheheh!%sden\n top çıkarmış diyolar!?!", hometown))
        case 'a', 'ı':
            aiResponse(fmt.Sprintf("naaaber pis\n%slı!\n", hometown))
        case 'e', 'i':
            aiResponse(fmt.Sprintf("nea!? %sden\n     adam çıkmaz ki beah!!!  hihöhöhö!!\n", hometown))
        }
    }
    laugh()
    aiResponse(fmt.Sprintf("\nneyse %s,\n kusura bakma...\n", userName))
    stage8()
}

// stage6 prints a joke and a proverb.
func stage6() {
    aiResponse("bak sana şindi konuyla ilgili bir fıkra...")
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
        userPrompt(fmt.Sprintf("%s!\nsana gözlerinin çok güzel olduğunu söyleyen olmuşmuydu hiç\n(e/h)? ", userName))
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(strings.ToLower(input))
        if input == "e" {
            aiResponse("yalan söylemiş!")
            laugh()
        } else {
            aiResponse("doğrudur. çünkü gözlerin güzel diil!")
            laugh()
        }
    }
    // Question 2: Money
    if getRandomInt(2) == 1 {
        userPrompt(fmt.Sprintf("\nyavrum\n%s\nayda 50 milyon kazanmak istermisin?\n(e/h)? ", userName))
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(strings.ToLower(input))
        if input == "e" {
            aiResponse("o zaman Ay'a gitmen lazım...")
            laugh()
        } else {
            aiResponse("iyi... zaten Ay'da sağlıklı çalışabileceğini sanmıyordum.")
            laugh()
        }
    }
    // Question 3: Name Origin
    if getRandomInt(2) == 1 {
        userPrompt(fmt.Sprintf("\n%s\nadı nerden geliyo? ", userName))
        reader.ReadString('\n') 
        aiResponse("üüüü! baya uzaktan geliyomuş!")
        laugh()
    }
    // Question 4: Holding a number
    if getRandomInt(2) == 1 {
        userPrompt(fmt.Sprintf("\n%s\nbi sayı tut.\ntuttunmu (e/h)?", userName))
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(strings.ToLower(input))
        if input == "e" {
            aiResponse("şimdi de bırak!")
            laugh()
        } else {
            aiResponse("bi sayıyı tutamadın allah belanı versin")
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
            userPrompt(fmt.Sprintf("\n%s, sana kısaca %s diyebilirmiyim??\n(e/h)? ", userName, nickname))
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
        userPrompt(fmt.Sprintf("\nnasılsınız lan\n%s?\niyimisin ki (e/h)? ", userName))
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(strings.ToLower(input))
        if input == "e" {
            randChoice := getRandomInt(3)
            if randChoice == 0 {
                aiResponse("niye iyisin? oturduğun yere bir bak bakiim...\njoysitick falan unutmuş olmasınlar?")
            } else if randChoice == 1 {
                aiResponse(fmt.Sprintf("iyi iyi... sen iyi olmaya devam et\n%s!\nuyu da büyü!\n", userName))
            } else {
                aiResponse(fmt.Sprintf("böyle bir hayatta nasıl iyi oluyorsunuz ki lan\n%s?\nbize de söyle yolunu biz de iyi olalım..\n", userName))
            }
        } else {
            randChoice := getRandomInt(3)
            if randChoice == 0 {
                aiResponse("bana ne lan! geber!")
            } else if randChoice == 1 {
                aiResponse("iyi iyi allah kötülük versin! he he he !!")
            } else {
                aiResponse(fmt.Sprintf("derdini anlat bana! açıl bana yavrucuum! utanma ben doktorum...\nKötü olmana sebep olan şey nedir %s", userName))
                reader.ReadString('\n') 
                aiResponse("\n??\nhahahahahahahaha!!! git allasen yaw! dert  ettiğin şeye bak!")
            }
        }
        laugh()
    }
    // Question 7: Student
    if getRandomInt(2) == 1 {
        userPrompt(fmt.Sprintf("\nneyse... %s\n      öğrencimisin? ", userName))
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(strings.ToLower(input))
        if input == "e" {
            randChoice := getRandomInt(2)
            if randChoice == 0 {
                aiResponse("wah! wah! wah! çok üzüldüm.. ailenin haberi varmı? ha!haha!!hohoho!!!\n")
            } else {
                aiResponse("nerde öğrencisin? okulda mı?? hihohohohhohohooo!!!\nespri konuşlandırdım!!\n")
            }
        } else {
            randChoice := getRandomInt(2)
            if randChoice == 0 {
                aiResponse("ulan insan en azından askerden yırtmak için öğrenci olur! Ama sen, tıss!\n")
            } else {
                userPrompt("hangi işle meşgulsun o vakit? ")
                reader.ReadString('\n') 
                aiResponse("siktir lan göt! cümle alem senin ne mal olduğunu biliyor.\n")
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
        userPrompt("oldu olcak kilonu da söyle bari... çok umurumda ya?")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        var err error
        weight, err = strconv.Atoi(input)
        if err != nil {
            aiResponse("Geçersiz giriş. Lütfen bir sayı girin.")
            continue
        }
        aiResponse(fmt.Sprintf("%d kilon var demek? Bakalım...", weight))
        if weight <= 39 {
            aiResponse("Rüzgarlı havada dışarı falan çıkma hehehe!")
            actDumb()
        } else if weight >= 40 && weight <= 59 {
            aiResponse("o kadar yemiş yersen ishal de olursun, kabız da!")
            actDumb()
        } else if weight >= 60 && weight <= 79 {
            aiResponse("sen normalsin o yüzden dalga geçmiicem... noormaal! noormaal! hehehe!!")
            actDumb()
        } else if weight >= 80 && weight <= 99 {
            randChoice := getRandomInt(3)
            if randChoice == 0 {
                aiResponse("Lütfen, oturduğun koltuk sağlam kalsın!")
            } else if randChoice == 1 {
                aiResponse("Maaşşallaaah! damızlıkmısın? hangi çiftlikte yetiştin? keh!keh!keh!!.")
            } else {
                aiResponse("Duba! dikkat et benim üstüme düşme!")
            }
            actDumb()
        } else if weight >= 100 {
            aiResponse("Anlamıştım... 2 saattir klavyenin anasını ağlattın")
            actDumb()
        }
        centerPrint("")
        break
    }
    stage5()
}

// stage3 asks for the user's height and responds accordingly.
func stage3() {
    userPrompt("boyun kaç cm senin?")
    var height int
    for {
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        var err error
        height, err = strconv.Atoi(input)
        if err != nil {
            aiResponse("Geçersiz giriş. Lütfen bir sayı girin.")
            continue
        }
        aiResponse(fmt.Sprintf("%d cm boyun var demek? Hmm...", height))
        if height <= 99 {
            aiResponse("Deden pigmelerin hangi kavminden lan?")
        } else if height >= 100 && height <= 149 {
            aiResponse("Kısa boylu olman önemli diil, diyeceğimi sanıyorsan yanılıyorsun pis cüce!")
        } else if height >= 150 && height <= 169 {
            aiResponse("Bacaklarına biraz gübre ektir. Faydası olur. kah!kih!koh!")
        } else if height >= 170 && height <= 189 {
            aiResponse("iyi... bana ne... sorduk mu?!")
        } else if height >= 190 && height <= 209 {
            aiResponse("Oha! fasülye sırığı!")
        } else if height >= 210 {
            aiResponse("Yok deve!! kaç santim dedik, milim demedik!")
            continue
        }
        centerPrint("")
        break
    }
    stage4()
}

// stage2 asks for the user's age and responds accordingly.
func stage2() {
    userPrompt("kaç yaşındasın?")
    var age int
    for {
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        var err error
        age, err = strconv.Atoi(input)
        if err != nil {
            aiResponse("Geçersiz giriş. Lütfen bir sayı girin.")
            continue
        }
        aiResponse(fmt.Sprintf("Öyle mi, %d yaşındasın demek?", age))
        if age <= 4 {
            aiResponse("çok küçükmüşsün be! sen git anan gelsin lan lavuk!")
        } else if age >= 5 && age <= 9 {
            userPrompt("sütünü içtin mi yavrum?\n(e/h)? ")
            choice, _ := reader.ReadString('\n')
            choice = strings.TrimSpace(strings.ToLower(choice))
            if choice == "e" {
                aiResponse("Beynine pek etkisi olmamış, git biraz da PEPSı iç!")
            } else {
                aiResponse("bok iç o zaman!")
            }
        } else if age >= 10 && age <= 17 {
            aiResponse("iyi iyi 18ine pek bişi kalmamış... Uyu da büyü!")
        } else if age >= 18 && age <= 24 {
            userPrompt("Oy kullancanmı genç?\n(e/h)? ")
            choice, _ := reader.ReadString('\n')
            choice = strings.TrimSpace(strings.ToLower(choice))
            if choice == "e" {
                aiResponse("ver de gör ebeninkini!")
            } else {
                aiResponse("Ulan sen ne biçim Tee.Cee vatandaşısın? Hayvan!...")
            }
        } else if age >= 25 && age <= 39 {
            aiResponse("vayy! naber morruk? Nerde eski programcılar dimi mirim?")
        } else if age >= 40 && age <= 59 {
            aiResponse("Yuh! bayağı yaşlısın... yaşlılar muhattabım diildir.. Git estetik yaptır gel...")
        } else if age >= 60 && age <= 98 {
            aiResponse("Ulan bunak! Klavyeyi nası görüyon? Geber de helvanı yiyelim. hehehe!")
        } else if age >= 99 {
            aiResponse("Kafa bulma lan göt")
            continue
        }
        centerPrint("")
        break
    }
    stage3()
}

// stage1 asks for the user's name and starts the conversation.
func stage1() {
    userPrompt("senin adın ne güzelim?")
    input, _ := reader.ReadString('\n')
    userName = strings.TrimSpace(input)
    aiResponse(fmt.Sprintf("Tanıştığıma memnun oldum, %s. Hadi başlayalım.", userName))
    stage2()
}

// stage0 is the initial welcome and introduction.
func stage0() {
	fmt.Println()
	centerPrint(ColorCyan + "Merhaba, hoş geldin." + ColorReset)
	time.Sleep(1 * time.Second)
	centerPrint(ColorCyan + "Ben yeni nesil bir terminal arayüzüyüm." + ColorReset)
	time.Sleep(1 * time.Second)
	aiResponse("Tanıştığıma memnun oldum. Hadi başlayalım.")
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
	stage0()
}

