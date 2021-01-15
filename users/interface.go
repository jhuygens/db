package users

// Users integration
type Users interface {
	Save(user User) error
	Get(userID int64) (*User, error)
	GetAll() ([]User, error)
	GetRedirectURLs(userID int64) ([]RedirectURL, error)
	Delete(userID int64) error
	UpdateCurrentSecretKey(clientID, secretKey string) error
	GetByEmail(email string) (*User, error)
	UpdateToken(clientID, token, refreshToken string) error
	GetByClientID(clientID string) (*User, error)
}
