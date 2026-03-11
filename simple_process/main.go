package main

import (
	"fmt"
	"os"
)

/* форматированный вывод в терминал */
func printHead(h string, f string, v ...any) {

	/*
		Управляющие команды для VT режима

		\033[ - escape последовательность (указывает что дальше будут команды терминала (через точку с запятой), а не текст)
		0  - сбросить все стили
		1  - полужирный
		32 - зеленый текст
		m  - конец управляющей последовательности
	*/

	fmt.Printf("\033[1;32m"+h+"\033[0m"+f, v...)
}

/* Точка входа в приложение. Адрес этого метода будет указан как начало программы */
func main() {

	cwd, err := os.Getwd() /* Получаем путь к рабочей директории */
	if err != nil {
		cwd = "==unknown=="
	}

	pid := os.Getpid() /* Получаем ID текущего процесса */
	uid := os.Getuid() /* Получаем ID пользователя от которого запущен процесс (в Windows вернет -1) */

	argv := os.Args      /* Получаем список переданных аргументов командной строки */
	envs := os.Environ() /* Получаем список переменных окружения */

	/* Выводим содержимое в терминал */
	printHead("CWD", ": %s\n", cwd)
	printHead("PID", ": %d\n", pid)
	printHead("UID", ": %d\n", uid)
	fmt.Println()

	printHead("argv", " [count: %d]:\n", len(argv))
	for i, arg := range argv {
		fmt.Printf("%d: %s\n", i, arg)
	}
	fmt.Println()

	printHead("Environment Variable", " [count: %d]:\n", len(envs))
	for i, env := range envs {
		fmt.Printf("%03d: %s\n", i, env)
	}

}
