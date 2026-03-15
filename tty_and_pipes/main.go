package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

/* sendOut метод собирает наш вывод */
func sendOut(v ...any) {
	/* выясняем куда идет вывод: в терминал или в pipe */
	isTermOut := term.IsTerminal(int(os.Stdout.Fd()))

	v = append(v, "to") /* собираем аргументы для Println */
	if isTermOut {      /* добавляем метку в зависимости от того, куда пишем */
		v = append(v, "TTY")
	} else {
		v = append(v, "PIPE")
	}

	fmt.Println(v...) /* выводим результат в stdout */
}

func main() {
	var data, from string
	/* выясняем откуда читаем данные: из терминала или из pipe */
	isTermIn := term.IsTerminal(int(os.Stdin.Fd()))

	if isTermIn { /* если мы получаем данные из терминала... */
		if len(os.Args) > 1 { /* и нам передали аргумент */
			data = os.Args[1] /* записываем первый аргумент в данные */
		} else {
			data = "<nil>" /* иначе сообщаем, что данных не передавали */
		}
		from = "TTY" /* запоминаем что данные читались из терминала */
	} else { /* если мы получаем данные не из терминала... */
		_, err := fmt.Scan(&data) /* читаем их из stdin */
		if err != nil {           /* если произошла ошибка */
			fmt.Fprint(os.Stderr, err.Error()) /* выводим ее в stderr */
			os.Exit(1)                         /* завершаем приложение с кодом ошибки 1 */
		}
		from = "PIPE" /* запоминаем что данные читались из pipe */
	}

	sendOut(data, "from", from) /* выводим результат */
}
