Что выведет программа? Объяснить вывод программы.

```
import (
    "fmt"
    "os"
)

func Foo() error {
    var err *os.PathError = nil
    return err
}

func main() {
    err := Foo()
    fmt.Println(err)
    fmt.Println(err == nil)
}

```

Ответ

```
Программа выведет <nil>, false

<!-- todo дописать почему -->

```
