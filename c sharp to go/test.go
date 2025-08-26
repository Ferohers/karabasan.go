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
	fmt.Println(jokes[jokeIndex])
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
	fmt.Println(laughs[getRandomInt(len(laughs))])
}

// actDumb has a 50% chance of printing a "dumb" joke.
func actDumb() {
	if getRandomInt(2) == 1 {
		fmt.Println("\ngeri zekalı taklidi yap bakiim...\nTamam tamam bukadar yeter!!!\n")
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
		fmt.Println(swears[0])
	}
	if getRandomInt(2) == 1 {
		fmt.Println(swears[1])
	}
	if getRandomInt(2) == 1 {
		fmt.Println(swears[2])
	}
	if getRandomInt(2) == 1 {
		fmt.Println(swears[3])
	}
	if getRandomInt(2) == 1 {
		fmt.Println(swears[4])
	}
}

// isVowel checks if a given character is a Turkish vowel.
func isVowel(r rune) bool {
	vowels := "aıueöüio"
	return strings.ContainsRune(vowels, r)
}

// stage10 concludes the game with a final joke.
func stage10() {
	fmt.Println("\nşimdik sana bi fıkra daha:\n")
	sayJoke()
	fmt.Println("Çıkmak için bir tuşa basın.")
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

	fmt.Println("şimdik sen bi sayı tut, ben bulmaya çalışiim. Ama dürüst ol.")
	fmt.Println("tahminimde yükselmen gerekirse 'y', düşmem gerekirse 'd' ile yanıt ver.")
	fmt.Println("sayıyı bulursam 'b' ile yanıt vermen yeterli.")

	for {
		guessCount++
		fmt.Printf(" %d  ??\n? ", guess)

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

	fmt.Printf(" %d  tahminde bildim...\n", guessCount)
	if guessCount < score {
		fmt.Println("kodum! kodum! kodum! hehehehe!")
	} else if guessCount > score {
		fmt.Println("lanet olsun! beni geçtin! %100 hile yapmışsındır!")
	} else {
		fmt.Println("hmm... eşitiz galiba...")
	}

	stage10()
}

// stage8 is the number guessing game where the user guesses the computer's number.
func stage8() {
	target := getRandomInt(100) + 1
	var guess int
	guessCount := 0

	fmt.Printf("%s,\n gel senlen oyun oynayak...\nben şimdik 1 ilen 100 arası bi sayı tutiim...\ntuttum.\n", userName)

	for {
		guessCount++
		fmt.Print("tahmin et bakalım..? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var err error
		guess, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Geçersiz giriş. Lütfen bir sayı girin.")
			continue
		}

		if guess == target {
			if guessCount <= 3 {
				fmt.Printf(" %d  tahminde nası bildin lan? walla brawo!!\n", guessCount)
			} else if guessCount <= 5 {
				fmt.Printf(" %d . denemede buldun!! tebrik etmek lazım şindi seni...\n", guessCount)
			} else if guessCount <= 10 {
				fmt.Printf(" %d tahminde buldun.. eh..\n", guessCount)
			} else if guessCount <= 20 {
				fmt.Printf("NİHAYET!!!  bişey  %d  kere sorulmaz ki ama, dimi?!\n", guessCount)
			} else if guessCount <= 30 {
				fmt.Printf("bir an ümidimi kesmiştim! neytse ki  %d  kerede buldun! aferin!\n", guessCount)
			} else {
				fmt.Printf(" %d \ntahminde bulundun...  sen,\n1- Türkçe bilmiyorsun...\n2- Klavye kullanmasını bilmiyorsun...\n3- ya da cinsel yönden bazısorunların var!!!\nE M B E S İ L !\n", guessCount)
			}
			score = guessCount
			stage9()
			break
		} else {
			if guess < 1 || guess > 100 {
				fmt.Println("Abartma! abartma!  1-100 arası dedik!")
			} else if guess < target {
				if target-guess > 20 {
					fmt.Println("çık çık")
				} else {
					fmt.Println("yaklaştın, acık daa çık!")
				}
			} else { // guess > target
				if guess-target > 20 {
					fmt.Println("aşşalara gel aşşalara")
				} else {
					fmt.Println("biraz daa düş!")
				}
			}
		}
	}
}

// stage7 asks for the user's hometown and responds based on the last vowel.
func stage7() {
	fmt.Printf("memleket nere %s?\n? ", userName)
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
			fmt.Printf("madem %slusun,\n buralara ne b*k yemeye geldin?! Ayrıca\n%sdan\n   adam falan çıkmaz!\n", hometown, hometown)
		case 'ü', 'ö':
			fmt.Printf("heheheh!%sden\n top çıkarmış diyolar!?!", hometown)
		case 'a', 'ı':
			fmt.Printf("naaaber pis\n%slı!\n", hometown)
		case 'e', 'i':
			fmt.Printf("nea!? %sden\n     adam çıkmaz ki beah!!!  hihöhöhö!!\n", hometown)
		}
	}

	laugh()
	fmt.Printf("\nneyse %s,\n kusura bakma...\n", userName)

	stage8()
}

