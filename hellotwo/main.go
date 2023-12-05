// main.go
package main
import (
	"fmt"
	"os"
)

func main() {
    fmt.Println("hello world from: ", os.Getenv("NAME"))
}