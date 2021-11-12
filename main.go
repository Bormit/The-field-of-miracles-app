//Package fieldOfDreams игра
package fieldOfDreams

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

//questions структура для создания образа вопросов.
type questions struct {
	quest  string // quest вопрос.
	answer string // answer ответ на вопрос.
}

var hp, money int = 3, 0 // hp жизни игрока money деньги.
var tryAnswer string     // tryAnswer попытка отгадки слова.

//createQuest создание среза с вопросами и ответами.Функция возвращает срез с вопросами (quest) и верными ответами (answer).
func createQuest() []questions {
	slice := []questions{
		{quest: "Что использовали в Китае для глажки белья вместо утюга?", answer: "сковорода"},
		{quest: "Как у западных и южных славян назывались селение, деревня, курень?", answer: "жупа"},
		{quest: "Польский ученый-математик Гуго Дионисий Штейнгауз, прославившийся также своими афоризмами, говорил: «Комплимент женщине должен быть правдивее, чем...»", answer: "правда"},
		{quest: "В Австралии на парковках возле некоторых торговых центров по ночам и вечерам включают классическую музыку, чтобы отпугнуть кого-то. Кого?", answer: "подросток"},
		{quest: "Пельмени издавна заготавливают в форме ушек. Что символизируют такие пельмени?", answer: "послушание"},
	}
	return slice
}

//generationQuest генерация случайного вопроса и ответа.Функция получает на вход срез с вопросами и ответами ,возвращает случайный элемент из среза.В формате вопрос,ответ.
func generationQuest(slice []questions) (string, string) {
	i := rand.Intn(len(slice))
	return slice[i].quest, slice[i].answer
}

//codingAnswer кодирование загаданного слова.Функция получает на вход верный ответ и кодирует его в символы "*",записывает закодированное слово в переменную tryAnswer.
func codingAnswer(answer string) {
	for i := 0; i < len([]rune(answer)); i++ {
		tryAnswer += "*"
	}
}

//firstOut первый вывод игры.Функция получает на вход вопрос и генерирует новую строку с данными(money, hp, quest, tryAnswer),возвращает строку с первым выводом игры.
func firstOut(quest string) string {
	return fmt.Sprintf("Деньги: %d \nЖизни: %d \nВопрос: %s \nСлово: %s", money, hp, quest, tryAnswer)
}

//rollCommand ожидание ввода команды /roll от игрока.Функция получает на вход интерфейс io.Reader для чтения потока байтов из консоли,возвращает ошибку если она была.
func rollCommand(stdin io.Reader) error {
	rd := bufio.NewReader(stdin)
	input, err := rd.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	if strings.ToLower(strings.TrimSpace(input)) != "/roll" {
		return errors.New("No command /roll")
	} else {
		return nil
	}
}

//waitInputLetter ожидание ввода буквы.Функция получает на вход интерфейс io.Reader для чтения потока байтов из консоли,возвращает символ,который ввел пользователь.
func waitInputLetter(stdin io.Reader) rune {
	rd := bufio.NewReader(stdin)
	input, _, err := rd.ReadRune()
	if err != nil {
		log.Panic(err)
	}
	return input
}

//drum прокрутка барабана.Функция возвращает строку с сообщением о получении монет.
func drum() string {
	giveMoney := rand.Intn(101)
	money += giveMoney
	return fmt.Sprintf("\nВы получили %d монет", giveMoney)
}

//checkLetter проверка угаданной буквы.Функция получает на вход символ который ввел пользователь и верный ответ,возвращает строку с сообщением о верности ответа.
func checkLetter(letter rune, answer string) string {
	var buffer = []rune(tryAnswer)

	if strings.Contains(answer, string(letter)) && !strings.Contains(tryAnswer, string(letter)) {

		for i, value := range []rune(answer) {
			if value == letter {
				buffer[i] = letter
			}
		}
		tryAnswer = string(buffer)
		return "Верно!"
	} else {
		hp--
		return "Неверно!"
	}
}

func main() {
	rand.Seed(time.Now().Unix()) //Опора для генератора чисел

	for hp != 0 { //Бесконечный цикл ,пока идет игра
		var isGame = true//Запускаем игру
		tryAnswer = ""//Обнуляем слово игрока
		quest, answer := generationQuest(createQuest())//Сохраняем случайный вопрос и ответ
		codingAnswer(answer)
		fmt.Println(firstOut(quest))

		for isGame {
			err := rollCommand(os.Stdin)
			if err != nil {
				os.Exit(1)
			}
			fmt.Println("\n\n\n\n", drum())
			fmt.Println(firstOut(quest))
			fmt.Print("Введите букву -> ")
			letter := waitInputLetter(os.Stdin)
			fmt.Println(checkLetter(letter, answer))
			if hp == 0 {
				isGame = false
			} else if tryAnswer == answer {
				fmt.Println("\n\n\n\n\nВы угадали!Новое слово:")
				isGame = false
			}
		}
	}
}