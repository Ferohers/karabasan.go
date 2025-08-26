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
)

// The global reader to handle all terminal input.
var reader = bufio.NewReader(os.Stdin)

// getRandomInt returns a random integer up to the given maximum (exclusive).
func getRandomInt(max int) int {
	return rand.Intn(max)
}

// countCharacters counts the number of non-space characters in a string.
func countCharacters(s string) int {
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
	fmt.Println(ColorCyan + jokes[jokeIndex] + ColorReset)
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
	fmt.Println(ColorCyan + laughs[getRandomInt(len(laughs))] + ColorReset)
}

// actDumb has a 50% chance of printing a "dumb" joke.
func actDumb() {
	if getRandomInt(2) == 1 {
		fmt.Println(ColorCyan + "\ngeri zekalı taklidi yap bakiim...\nTamam tamam bukadar yeter!!!\n" + ColorReset)
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
		fmt.Println(ColorCyan + swears[0] + ColorReset)
	}
	if getRandomInt(2) == 1 {
		fmt.Println(ColorCyan + swears[1] + ColorReset)
	}
	if getRandomInt(2) == 1 {
		fmt.Println(ColorCyan + swears[2] + ColorReset)
	}
	if getRandomInt(2) == 1 {
		fmt.Println(ColorCyan + swears[3] + ColorReset)
	}
	if getRandomInt(2) == 1 {
		fmt.Println(ColorCyan + swears[4] + ColorReset)
	}
}

// isVowel checks if a given character is a Turkish vowel.
func isVowel(r rune) bool {
	vowels := "aıueöüio"
	return strings.ContainsRune(vowels, r)
}

