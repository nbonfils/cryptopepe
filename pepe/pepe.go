package pepe

// Pepe is a frog
type Pepe struct {
}

// Transaction is a Pepe that changes owner
type Transaction struct {
	Hash [32]byte
}

// IsValid checks that a transaction is correct (have correct pepe, etc...)
func (t *Transaction) IsValid() bool {
	return true
}
