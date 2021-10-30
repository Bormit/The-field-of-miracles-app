package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var slice = []questions{
	{quest: "Что использовали в Китае для глажки белья вместо утюга?", answer: "сковорода"},
	{quest: "Как у западных и южных славян назывались селение, деревня, курень?", answer: "жупа"},
	{quest: "Польский ученый-математик Гуго Дионисий Штейнгауз, прославившийся также своими афоризмами, говорил: «Комплимент женщине должен быть правдивее, чем...»", answer: "правда"},
	{quest: "В Австралии на парковках возле некоторых торговых центров по ночам и вечерам включают классическую музыку, чтобы отпугнуть кого-то. Кого?", answer: "подросток"},
	{quest: "Пельмени издавна заготавливают в форме ушек. Что символизируют такие пельмени?", answer: "послушание"},
}

func TestCreateQuest(t *testing.T) {
	if !reflect.DeepEqual(createQuest(), slice) {
		t.Error("Неправильно созданный срез с вопросами и ответами")
	}
}

func TestGenerationQuest(t *testing.T) {
	var isQuest, isAnswer = false, false
	quest, answer := generationQuest(slice)

	for _, value := range slice {
		if value.quest == quest {
			isQuest = true
		}
		if value.answer == answer {
			isAnswer = true
		}
	}
	if !isQuest || !isAnswer {
		t.Error("Ошибка в генерации случайного вопроса и ответа")
	}
}

func TestCodingAnswer(t *testing.T) {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(slice))
	var sum int

	codingAnswer(slice[i].answer)
	for j := 0; j < len([]rune(slice[i].answer)); j++ {
		if tryAnswer[j] == '*' {
			sum++
		}
	}
	if sum != len([]rune(slice[i].answer)) {
		t.Error("Ошибка в кодировании загаданного слова")
	}
}

func TestFirstOut(t *testing.T) {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(slice))

	if firstOut(slice[i].quest) != fmt.Sprintf("Деньги: %d \nЖизни: %d \nВопрос: %s \nСлово: %s", money, hp, slice[i].quest, tryAnswer) {
		t.Error("Ошибка в первом выводе игры")
	}
}