// stage10 concludes the game with a final joke.
func stage10() {
	fmt.Println(ColorCyan + "\nşimdik sana bi fıkra daha:\n" + ColorReset)
	sayJoke()
	fmt.Println(ColorGreen + "Çıkmak için bir tuşa basın." + ColorReset)
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

	fmt.Println(ColorCyan + "şimdik sen bi sayı tut, ben bulmaya çalışiim. Ama dürüst ol." + ColorReset)
	fmt.Println(ColorCyan + "tahminimde yükselmen gerekirse 'y', düşmem gerekirse 'd' ile yanıt ver." + ColorReset)
	fmt.Println(ColorCyan + "sayıyı bulursam 'b' ile yanıt vermen yeterli." + ColorReset)

	for {
		guessCount++
		fmt.Printf(ColorCyan + " %d  ??\n" + ColorReset, guess)
		fmt.Print(ColorGreen + "? " + ColorReset)

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

	fmt.Printf(ColorCyan + " %d  tahminde bildim...\n" + ColorReset, guessCount)
	if guessCount < score {
		fmt.Println(ColorCyan + "kodum! kodum! kodum! hehehehe!" + ColorReset)
	} else if guessCount > score {
		fmt.Println(ColorCyan + "lanet olsun! beni geçtin! %100 hile yapmışsındır!" + ColorReset)
	} else {
		fmt.Println(ColorCyan + "hmm... eşitiz galiba..." + ColorReset)
	}

	stage10()
}

// stage8 is the number guessing game where the user guesses the computer's number.
func stage8() {
	target := getRandomInt(100) + 1
	var guess int
	guessCount := 0

	fmt.Printf(ColorCyan + "%s,\n gel senlen oyun oynayak...\nben şimdik 1 ilen 100 arası bi sayı tutiim...\ntuttum.\n" + ColorReset, userName)

	for {
		guessCount++
		fmt.Print(ColorGreen + "tahmin et bakalım..? " + ColorReset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var err error
		guess, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println(ColorCyan + "Geçersiz giriş. Lütfen bir sayı girin." + ColorReset)
			continue
		}

		if guess == target {
			if guessCount <= 3 {
				fmt.Printf(ColorCyan + " %d  tahminde nası bildin lan? walla brawo!!\n" + ColorReset, guessCount)
			} else if guessCount <= 5 {
				fmt.Printf(ColorCyan + " %d . denemede buldun!! tebrik etmek lazım şindi seni...\n" + ColorReset, guessCount)
			} else if guessCount <= 10 {
				fmt.Printf(ColorCyan + " %d tahminde buldun.. eh..\n" + ColorReset, guessCount)
			} else if guessCount <= 20 {
				fmt.Printf(ColorCyan + "NİHAYET!!!  bişey  %d  kere sorulmaz ki ama, dimi?!\n" + ColorReset, guessCount)
			} else if guessCount <= 30 {
				fmt.Printf(ColorCyan + "bir an ümidimi kesmiştim! neytse ki  %d  kerede buldun! aferin!\n" + ColorReset, guessCount)
			} else {
				fmt.Printf(ColorCyan + " %d \ntahminde bulundun...  sen,\n1- Türkçe bilmiyorsun...\n2- Klavye kullanmasını bilmiyorsun...\n3- ya da cinsel yönden bazısorunların var!!!\nE M B E S İ L !\n" + ColorReset, guessCount)
			}
			score = guessCount
			stage9()
			break
		} else {
			if guess < 1 || guess > 100 {
				fmt.Println(ColorCyan + "Abartma! abartma!  1-100 arası dedik!" + ColorReset)
			} else if guess < target {
				if target-guess > 20 {
					fmt.Println(ColorCyan + "çık çık" + ColorReset)
				} else {
					fmt.Println(ColorCyan + "yaklaştın, acık daa çık!" + ColorReset)
				}
			} else { // guess > target
				if guess-target > 20 {
					fmt.Println(ColorCyan + "aşşalara gel aşşalara" + ColorReset)
				} else {
					fmt.Println(ColorCyan + "biraz daa düş!" + ColorReset)
				}
			}
		}
	}
}

// stage7 asks for the user's hometown and responds based on the last vowel.
func stage7() {
	fmt.Printf(ColorGreen + "memleket nere %s?\n? " + ColorReset, userName)
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
		switch lastVowel {
		case 'u', 'o':
			fmt.Printf(ColorCyan + "madem %slusun,\n buralara ne b*k yemeye geldin?! Ayrıca\n%sdan\n   adam falan çıkmaz!\n" + ColorReset, hometown, hometown)
		case 'ü', 'ö':
			fmt.Printf(ColorCyan + "heheheh!%sden\n top çıkarmış diyolar!?!" + ColorReset, hometown)
		case 'a', 'ı':
			fmt.Printf(ColorCyan + "naaaber pis\n%slı!\n" + ColorReset, hometown)
		case 'e', 'i':
			fmt.Printf(ColorCyan + "nea!? %sden\n     adam çıkmaz ki beah!!!  hihöhöhö!!\n" + ColorReset, hometown)
		}
	}

	laugh()
	fmt.Printf(ColorCyan + "\nneyse %s,\n kusura bakma...\n" + ColorReset, userName)

	stage8()
}

// stage6 prints a joke and a proverb.
func stage6() {
	fmt.Println(ColorCyan + "bak sana şindi konuyla ilgili bir fıkra..." + ColorReset)
	sayJoke()
	laugh()

	proverbs := []string{
		"yani sakla samanı gelir zamanı.",
		"yani arkadaşlarımızı dikkatli seçmemiz lazım.",
		"buradan alınacak ders: Göte giren şemsiye açılmaz..",
	}

	fmt.Printf(ColorCyan + "\n%s\n" + ColorReset, proverbs[getRandomInt(len(proverbs))])
	laugh()
	fmt.Println()
	stage7()
}

