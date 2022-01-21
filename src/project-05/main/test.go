package main
import (
	"fmt"
	"net/http"
	"log"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sayHello")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}