package foo

import "fmt"

func Greet(name string) string {
    return fmt.Sprintf("%s, 你好！version 1.0.0", name)
}
