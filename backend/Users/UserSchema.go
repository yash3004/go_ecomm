package Users

type User struct {
	UserID   string
	Username string
	Password string
	Email    string
	PhoneNo  string
}

type ReadUser struct {
	UserID   string
	Username string
	Email    string
	PhoneNo  string
}

type AuthUser struct {
	UserID   string
	Username string
	Password string
}

type UserAddress struct {
	Addressid   string
	UserId      string
	Address1    string
	Address2    string
	City        string
	State       string
	Country     string
	Postal_code int
	PhoneNo     int
}
