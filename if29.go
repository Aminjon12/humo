package main
import "fmt"
func main(){
	var num int
	fmt.Println("Введите число")
	fmt.Scanln(&num)
	str := ""
	positive := "положительное "
	negative := "отрицательное "
	even := "четное"
	odd := "нечетное"
	zero := "нулевое"
	if num > 0{
		str += positive
		if num % 2 == 0{
			str += even
		} else {
			str += odd
		}
	} else if num < 0 {
		str += negative
		if num % 2 == 0{
			str += even
		} else {
			str += odd
		}
	} else {
		str += zero
	}
fmt.Println(str)
}