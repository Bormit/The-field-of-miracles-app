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

func returnAll(quest string,answer string) string {
	var buf string
	//for i:=0;i<len(answer);i++{
	//	buf+="*"
	//}

	return fmt.Sprintf("Деньги: %d \nЖизни: %d \nВопрос: %s\nСлово: %s",money, hp, quest, buf)
}

func main() {
	rand.Seed(time.Now().Unix()) //Опора для генератора чисел
	//for isGame { //Бесконечный цикл ,пока идет игра
	quest, answer := generationQuest(createQuest())
	fmt.Println(quest,answer,len(answer))
	//fmt.Println(returnAll(quest,answer))
	//}
}
