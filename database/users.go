package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

var userList []User

func (u User) Store() User {
	if u.ID != 0 {
		return u
	}
	u.ID = len(userList) + 1
	userList = append(userList, u)
	return u
}

func Find(email, password string) *User {
	for _, u := range userList {
		if u.Email == email && u.Password == password {
			return &u
		}
	}
	return nil
}
