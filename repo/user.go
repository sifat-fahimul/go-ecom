package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

type UserRepo interface {
	Create(p User) (*User, error)
	Find(email, password string) (*User, error)
	// List() ([]*User, error)
	// Delete(id int) error
	// Update(p User) (*User, error)
}

type userRepo struct {
	userList []User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}
	user.ID = len(r.userList) + 1
	r.userList = append(r.userList, user)
	return &user, nil
}
func (r userRepo) Find(email, password string) (*User, error) {
	for _, u := range r.userList {
		if u.Email == email && u.Password == password {
			return &u, nil
		}
	}
	return nil, nil

}

// func (r userRepo) List() ([]*User, error)
// func (r userRepo) Delete(id int) error
// func (r userRepo) Update(p User) (*User, error)
