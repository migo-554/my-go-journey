package main

import "fmt"

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
	for {
		var stress int
		fmt.Println("Введите свой уровень стресса от 1 до 10:")
		_, err := fmt.Scan(&stress)
		StatStress := Proverka(err)
		if StatStress == "not" {
			continue
		}

		var words int
		fmt.Println("Теперь введите колличество выученных сегодня слов на английском")
		_, err1 := fmt.Scan(&words)
		StatWords := Proverka(err1)
		if StatWords == "not" {
			continue
		}

		var watchingYoutube int
		fmt.Println("Введите колличество минут проебанных за ютубом")
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
			fmt.Printf("Системы стабильный продолжай в том же духе")
		}
		break

	}
}
