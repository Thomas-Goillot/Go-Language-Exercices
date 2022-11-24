package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now()
	fmt.Println("Nous sommes le", day.Day(), getMonthName(int(day.Month())), "il est", day.Hour(), "h", addZeroIfNecessary(day.Minute()))
	fmt.Printf("Nous sommes le %d %s, il est %dh%s", day.Day(), getMonthName(int(day.Month())), day.Hour(), addZeroIfNecessary(day.Minute()))
}

func getMonthName(month int) string {
	switch month {
	case 1:
		return "Janvier"
	case 2:
		return "Février"
	case 3:
		return "Mars"
	case 4:
		return "Avril"
	case 5:
		return "Mai"
	case 6:
		return "Juin"
	case 7:
		return "Juillet"
	case 8:
		return "Août"
	case 9:
		return "Septembre"
	case 10:
		return "Octobre"
	case 11:
		return "Novembre"
	case 12:
		return "Décembre"
	default:
		return "Inconnu"
	}
}

func addZeroIfNecessary(number int) string {
	if number <= 9 {
		return "0" + fmt.Sprintf("%d", number)
	}
	return fmt.Sprintf("%d", number)
}
