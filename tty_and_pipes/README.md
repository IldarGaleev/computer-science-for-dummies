# PIPE or TTY

Пример демонстрирует как приложение может менять свое поведение в зависимоти от того с чем оно работает: с TTY или с PIPE.

Код [примера](main.go) написан на языке `Go` и снабжен комментариями.

## Запуск и ожидаемый вывод

### Сборка из исходников

1. Скачать и установить go (использовался go 1.25.7)
2. Собрать пример
   ```bash
   cd tty_and_pipes/ # если вы в корневом каталоге проекта
   go build
   ```
3. Должен появиться исполняемый файл `tty_and_pipes` (`tty_and_pipes.exe` на Windows)
4. В Unix системах добавьте атрибут исполняемого файла
   ```bash
   chmod +x tty_and_pipes
   ```

### Запуск примера

1. Войдите в каталог с исполняемым файлом (если вы еще не в нем)
2. Попробуйте следующие примеры:
   - `./tty_and_pipes "hello"`
   - `./tty_and_pipes "hello" | cat`
   - `./tty_and_pipes "hello" > file1.txt`
   - `echo "hello" | ./tty_and_pipes`
   - `echo "hello" | ./tty_and_pipes | cat`
   - `echo "hello" | ./tty_and_pipes > file2.txt`