// stage5 contains a series of random questions.
func stage5() {
	// Question 1: Eyes
	if getRandomInt(2) == 1 {
		fmt.Printf(ColorCyan + "%s!\n" + ColorReset, userName)
		fmt.Print(ColorGreen + "sana gözlerinin çok güzel olduğunu söyleyen olmuşmuydu hiç\n(e/h)? " + ColorReset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			fmt.Println(ColorCyan + "yalan söylemiş!" + ColorReset)
			laugh()
		} else {
			fmt.Println(ColorCyan + "doğrudur. çünkü gözlerin güzel diil!" + ColorReset)
			laugh()
		}
	}

	// Question 2: Money
	if getRandomInt(2) == 1 {
		fmt.Printf(ColorCyan + "\nyavrum\n%s\n" + ColorReset, userName)
		fmt.Print(ColorGreen + "ayda 50 milyon kazanmak istermisin?\n(e/h)? " + ColorReset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			fmt.Println(ColorCyan + "o zaman Ay'a gitmen lazım..." + ColorReset)
			laugh()
		} else {
			fmt.Println(ColorCyan + "iyi... zaten Ay'da sağlıklı çalışabileceğini sanmıyordum." + ColorReset)
			laugh()
		}
	}

	// Question 3: Name Origin
	if getRandomInt(2) == 1 {
		fmt.Printf(ColorCyan + "\n%s\n" + ColorReset, userName)
		fmt.Print(ColorGreen + "adı nerden geliyo?\n? " + ColorReset)
		reader.ReadString('\n') // Just read and discard the user's input
		fmt.Println(ColorCyan + "üüüü! baya uzaktan geliyomuş!" + ColorReset)
		laugh()
	}

	// Question 4: Holding a number
	if getRandomInt(2) == 1 {
		fmt.Printf(ColorCyan + "\n%s\n" + ColorReset, userName)
		fmt.Print(ColorGreen + "bi sayı tut.\ntuttunmu (e/h)\n? " + ColorReset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			fmt.Println(ColorCyan + "şimdi de bırak!" + ColorReset)
			laugh()
		} else {
			fmt.Println(ColorCyan + "bi sayıyı tutamadın allah belanı versin" + ColorReset)
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
			fmt.Printf(ColorCyan + "\n%s, sana kısaca %s diyebilirmiyim??\n" + ColorReset, userName, nickname)
			fmt.Print(ColorGreen + "(e/h)? " + ColorReset)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(strings.ToLower(input))
			if input == "e" {
				fmt.Println(ColorCyan + "iyi... ama ben demek istemiyorum!" + ColorReset)
				laugh()
			} else {
				fmt.Printf(ColorCyan + "%s! %s! %s!\n" + ColorReset, nickname, nickname, nickname)
				laugh()
			}
		}
	}

	// Question 6: How are you?
	if getRandomInt(2) == 1 {
		fmt.Printf(ColorCyan + "\nnasılsınız lan\n%s?\n" + ColorReset, userName)
		fmt.Print(ColorGreen + "iyimisin ki (e/h)\n? " + ColorReset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			randChoice := getRandomInt(3)
			if randChoice == 0 {
				fmt.Println(ColorCyan + "niye iyisin? oturduğun yere bir bak bakiim...\njoysitick falan unutmuş olmasınlar?" + ColorReset)
			} else if randChoice == 1 {
				fmt.Printf(ColorCyan + "iyi iyi... sen iyi olmaya devam et\n%s!\nuyu da büyü!\n" + ColorReset, userName)
			} else {
				fmt.Printf(ColorCyan + "böyle bir hayatta nasıl iyi oluyorsunuz ki lan\n%s?\nbize de söyle yolunu biz de iyi olalım..\n" + ColorReset, userName)
			}
		} else {
			randChoice := getRandomInt(3)
			if randChoice == 0 {
				fmt.Println(ColorCyan + "bana ne lan! geber!" + ColorReset)
			} else if randChoice == 1 {
				fmt.Println(ColorCyan + "iyi iyi allah kötülük versin! he he he !!" + ColorReset)
			} else {
				fmt.Printf(ColorCyan + "derdini anlat bana! açıl bana yavrucuum! utanma ben doktorum...\nKötü olmana sebep olan şey nedir %s" + ColorReset, userName)
				reader.ReadString('\n') // Read and discard
				fmt.Println(ColorCyan + "\n??\nhahahahahahahaha!!! git allasen yaw! dert  ettiğin şeye bak!" + ColorReset)
			}
		}
		laugh()
	}

	// Question 7: Student
	if getRandomInt(2) == 1 {
		fmt.Printf(ColorCyan + "\nneyse... %s\n" + ColorReset, userName)
		fmt.Print(ColorGreen + "      öğrencimisin?\n? " + ColorReset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			randChoice := getRandomInt(2)
			if randChoice == 0 {
				fmt.Println(ColorCyan + "wah! wah! wah! çok üzüldüm.. ailenin haberi varmı? ha!haha!!hohoho!!!\n" + ColorReset)
			} else {
				fmt.Println(ColorCyan + "nerde öğrencisin? okulda mı?? hihohohohhohohooo!!!\nespri konuşlandırdım!!\n" + ColorReset)
			}
		} else {
			randChoice := getRandomInt(2)
			if randChoice == 0 {
				fmt.Println(ColorCyan + "ulan insan en azından askerden yırtmak için öğrenci olur! Ama sen, tıss!\n" + ColorReset)
			} else {
				fmt.Print(ColorGreen + "hangi işle meşgulsun o vakit?\n? " + ColorReset)
				reader.ReadString('\n') // Read and discard
				fmt.Println(ColorCyan + "siktir lan göt! cümle alem senin ne mal olduğunu biliyor.\n" + ColorReset)
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
		fmt.Print(ColorGreen + "oldu olcak kilonu da söyle bari... çok umurumda ya...\n? " + ColorReset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		weight, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println(ColorCyan + "Geçersiz giriş. Lütfen bir sayı girin." + ColorReset)
			continue
		}

		if weight <= 39 {
			fmt.Println(ColorCyan + "Rüzgarlı havada dışarı falan çıkma hehehe!" + ColorReset)
			actDumb()
		} else if weight >= 40 && weight <= 59 {
			fmt.Println(ColorCyan + "o kadar yemiş yersen ishal de olursun, kabız da!" + ColorReset)
			actDumb()
		} else if weight >= 60 && weight <= 79 {
			fmt.Println(ColorCyan + "sen normalsin o yüzden dalga geçmiicem... noormaal! noormaal! hehehe!!" + ColorReset)
			actDumb()
		} else if weight >= 80 && weight <= 99 {
			randChoice := getRandomInt(3)
			if randChoice == 0 {
				fmt.Println(ColorCyan + "Lütfen, oturduğun koltuk sağlam kalsın!" + ColorReset)
			} else if randChoice == 1 {
				fmt.Println(ColorCyan + "Maaşşallaaah! damızlıkmısın? hangi çiftlikte yetiştin? keh!keh!keh!!." + ColorReset)
			} else {
				fmt.Println(ColorCyan + "Duba! dikkat et benim üstüme düşme!" + ColorReset)
			}
			actDumb()
		} else if weight >= 100 {
			fmt.Println(ColorCyan + "Anlamıştım... 2 saattir klavyenin anasını ağlattın" + ColorReset)
			actDumb()
		}
		fmt.Println()
		break
	}
	stage5()
}

// stage3 asks for the user's height and responds accordingly.
func stage3() {
	var height int
	for {
		fmt.Print(ColorGreen + "boyun kaç cm senin?\n? " + ColorReset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		height, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println(ColorCyan + "Geçersiz giriş. Lütfen bir sayı girin." + ColorReset)
			continue
		}

		if height <= 99 {
			fmt.Println(ColorCyan + "Deden pigmelerin hangi kavminden lan?" + ColorReset)
		} else if height >= 100 && height <= 149 {
			fmt.Println(ColorCyan + "Kısa boylu olman önemli diil, diyeceğimi sanıyorsan yanılıyorsun pis cüce!" + ColorReset)
		} else if height >= 150 && height <= 169 {
			fmt.Println(ColorCyan + "Bacaklarına biraz gübre ektir. Faydası olur. kah!kih!koh!" + ColorReset)
		} else if height >= 170 && height <= 189 {
			fmt.Println(ColorCyan + "iyi... bana ne... sorduk mu?!" + ColorReset)
		} else if height >= 190 && height <= 209 {
			fmt.Println(ColorCyan + "Oha! fasülye sırığı!" + ColorReset)
		} else if height >= 210 {
			fmt.Println(ColorCyan + "Yok deve!! kaç santim dedik, milim demedik!" + ColorReset)
			continue
		}
		fmt.Println()
		break
	}
	stage4()
}

// stage2 asks for the user's age and responds accordingly.
func stage2() {
	var age int
	for {
		fmt.Print(ColorGreen + "kaç yaşındasın?\n? " + ColorReset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		age, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println(ColorCyan + "Geçersiz giriş. Lütfen bir sayı girin." + ColorReset)
			continue
		}
		if age <= 4 {
			fmt.Println(ColorCyan + "çok küçükmüşsün be! sen git anan gelsin lan lavuk!" + ColorReset)
		} else if age >= 5 && age <= 9 {
			fmt.Print(ColorGreen + "sütünü içtin mi yavrum?\n(e/h)? " + ColorReset)
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(strings.ToLower(choice))
			if choice == "e" {
				fmt.Println(ColorCyan + "Beynine pek etkisi olmamış, git biraz da PEPSı iç!" + ColorReset)
			} else {
				fmt.Println(ColorCyan + "bok iç o zaman!" + ColorReset)
			}
		} else if age >= 10 && age <= 17 {
			fmt.Println(ColorCyan + "iyi iyi 18ine pek bişi kalmamış... Uyu da büyü!" + ColorReset)
		} else if age >= 18 && age <= 24 {
			fmt.Print(ColorGreen + "Oy kullancanmı genç?\n(e/h)? " + ColorReset)
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(strings.ToLower(choice))
			if choice == "e" {
				fmt.Println(ColorCyan + "ver de gör ebeninkini!" + ColorReset)
			} else {
				fmt.Println(ColorCyan + "Ulan sen ne biçim Tee.Cee vatandaşısın? Hayvan!..." + ColorReset)
			}
		} else if age >= 25 && age <= 39 {
			fmt.Println(ColorCyan + "vayy! naber morruk? Nerde eski programcılar dimi mirim?" + ColorReset)
		} else if age >= 40 && age <= 59 {
			fmt.Println(ColorCyan + "Yuh! bayağı yaşlısın... yaşlılar muhattabım diildir.. Git estetik yaptır gel..." + ColorReset)
		} else if age >= 60 && age <= 98 {
			fmt.Println(ColorCyan + "Ulan bunak! Klavyeyi nası görüyon? Geber de helvanı yiyelim. hehehe!" + ColorReset)
		} else if age >= 99 {
			fmt.Println(ColorCyan + "Kafa bulma lan göt" + ColorReset)
			continue
		}
		fmt.Println()
		break
	}
	stage3()
}

// stage1 asks for the user's name and starts the conversation.
func stage1() {
	fmt.Print(ColorGreen + "senin adın ne güzelim\n? " + ColorReset)
	input, _ := reader.ReadString('\n')
	userName = strings.TrimSpace(input)

	fmt.Println()
	charCount := countCharacters(userName)

	if charCount <= 2 {
		fmt.Printf(ColorCyan + "Uzak doğudan mısın yoksa başka bir gezegenden mi?\n %d\n harfli ismini biraz zor telafuz ediyorum da...\n%c...\n%ch%s!!!\neee.. olmadı galiba... hehehehehee!\n" + ColorReset, charCount, userName[0], userName[0], userName)
	} else if charCount >= 8 {
		fmt.Println(ColorCyan + "maaşşallaaaah!\nnüfus memuru ananı babanı pek sevmiyormuş galiba!!!" + ColorReset)
		laugh()
	}

	fmt.Printf(ColorCyan + "%s...\n" + ColorReset, userName)
	stage2()
}

// stage0 is the initial welcome message.
func stage0() {
	fmt.Println(ColorCyan + "merhaba" + ColorReset)
	fmt.Println(ColorCyan + "ben karabasan..." + ColorReset)
	stage1()
}

// main is the entry point of the Go application.
func main() {
	rand.Seed(time.Now().UnixNano())
	stage0()
}
