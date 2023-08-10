package exstringer

import "fmt"

type Currency int

type Country string

const (
	BLR Currency = 0
	USD Currency = 1
)

const (
	BRASIL Country = "BRASIL"
	USA    Country = "USA"
)

type Account struct {
	Amount float64
	Unit   Currency
}

func createAccount(country Country) (*Account, bool) {
	switch country {
	case BRASIL:
		return &Account{Amount: 0, Unit: BLR}, true
	case USA:
		return &Account{Amount: 0, Unit: USD}, true
	default:
		return nil, false
	}
}

func (account *Account) depositAmount(value float64) {
	account.Amount += value
}

func (account *Account) showAmount() string {
	return account.Stringer()
}

func (curr Currency) Stringer() string {
	currencies := []string{"BLR", "USD"}
	return currencies[curr]
} 

func (account *Account) Stringer() string {
	return fmt.Sprintf("Você tem %.2f %s", account.Amount, account.Unit.Stringer())
}

func Exec() {
	var account *Account

	if acc, ok := createAccount(BRASIL); ok {
		account = acc
	}

	account.depositAmount(200)
	fmt.Print(account.showAmount())
}