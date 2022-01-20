package main
import "fmt"

func getSum() func(int) int {
	sum := 0;
	return func(i int) int {
		sum = sum + i
		return sum
	}
}


// 闭包
func main() {
	f := getSum();
	fmt.Println(f(1))	// 1
	fmt.Println(f(2))	// 3
	fmt.Println(f(3))	// 6
	fmt.Println(f(4))	// 10
}