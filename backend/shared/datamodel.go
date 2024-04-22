package shared

import (
	"fmt"
	"time"
)

type Transaction struct {
	Email         string `json:"email"`
	Amount        int    `json:"amount"`
	Month         string `json:"month"`
	TransactionID string `json:"transactionId"`
}

func (t *Transaction) GenerateTransactionID() {
	timestamp := time.Now().UnixNano()
	t.TransactionID = fmt.Sprintf("%x", timestamp) // Simple hash for demonstration
}
