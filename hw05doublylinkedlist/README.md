# Двусвязный список

Цель:  Реализовать двусвязный список https://en.wikipedia.org/wiki/Doubly_linked_list ​
Завести в репозитории отдельный пакет (модуль) для этого ДЗ
Реализовать типы List и Item (см. ниже) и методы у них.
Написать unit-тесты проверяющие работу всех методов.
                                      
Ожидаемые типы (псевдокод):
```bash
List: // тип контейнер
    Len() // длинна списка
    First() // первый Item
    Last() // последний Item
    PushFront(v interface{}) // добавить значение в начало
    PushBack(v interface{}) // добавить значение в конец
    Remove(i Item) // удалить элемент

Item: // элемент списка
    Value() interface{} // возвращает значение
    Next() *Item // следующий Item
    Prev() *Item // предыдущий
```

Критерии оценки: 
* Сложность всех операций должна быть O(1), т.е. не должно быть мест 
где осуществляется полный обход списка.
* Пакет должен проходить все тесты.
* Код должен проходить проверки go vet и golint
* У преподавателя должна быть возможность скачать и проверить пакет с 
помощью go get / go test

## Usage

To update dependencies, run tests and lint:
```bash
make
```

### Other make commands

* `bootstrap` - download golangci-lint, if it is not exists
* `lint` - run golangci linter
* `test` - run all tests
* `vendor` - actualize and update dependencies
