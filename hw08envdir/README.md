# Утилита envdir

## Цель  

Реализовать утилиту envdir на Go. Эта утилита позволяет запускать программы получая переменные 
окружения из определенной директории. См man envdir Пример go-envdir /path/to/evndir command arg1 arg2

1. Завести в репозитории отдельный пакет (модуль) для этого ДЗ
2. Реализовать функцию вида ReadDir(dir string) (map[string]string, error), которая сканирует указанный 
каталог и возвращает все переменные окружения, определенные в нем.
3. Реализовать функцию вида RunCmd(cmd []string, env map[string]string) int , которая запускает программу 
с аргументами (cmd) c переопределнным окружением.
4. Реализовать функцию main, анализирующую аргументы командной строки и вызывающую ReadDir и RunCmd

5. Протестировать утилиту.
    * Тестировать можно утилиту целиком с помощью shell скрипта, а можно написать 
    unit тесты на отдельные функции.

## Критерии оценки
1. Стандартные потоки ввода/вывода/ошибок должны пробрасываться в вызываемую программу.
2. Код выхода утилиты envdir должен совпадать с кодом выхода программы.
4. Код должен проходить проверки go vet и golint
5. У преподавателя должна быть возможность скачать и установить пакет с помощью go get / go install 

## Usage

To compile and run binary use:
```bash
make
```

### Other make commands

* `clean` - remove compiled binary file
* `compile` - compile binary file 
* `run` - compile and run the binary file
* `bootstrap` - download golangci-lint, if it is not exists
* `lint` - run golangci linter
* `test` - run all tests
* `demo` - run binary file with demo arguments
* `vendor` - actualize and update dependencies
