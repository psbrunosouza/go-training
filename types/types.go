package mytypes

import (
	"fmt"
	"time"
)

type BalanceType float64

type Restaurant struct {
	Name          string
	BirthDate     time.Time
	IsPetFriendly bool
}

type CreditCard struct {
	Balance BalanceType
}

func CreateRestaurant() {
	restaurant := Restaurant{
		Name: "American Day", 
		BirthDate: time.Date(1996, 06, 15, 0, 0, 0, 0, time.Now().Location()), 
		IsPetFriendly: false,
	}

	fmt.Println(restaurant)
}

func (creditCard *CreditCard) Payment(payment BalanceType) BalanceType {
	creditCard.Balance -= payment
	return creditCard.Balance
}