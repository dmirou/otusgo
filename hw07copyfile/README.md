# Копирование файлов

Цель:  Реализовать утилиту копирования файлов. 
Утилита должна принимать следующие аргументы 
* файл источник (From) 
* файл копия (To) 
* Отступ в источнике (Offset), по умолчанию - 0 
* Количество копируемых байт (Limit), по умолчанию - весь файл из From 

Выводить в консоль прогресс копирования в %, например с помощью github.com/cheggaaa/pb  
Программа может НЕ обрабатывать файлы, у которых не известна длинна (например /dev/urandom).  

* Завести в репозитории отдельный пакет (модуль) для этого ДЗ
* Реализовать функцию вида Copy(from string, to string, limit int, offset int) error
* Написать unit-тесты на функцию Copy
* Реализовать функцию main, анализирующую параметры командной строки и вызывающую Copy
* Проверить установку и работу утилиты руками

Критерии оценки: 
* Функция должна проходить все тесты
* Все необходимые для тестов файлы должны создаваться в самом тесте
* Код должен проходить проверки go vet и golint
* У преподавателя должна быть возможность скачать, проверить и установить пакет с помощью go get / go test / go install

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
* `vendor` - actualize and update dependencies
