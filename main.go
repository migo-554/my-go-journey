package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Dayreport struct {
	Date            string
	Stress          int
	Words           int
	WatchingYoutube int
}

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

	report := Dayreport{
		Date: time.Now().Format("02.01.2006"),
	}
	for {

		fmt.Printf("Сегодняшнее время в кемерово: %s\n", now.Format("15:04:03"))
		fmt.Print("Введите свой уровень стресса от 1 до 10:")
		_, err := fmt.Scan(&report.Stress)
		StatStress := Proverka(err)
		if StatStress == "not" {
			continue
		}

		fmt.Print("Теперь введите колличество выученных сегодня слов на английском:")
		_, err1 := fmt.Scan(&report.Words)
		StatWords := Proverka(err1)
		if StatWords == "not" {
			continue
		}

		fmt.Print("Введите колличество минут проебанных сегодня за ютубом:")
		_, err2 := fmt.Scan(&report.WatchingYoutube)
		StatYoutube := Proverka(err2)
		if StatYoutube == "not" {
			continue
		}

		if report.Stress > 8 {
			fmt.Println("Реактор перегружен Наисрочнейше нужен отдых")

		} else if report.Words > 10 && report.WatchingYoutube == 0 {
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

		jsonData, err := json.Marshal(report)
		if err != nil {
			fmt.Println("Ошибка кодирования JSON:", err)
			return
		}

		_, err = f.Write(jsonData)
		if err == nil {
			f.WriteString("\n")
			fmt.Println("Данне успешно записались в формате JSON!")
		} else {
			fmt.Println("Ошибка в записи в файл", err)
		}

		break

	}
}
