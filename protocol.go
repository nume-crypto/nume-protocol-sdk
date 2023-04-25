package nume

type Protocol struct {
	Request *Request
}

func (protocol *Protocol) GetCurrencies() (map[string]interface{}, error) {
	url := "get-currencies/"
	return protocol.Request.Get(url, nil)
}

func (protocol *Protocol) MonitorStates() (map[string]interface{}, error) {
	url := "monitor-states"
	return protocol.Request.Get(url, nil)
}

func (protocol *Protocol) GetBlockNumber() (map[string]interface{}, error) {
	url := "get-block-number"
	return protocol.Request.Get(url, nil)
}

func (protocol *Protocol) ContractAbi() (map[string]interface{}, error) {
	url := "contract-abi"
	return protocol.Request.Get(url, nil)
}
