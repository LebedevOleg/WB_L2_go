Что выведет программа? Объяснить вывод программы.

```
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func asChan(vs ...int) <-chan int {
   c := make(chan int)

   go func() {
       for _, v := range vs {
           c <- v
           time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
      }

      close(c)
  }()
  return c
}

func merge(a, b <-chan int) <-chan int {
   c := make(chan int)
   go func() {
       for {
           select {
               case v := <-a:
                   c <- v
               case v := <-b:
                   c <- v
           }
      }
   }()
 return c
}

func main() {

   a := asChan(1, 3, 5, 7)
   b := asChan(2, 4 ,6, 8)
   c := merge(a, b )
   for v := range c {
       fmt.Println(v)
   }
}

```

Ответ

```
Программа выведет в рандомном порядке цифры от 1 до 8, а затем начнет бесконечно выводить default значения int (0)

после того как закроются каналы a и b, select в функции merge начнет получать default значения для int (0), получается это потому, что при получении значения из канала мы не проверяем закрыт он или начнет
из-за этого мы не можем выйти из бесконечного цикла и никогда не попадем на return
```
