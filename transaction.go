package nume

type Transaction struct {
	Request *Request
}

func (transaction *Transaction) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := "store-transaction"
	return transaction.Request.Post(url, data)
}
