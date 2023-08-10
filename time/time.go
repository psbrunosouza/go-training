package good_times

import (
	"fmt"
	"time"
)

/*
	criar uma data nova
*/
func CreateADate() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("02/01/2006"))
	fmt.Println(time.Now().Year())
	value, _ := time.Parse("02/01/2006 03:04:05", "27/04/2022 07:34:45")
	result := value.Format("02/01/2006")
	fmt.Println(result)
} 