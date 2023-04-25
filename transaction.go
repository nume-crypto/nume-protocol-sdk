package nume

type Transaction struct {
	Request *Request
}

func (transaction *Transaction) Create(data map[string]interface{}) (map[string]interface{}, error) {
	url := "store-transaction"
	return transaction.Request.Post(url, data)
}

func (transaction *Transaction) Get(tx_hash string) (map[string]interface{}, error) {
	url := "transaction/" + tx_hash
	return transaction.Request.Get(url, nil)
}
