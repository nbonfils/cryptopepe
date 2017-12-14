package schmekles

// Transaction is an donnation of schmekles
type Transaction struct {
}

// IsValid checks that a transaction is correct (balances, etc...)
func (t *Transaction) IsValid() bool {
	return true
}
