package main

import "fmt"

/*func main() {
	//today := "Wednesday"

	switch today := "Wednesday"; today {
	case "Saturday":
		fmt.Println("Today is Saturday")
	case "Monday", "Tuesday":
		fmt.Println("Today is Weekdays")
		fallthrough

	default:
		fmt.Println("วันนี้ไม่ใช่วันของมึงค่ะ!")
	}
}*/

//WORKSHOP

func main() {

	switch rate := 8.4; {
	case rate < 5.0:
		fmt.Println("Disappointed")
	case rate >= 5.0 && rate < 7.0:
		fmt.Println("Normal")
	case rate >= 7.0 && rate < 10.0:
		fmt.Println("Good")

	default:
		fmt.Println("WTF!")
	}
}