// stage6 prints a joke and a proverb.
func stage6() {
	fmt.Println("bak sana şindi konuyla ilgili bir fıkra...")
	sayJoke()
	laugh()

	proverbs := []string{
		"yani sakla samanı gelir zamanı.",
		"yani arkadaşlarımızı dikkatli seçmemiz lazım.",
		"buradan alınacak ders: Göte giren şemsiye açılmaz..",
	}

	fmt.Printf("\n%s\n", proverbs[getRandomInt(len(proverbs))])
	laugh()
	fmt.Println()
	stage7()
}

// stage5 contains a series of random questions.
func stage5() {
	// Question 1: Eyes
	if getRandomInt(2) == 1 {
		fmt.Printf("%s!\n", userName)
		fmt.Print("sana gözlerinin çok güzel olduğunu söyleyen olmuşmuydu hiç\n(e/h)? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			fmt.Println("yalan söylemiş!")
			laugh()
		} else {
			fmt.Println("doğrudur. çünkü gözlerin güzel diil!")
			laugh()
		}
	}

	// Question 2: Money
	if getRandomInt(2) == 1 {
		fmt.Printf("\nyavrum\n%s\n", userName)
		fmt.Print("ayda 50 milyon kazanmak istermisin?\n(e/h)? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			fmt.Println("o zaman Ay'a gitmen lazım...")
			laugh()
		} else {
			fmt.Println("iyi... zaten Ay'da sağlıklı çalışabileceğini sanmıyordum.")
			laugh()
		}
	}

	// Question 3: Name Origin
	if getRandomInt(2) == 1 {
		fmt.Printf("\n%s\n", userName)
		fmt.Print("adı nerden geliyo?\n? ")
		reader.ReadString('\n') // Just read and discard the user's input
		fmt.Println("üüüü! baya uzaktan geliyomuş!")
		laugh()
	}

	// Question 4: Holding a number
	if getRandomInt(2) == 1 {
		fmt.Printf("\n%s\n", userName)
		fmt.Print("bi sayı tut.\ntuttunmu (e/h)\n? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			fmt.Println("şimdi de bırak!")
			laugh()
		} else {
			fmt.Println("bi sayıyı tutamadın allah belanı versin")
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
			fmt.Printf("\n%s, sana kısaca %s diyebilirmiyim??\n(e/h)? ", userName, nickname)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(strings.ToLower(input))
			if input == "e" {
				fmt.Println("iyi... ama ben demek istemiyorum!")
				laugh()
			} else {
				fmt.Printf("%s! %s! %s!\n", nickname, nickname, nickname)
				laugh()
			}
		}
	}

	// Question 6: How are you?
	if getRandomInt(2) == 1 {
		fmt.Printf("\nnasılsınız lan\n%s?\n", userName)
		fmt.Print("iyimisin ki (e/h)\n? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			randChoice := getRandomInt(3)
			if randChoice == 0 {
				fmt.Println("niye iyisin? oturduğun yere bir bak bakiim...\njoysitick falan unutmuş olmasınlar?")
			} else if randChoice == 1 {
				fmt.Printf("iyi iyi... sen iyi olmaya devam et\n%s!\nuyu da büyü!\n", userName)
			} else {
				fmt.Printf("böyle bir hayatta nasıl iyi oluyorsunuz ki lan\n%s?\nbize de söyle yolunu biz de iyi olalım..\n", userName)
			}
		} else {
			randChoice := getRandomInt(3)
			if randChoice == 0 {
				fmt.Println("bana ne lan! geber!")
			} else if randChoice == 1 {
				fmt.Println("iyi iyi allah kötülük versin! he he he !!")
			} else {
				fmt.Printf("derdini anlat bana! açıl bana yavrucuum! utanma ben doktorum...\nKötü olmana sebep olan şey nedir %s", userName)
				reader.ReadString('\n') // Read and discard
				fmt.Println("\n??\nhahahahahahahaha!!! git allasen yaw! dert  ettiğin şeye bak!")
			}
		}
		laugh()
	}

	// Question 7: Student
	if getRandomInt(2) == 1 {
		fmt.Printf("\nneyse... %s\n", userName)
		fmt.Print("      öğrencimisin?\n? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "e" {
			randChoice := getRandomInt(2)
			if randChoice == 0 {
				fmt.Println("wah! wah! wah! çok üzüldüm.. ailenin haberi varmı? ha!haha!!hohoho!!!\n")
			} else {
				fmt.Println("nerde öğrencisin? okulda mı?? hihohohohhohohooo!!!\nespri konuşlandırdım!!\n")
			}
		} else {
			randChoice := getRandomInt(2)
			if randChoice == 0 {
				fmt.Println("ulan insan en azından askerden yırtmak için öğrenci olur! Ama sen, tıss!\n")
			} else {
				fmt.Print("hangi işle meşgulsun o vakit?\n? ")
				reader.ReadString('\n') // Read and discard
				fmt.Println("siktir lan göt! cümle alem senin ne mal olduğunu biliyor.\n")
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
		fmt.Print("oldu olcak kilonu da söyle bari... çok umurumda ya...\n? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		weight, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Geçersiz giriş. Lütfen bir sayı girin.")
			continue
		}

		if weight <= 39 {
			fmt.Println("Rüzgarlı havada dışarı falan çıkma hehehe!")
			actDumb()
		} else if weight >= 40 && weight <= 59 {
			fmt.Println("o kadar yemiş yersen ishal de olursun, kabız da!")
			actDumb()
		} else if weight >= 60 && weight <= 79 {
			fmt.Println("sen normalsin o yüzden dalga geçmiicem... noormaal! noormaal! hehehe!!")
			actDumb()
		} else if weight >= 80 && weight <= 99 {
			randChoice := getRandomInt(3)
			if randChoice == 0 {
				fmt.Println("Lütfen, oturduğun koltuk sağlam kalsın!")
			} else if randChoice == 1 {
				fmt.Println("Maaşşallaaah! damızlıkmısın? hangi çiftlikte yetiştin? keh!keh!keh!!.")
			} else {
				fmt.Println("Duba! dikkat et benim üstüme düşme!")
			}
			actDumb()
		} else if weight >= 100 {
			fmt.Println("Anlamıştım... 2 saattir klavyenin anasını ağlattın")
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
		fmt.Print("boyun kaç cm senin?\n? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		height, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Geçersiz giriş. Lütfen bir sayı girin.")
			continue
		}

		if height <= 99 {
			fmt.Println("Deden pigmelerin hangi kavminden lan?")
		} else if height >= 100 && height <= 149 {
			fmt.Println("Kısa boylu olman önemli diil, diyeceğimi sanıyorsan yanılıyorsun pis cüce!")
		} else if height >= 150 && height <= 169 {
			fmt.Println("Bacaklarına biraz gübre ektir. Faydası olur. kah!kih!koh!")
		} else if height >= 170 && height <= 189 {
			fmt.Println("iyi... bana ne... sorduk mu?!")
		} else if height >= 190 && height <= 209 {
			fmt.Println("Oha! fasülye sırığı!")
		} else if height >= 210 {
			fmt.Println("Yok deve!! kaç santim dedik, milim demedik!")
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
		fmt.Print("kaç yaşındasın?\n? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		age, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Geçersiz giriş. Lütfen bir sayı girin.")
			continue
		}
		if age <= 4 {
			fmt.Println("çok küçükmüşsün be! sen git anan gelsin lan lavuk!")
		} else if age >= 5 && age <= 9 {
			fmt.Print("sütünü içtin mi yavrum?\n(e/h)? ")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(strings.ToLower(choice))
			if choice == "e" {
				fmt.Println("Beynine pek etkisi olmamış, git biraz da PEPSı iç!")
			} else {
				fmt.Println("bok iç o zaman!")
			}
		} else if age >= 10 && age <= 17 {
			fmt.Println("iyi iyi 18ine pek bişi kalmamış... Uyu da büyü!")
		} else if age >= 18 && age <= 24 {
			fmt.Print("Oy kullancanmı genç?\n(e/h)? ")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(strings.ToLower(choice))
			if choice == "e" {
				fmt.Println("ver de gör ebeninkini!")
			} else {
				fmt.Println("Ulan sen ne biçim Tee.Cee vatandaşısın? Hayvan!...")
			}
		} else if age >= 25 && age <= 39 {
			fmt.Println("vayy! naber morruk? Nerde eski programcılar dimi mirim?")
		} else if age >= 40 && age <= 59 {
			fmt.Println("Yuh! bayağı yaşlısın... yaşlılar muhattabım diildir.. Git estetik yaptır gel...")
		} else if age >= 60 && age <= 98 {
			fmt.Println("Ulan bunak! Klavyeyi nası görüyon? Geber de helvanı yiyelim. hehehe!")
		} else if age >= 99 {
			fmt.Println("Kafa bulma lan göt")
			continue
		}
		fmt.Println()
		break
	}
	stage3()
}

// stage1 asks for the user's name and starts the conversation.
func stage1() {
	fmt.Print("senin adın ne güzelim\n? ")
	input, _ := reader.ReadString('\n')
	userName = strings.TrimSpace(input)

	fmt.Println()
	charCount := countCharacters(userName)

	if charCount <= 2 {
		fmt.Printf("Uzak doğudan mısın yoksa başka bir gezegenden mi?\n %d\n harfli ismini biraz zor telafuz ediyorum da...\n%c...\n%ch%s!!!\neee.. olmadı galiba... hehehehehee!\n", charCount, userName[0], userName[0], userName)
	} else if charCount >= 8 {
		fmt.Println("maaşşallaaaah!\nnüfus memuru ananı babanı pek sevmiyormuş galiba!!!")
		laugh()
	}

	fmt.Printf("%s...\n", userName)
	stage2()
}

// stage0 is the initial welcome message.
func stage0() {
	fmt.Println("merhaba")
	fmt.Println("ben karabasan...")
	stage1()
}

// main is the entry point of the Go application.
func main() {
	rand.Seed(time.Now().UnixNano())
	stage0()
}
