package main
import "fmt"

func main() {
	defer fmt.Println("111")
	defer fmt.Println("222")

	fmt.Println("333")
	fmt.Println("444")
}