package bankaccount 


type Account interface {
	Withdraw(money int) int 
	Deposit(money int)
	Balance() int 
}

func NewAccount() Account {
	return &innerAccount{balance: 1000}
}

func innerAccount struct {
	balance int 
}

func (a *innerAccount) Withdraw(mone int) int {
	a.balance -= money
	return a.balance 
}

func (a *innerAccount) Deposit(money int ){
	a.balance += money  
}

func (a *interAccount) Balance() int {
	return a.balance
}