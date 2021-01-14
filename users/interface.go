package users

// Users integration
type Users interface {
	Save(user User) error
	Get(userID int64) (*User, error)
	GetAll() ([]User, error)
	GetRedirectURLs(userID int64) ([]RedirectURL, error)
	Delete(userID int64) error
}
