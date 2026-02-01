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

func logMessage(level string, messages ...string) {
	f1, err := os.OpenFile("logs.json", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии json файла", err)
		return
	}
	defer f1.Close()

	_, err = f1.WriteString(fmt.Sprintf("[%s]---%s---\n", level, time.Now().Format("15.04.05 (06 01)")))
	if err != nil {
		fmt.Println("Ошибка при запись в логи ", err)
	}
	for _, s := range messages {
		_, err := f1.WriteString("- " + s + "\n")
		if err != nil {
			fmt.Println("Ошибка записи строки ", err)
		}
	}

}
func main() {
	defer func() {
		if r := recover(); r != nil {
			errorMasage := fmt.Sprint(r)
			logMessage("Паника поймана ", errorMasage)
		}
	}()

	report := Dayreport{
		Date: time.Now().Format("02.01.2006"),
	}

	var discard string

	for {
		now := time.Now()
		fmt.Printf("Сегодняшнее время в кемерово: %s\n", now.Format("15.04.05 (06 01)"))
		fmt.Print("Введите свой уровень стресса от 1 до 10:")
		_, err := fmt.Scan(&report.Stress)
		if err != nil {
			logMessage("Ошбика при вводе в report.Stress", err.Error())
			fmt.Scanln(&discard)
			continue
		}
		if report.Stress > 10 || report.Stress <= 0 {
			fmt.Println("Введеное число вне диапазона значений")
			continue
		}
		fmt.Print("Теперь введите колличество выученных сегодня слов на английском:")
		_, err = fmt.Scan(&report.Words)
		if err != nil {
			logMessage("Ошбика при вводе в report.Words", err.Error())
			fmt.Scanln(&discard)
			continue
		}
		if report.Words < 0 {
			fmt.Println("Введеное число вне диапазона значений")
			continue
		}

		fmt.Print("Введите колличество минут проебанных сегодня за ютубом:")
		_, err = fmt.Scan(&report.WatchingYoutube)
		if err != nil {
			logMessage("Ошбика при вводе в report.WatchingYoutube", err.Error())
			fmt.Scanln(&discard)
			continue
		}
		if report.WatchingYoutube < 0 {
			fmt.Println("Введеное число вне диапазона значений")
			continue
		}

		if report.Stress > 8 {
			fmt.Println("Реактор перегружен Наисрочнейше нужен отдых")

		} else if report.Words > 10 && report.WatchingYoutube == 0 {
			fmt.Println("Корабль идет на гиперскорости")
		} else {
			fmt.Println("Системы стабильный продолжай в том же духе")
		}

		f, err := os.OpenFile("Progress.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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
			fmt.Println("Данные успешно записались в формате JSON!")
		} else {
			fmt.Println("Ошибка в записи в файл", err)
		}

		break

	}
}
