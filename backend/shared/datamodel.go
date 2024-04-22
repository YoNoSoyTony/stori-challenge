package shared

import (
	"crypto/sha256"
	"encoding/hex"
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
	data := fmt.Sprintf("%s-%d-%s-%d", t.Email, t.Amount, t.Month, timestamp)
	hash := sha256.Sum256([]byte(data))
	t.TransactionID = hex.EncodeToString(hash[:])
}
