package main
import "fmt"
func main(){
	var a int
	var b int
	fmt.Println("Введите числа")
	fmt.Scanln(&a, &b)
	for i := a ; i <= b; i++ {
		for j := 0; j < i-a+1; j++ {
			fmt.Println(i)
		}
	}
}