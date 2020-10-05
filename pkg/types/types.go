package types

type Money           int64      //amount
type PaymentCategory string     //category
type PaymentStatus   string     //status
type Phone           string     //number phone

//Status Payment
const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

//Info about payment
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

//User account info
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}

//Favorites tab
type Favorite struct {
	ID        string
	AccountID int64
	Name      string
	Amount    Money
	Category  PaymentCategory
}
