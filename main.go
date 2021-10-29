package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type questions struct { //Структура для создания образа вопросов
	quest  string //Вопрос
	answer string // Ответ на вопрос
}

var hp, money int = 3, 0 // Жизни игрока и деньги
var tryAnswer string //Попытка отгадки слова

func createQuest() []questions { //Создание среза с вопросами и ответами
	slice := []questions{
		{quest: "Что использовали в Китае для глажки белья вместо утюга?", answer: "сковорода"},
		{quest: "Как у западных и южных славян назывались селение, деревня, курень?", answer: "жупа"},
		{quest: "Польский ученый-математик Гуго Дионисий Штейнгауз, прославившийся также своими афоризмами, говорил: «Комплимент женщине должен быть правдивее, чем...»", answer: "правда"},
		{quest: "В Австралии на парковках возле некоторых торговых центров по ночам и вечерам включают классическую музыку, чтобы отпугнуть кого-то. Кого?", answer: "подросток"},
		{quest: "Пельмени издавна заготавливают в форме ушек. Что символизируют такие пельмени?", answer: "послушание"},
	}
	return slice
}

func generationQuest(slice []questions) (string, string) { // Генерация случайного вопроса и ответа
	i := rand.Intn(len(slice)) //Генерация индекса для выбора случайного элемента среза
	return slice[i].quest, slice[i].answer
}

func codingAnswer(answer string){ //Кодирование слова загаданного
	for i:=0;i<len([]rune(answer));i++{
		tryAnswer+="*"
	}
}


func firstOut(quest string)string{ //Первый вывод игры
	return fmt.Sprintf("Деньги: %d \nЖизни: %d \nВопрос: %s \nСлово: %s",money,hp,quest,tryAnswer)
}

func rollCommand()error{ //Ожидание ввода команды /roll от игрока
	rd:=bufio.NewReader(os.Stdin)
	input,err:=rd.ReadString('\n')
	if err!=nil{
		log.Panic(err)
	}
	if strings.ToLower(strings.TrimSpace(input))!="/roll"{
		return errors.New("No command /roll")
	}else{
		return nil
	}
}

func waitInputLetter()rune{//Ожидание ввода буквы
	fmt.Print("Введите букву -> ")
	rd:=bufio.NewReader(os.Stdin)
	input,_,err:=rd.ReadRune()
	if err!=nil{
		log.Panic(err)
	}
	return input
}

func drum(){//Прокрутка барабана
	giveMoney:=rand.Intn(101)
	money+=giveMoney
	fmt.Println("\n\n\n\n",fmt.Sprintf("\nВы получили %d монет",giveMoney))
}

func checkLetter(letter rune,answer string)string{//Проверка угаданной буквы
	var buffer = []rune(tryAnswer)

	if strings.Contains(answer,string(letter)) && !strings.Contains(tryAnswer,string(letter)){

		for i,value:=range []rune(answer){
			if value==letter{
				buffer[i]=letter
			}
		}
		tryAnswer=string(buffer)
		return "Верно!"
	}else{
		hp--
		return "Неверно!"
	}
}


func main() {
	rand.Seed(time.Now().Unix()) //Опора для генератора чисел

	for hp!=0 { //Бесконечный цикл ,пока идет игра
		var isGame = true
		tryAnswer=""
		quest, answer := generationQuest(createQuest())
		codingAnswer(answer)
		fmt.Println(firstOut(quest))

		for isGame {
			err := rollCommand()
			if err != nil {
				os.Exit(1)
			}
			drum()
			fmt.Println(firstOut(quest))
			letter := waitInputLetter()
			fmt.Println(checkLetter(letter, answer))
			if hp==0{
				isGame=false
			}else if tryAnswer==answer{
				fmt.Println("\n\n\n\n\nВы угадали!Новое слово:")
				isGame=false
			}
		}
	}
}