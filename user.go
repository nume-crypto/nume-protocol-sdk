package nume

type User struct {
	Request *Request
}

func (user *User) Get(u string) (map[string]interface{}, error) {
	url := "fetch-user-data/" + u
	return user.Request.Get(url, nil)
}

func (user *User) GetUserBalance(u string) (map[string]interface{}, error) {
	url := "fetch-user-balance/" + u
	return user.Request.Get(url, nil)
}

func (user *User) GetUserCurrencyBalance(u string, c string) (map[string]interface{}, error) {
	url := "fetch-user-balance/" + u + "/" + c
	return user.Request.Get(url, nil)
}

func (user *User) GetUserProof(u string, c string) (map[string]interface{}, error) {
	url := "fetch-user-proof/" + u + "/" + c
	return user.Request.Get(url, nil)
}
