package nume

type User struct {
	Request *Request
}

func (user *User) Create(data map[string]interface{}) (map[string]interface{}, error) {
	url := "initiate-registration"
	return user.Request.Post(url, data)
}

func (user *User) Get(u string) (map[string]interface{}, error) {
	url := "fetch-user-data/" + u
	return user.Request.Get(url, nil)
}

func (user *User) GetUserBalance(u string) (map[string]interface{}, error) {
	url := "fetch-user-balance/" + u
	return user.Request.Get(url, nil)
}
