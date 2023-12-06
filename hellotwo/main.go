// main.go
package main
import (
	"fmt"
	"net/http"
)


func henlo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}


func main() {
    http.HandleFunc("/", henlo)
}