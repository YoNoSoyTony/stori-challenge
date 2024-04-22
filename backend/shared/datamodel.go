package shared

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func CalculateMetrics(transactions []Transaction) (Metrics, error) {
	var metrics Metrics
	var totalPositive, totalNegative, totalTransactions int
	monthCounts := make(map[string]int)

	for _, transaction := range transactions {
		totalTransactions++
		if transaction.Amount > 0 {
			totalPositive += transaction.Amount
		} else {
			totalNegative += transaction.Amount
		}
		monthCounts[transaction.Month]++
	}

	metrics.Balance = totalPositive - totalNegative
	if totalPositive > 0 {
		metrics.PositiveAverage = totalPositive / totalTransactions
	}
	if totalNegative < 0 {
		metrics.NegativeAverage = totalNegative / totalTransactions
	}
	metrics.TransactionsByMonth = monthCounts
	metrics.Transactions = transactions

	return metrics, nil
}

// Metrics represents the calculated metrics for transactions.
type Metrics struct {
	Balance             int            `json:"balance"`
	PositiveAverage     int            `json:"positiveAverage"`
	NegativeAverage     int            `json:"negativeAverage"`
	TransactionsByMonth map[string]int `json:"transactionsByMonth"`
	Transactions        []Transaction  `json:"transactions"`
}

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
