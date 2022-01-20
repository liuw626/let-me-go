package main
import "fmt"

func main() {
	defer func() {
		err := recover()
		if (err != nil) {
			fmt.Println("err:", err)
		}
		fmt.Println("function finished")
	}()

	i := 1
	j := 0

	k := i / j
	fmt.Println("k:", k)
}