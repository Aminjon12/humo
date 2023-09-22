package main
import "fmt"
	
func main(){
	var year int
	fmt.Println("Введите год")
	fmt.Scanln(&year)
	if year % 4 == 0 {
		if year % 100 == 0{
			if year % 400 == 0{
				fmt.Println("366")
			} else {
				fmt.Println("365")
			}
		} else {
			fmt.Println("366")
		}
	} else {
		fmt.Println("365")
	}
	}
