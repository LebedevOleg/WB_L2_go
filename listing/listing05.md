Что выведет программа? Объяснить вывод программы.

```
package main

type customError struct {
     msg string
}

func (e *customError) Error() string {
    return e.msg
}

func test() *customError {
     {
         // do something
     }
     return nil
}

func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}

```

Ответ

```
Программа выведет строку "error"
дело в том, что возвращаемое значение из функции test() является ссылка на созданную структуру
тоесть нам возвращается ссылка на объект структуры customError уже в котором лежит значение nil
Поскольку сама ссылка не nil нам и выводится ошибка
```
