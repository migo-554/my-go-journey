package main

import (
	"fmt"
	"os"
	"time"
)

func Proverka(err error) string {
	if err != nil {
		fmt.Println("Нужно вводить числа попробуй еще раз")
		var discard string
		fmt.Scanln(&discard)
		return "not"
	} else {
		return "ok"
	}
}

func main() {
	now := time.Now()
	for {
		fmt.Printf("Сегодняшнее время в кемерово: %s\n", now.Format("15:04:03"))
		var stress int
		fmt.Print("Введите свой уровень стресса от 1 до 10:")
		_, err := fmt.Scan(&stress)
		StatStress := Proverka(err)
		if StatStress == "not" {
			continue
		}

		var words int
		fmt.Print("Теперь введите колличество выученных сегодня слов на английском:")
		_, err1 := fmt.Scan(&words)
		StatWords := Proverka(err1)
		if StatWords == "not" {
			continue
		}

		var watchingYoutube int
		fmt.Print("Введите колличество минут проебанных сегодня за ютубом:")
		_, err2 := fmt.Scan(&watchingYoutube)
		StatYoutube := Proverka(err2)
		if StatYoutube == "not" {
			continue
		}

		if stress > 8 {
			fmt.Println("Реактор перегружен Наисрочнейше нужен отдых")

		} else if words > 10 && watchingYoutube == 0 {
			fmt.Println("Корабль идет на гиперскорости")
		} else {
			fmt.Println("Системы стабильный продолжай в том же духе")
		}

		f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Ошибка доступа к журналу ", err)
			return
		}
		defer f.Close()

		report := fmt.Sprintf("[%s] Стресс: %d, Слов %d\n ", time.Now().Format("01.02.2006"), stress, words)
		_, err = f.WriteString(report)
		if err != nil {
			fmt.Println("Ошибка с записью ", err)
			return
		} else {
			fmt.Println("Данные успешно записались!")
		}

		break

	}
}
