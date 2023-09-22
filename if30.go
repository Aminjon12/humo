package main
import "fmt"
func main(){
	fmt.Println("ВВедите число")
	var num int
	fmt.Scanln(&num)
	str := ""
	decades := "двузначное "
	hundreds := "трезначное "
	even := "четное"
	odd := "нечетное"
	one := "однозначное "
	if num/10 >= 10{
		str += hundreds
		if num % 2 == 0{
			str += even
		} else {
			str += odd
		}
	} else if num/10 < 10 && num/10 > 1 {
		str += decades
		if num % 2 == 0{
			str += even
		} else {
			str += odd
		}
	} else {
		str += one
		if num % 2 == 0{
			str += even
		} else {
			str += odd
		}
	}
	fmt.Println(str)
}